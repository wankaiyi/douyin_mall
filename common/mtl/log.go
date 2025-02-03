package mtl

import (
	"douyin_mall/common/infra/kafka"
	"douyin_mall/common/utils/env"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

type DouyinMallJSONFormatter struct {
	TimeZone    *time.Location
	ServiceName string
}

func (f *DouyinMallJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	entry.Time = entry.Time.In(f.TimeZone)

	data := logrus.Fields{
		"time":    entry.Time.Format(time.DateTime),
		"level":   entry.Level.String(),
		"msg":     entry.Message,
		"service": f.ServiceName,
	}

	for k, v := range entry.Data {
		data[k] = v
	}

	serialized, err := sonic.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON: %v", err)
	}
	return append(serialized, '\n'), nil
}

func InitLog(logFileName string, logMaxSize int, logMaxBackups int, logMaxAge int, logLevel klog.Level, clsKafkaUser string, clsKafkaPassword string, clsKafkaTopicId string, serviceName string) {
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("Error loading location:", err)
		panic(err)
	}

	log := logrus.New()
	log.SetReportCaller(true)
	log.SetFormatter(&DouyinMallJSONFormatter{
		TimeZone:    location,
		ServiceName: serviceName,
	})
	logger := kitexlogrus.NewLogger(kitexlogrus.WithLogger(log))
	klog.SetLogger(logger)

	if currentEnv, err := env.GetString("env"); err == nil && currentEnv == "dev" {
		klog.SetLevel(klog.LevelDebug)
		klog.SetOutput(os.Stdout)
	} else {
		klog.SetLevel(logLevel)
		fileWriter := zapcore.AddSync(&lumberjack.Logger{
			Filename:   logFileName,
			MaxSize:    logMaxSize,
			MaxBackups: logMaxBackups,
			MaxAge:     logMaxAge,
			LocalTime:  true,
		})

		kafkaWriter := kafka.NewKafkaWriter(
			clsKafkaUser,
			clsKafkaPassword,
			clsKafkaTopicId,
		)

		writeSyncers := zapcore.NewMultiWriteSyncer(fileWriter, kafkaWriter)
		asyncWriter := &zapcore.BufferedWriteSyncer{
			WS:            writeSyncers,
			FlushInterval: time.Second * 5,
		}
		klog.SetOutput(asyncWriter)
		server.RegisterShutdownHook(func() {
			asyncWriter.Sync()
		})
	}
}
