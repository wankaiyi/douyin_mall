package kafka

import (
	"context"
	"douyin_mall/payment/conf"
	"github.com/IBM/sarama"
	"strings"
	"time"
)

type Consumer struct {
	producer sarama.SyncProducer
	delay    time.Duration
}

func NewConsumer(producer sarama.SyncProducer, delay time.Duration) *Consumer {
	return &Consumer{producer, delay}
}

func (c *Consumer) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}
func (c *Consumer) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	for message := range claim.Messages() {
		now := time.Now().Local()
		if now.Sub(message.Timestamp) > c.delay {
			if string(message.Key) != "check-payment" {
				continue
			}
			//转发消息
			_, _, err := c.producer.SendMessage(&sarama.ProducerMessage{
				Topic: "check-topic-test",
				Key:   sarama.ByteEncoder(message.Key),
				Value: sarama.ByteEncoder(message.Value)})
			if err == nil {
				session.MarkMessage(message, "")
			}
			continue
		}
		time.Sleep(1 * time.Second)
		return nil

	}
	return nil
}
func DelayQueueInit() {
	config := sarama.NewConfig()
	strategies := make([]sarama.BalanceStrategy, 1)
	strategies[0] = sarama.NewBalanceStrategyRange()
	config.Consumer.Group.Rebalance.GroupStrategies = strategies
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	config.Producer.RequiredAcks = sarama.WaitForAll          // 等待所有副本确认
	config.Producer.Retry.Max = 5                             // 重试次数
	config.Producer.Return.Successes = true                   // 返回成功信息
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 随机分区

	bizKafka := conf.GetConf().Kafka.BizKafka
	brokers := strings.Split(bizKafka.BootstrapServers, ",")
	consumerGroup, err := sarama.NewConsumerGroup(brokers, "check-group-test", config)
	// 连接 Kafka Broker

	producer, err := sarama.NewSyncProducer(brokers, config)

	var consumer = NewConsumer(producer, 5*time.Second)

	for {
		if err = consumerGroup.Consume(context.Background(),
			[]string{"__delay-seconds-5"}, consumer); err != nil {
			panic(err)
		}
	}
}
