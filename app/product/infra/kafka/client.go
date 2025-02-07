package kafka

import (
	"douyin_mall/product/conf"
	"github.com/segmentio/kafka-go"
	"time"
)

var (
	Writer *kafka.Writer
	Reader *kafka.Reader
)

func InitClient() {
	conf := conf.GetConf().KafkaEs
	writer := kafka.Writer{
		Addr:                   kafka.TCP(conf.KafkaWriter.Addr),
		Topic:                  conf.KafkaWriter.Topic,
		Balancer:               &kafka.Hash{},
		WriteTimeout:           time.Second,
		RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: true,
	}
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{conf.KafkaReader.Addr},
		Topic:          conf.KafkaReader.Topic,
		GroupID:        conf.KafkaReader.GroupID,
		StartOffset:    kafka.FirstOffset,
		CommitInterval: time.Second,
	})
	Writer = &writer
	Reader = reader
}

func CloseClient() {
	Writer.Close()
	Reader.Close()
}
