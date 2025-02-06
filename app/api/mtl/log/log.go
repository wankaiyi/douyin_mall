package log

import (
	"context"
	"douyin_mall/common/infra/kafka"
	"douyin_mall/common/mtl"
	"douyin_mall/common/utils/env"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var AsyncWriter *zapcore.BufferedWriteSyncer

func InitLog(serviceName string, logLevel hlog.Level, logFileName string, maxSize int, maxBackups int, maxAge int, h *server.Hertz, clsKafkaUser, clsKafkaPassword, clsKafkaTopicId string) {
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("Error loading location:", err)
		panic(err)
	}

	log := logrus.New()
	log.SetReportCaller(true)
	log.SetFormatter(&mtl.DouyinMallJSONFormatter{
		TimeZone:    location,
		ServiceName: serviceName,
	})
	logger := hertzlogrus.NewLogger(hertzlogrus.WithLogger(log))
	hlog.SetLogger(logger)
	if currentEnv, err := env.GetString("env"); err == nil && currentEnv == "dev" {
		hlog.SetLevel(hlog.LevelDebug)
		logger.SetOutput(os.Stdout)
	} else {
		hlog.SetLevel(logLevel)
		fileWriter := zapcore.AddSync(&lumberjack.Logger{
			Filename:   logFileName,
			MaxSize:    maxSize,
			MaxBackups: maxBackups,
			MaxAge:     maxAge,
		})

		kafkaWriter := kafka.NewKafkaWriter(
			clsKafkaUser,
			clsKafkaPassword,
			clsKafkaTopicId,
		)

		writeSyncers := zapcore.NewMultiWriteSyncer(fileWriter, kafkaWriter)
		AsyncWriter = &zapcore.BufferedWriteSyncer{
			WS:            writeSyncers,
			FlushInterval: time.Second * 5,
		}

		hlog.SetOutput(AsyncWriter)
		h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
			AsyncWriter.Sync()
		})
		logger.SetOutput(AsyncWriter)
	}

}
