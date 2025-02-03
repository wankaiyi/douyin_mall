package mtl

import (
	"douyin_mall/common/infra/kafka"
	"douyin_mall/common/utils/env"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

func InitLog(logFileName string, logMaxSize int, logMaxBackups int, logMaxAge int, logLevel klog.Level, clsKafkaUser string, clsKafkaPassword string, clsKafkaTopicId string) {
	logger := kitexlogrus.NewLogger()
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
