package producer

import (
	"context"
	"douyin_mall/checkout/conf"
	"douyin_mall/checkout/infra/kafka/constant"
	"douyin_mall/checkout/infra/kafka/model"
	commonModel "douyin_mall/common/infra/kafka/model"
	"douyin_mall/common/infra/kafka/tracing"
	"github.com/IBM/sarama"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"go.opentelemetry.io/otel"
)

var (
	producer sarama.SyncProducer
	err      error
)

func InitDelayStockCompensationProducer() {
	config := sarama.NewConfig()
	// 保证消息不丢失
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewHashPartitioner
	config.Producer.Idempotent = true
	config.Producer.Transaction.Retry.Backoff = 10
	config.Producer.Retry.Max = 3
	config.Producer.Transaction.ID = "delay-stock-compensation-producer"
	config.Net.MaxOpenRequests = 1

	producer, err = sarama.NewSyncProducer(conf.GetConf().Kafka.BizKafka.BootstrapServers, config)
	if err != nil {
		panic(err.Error())
	}

	server.RegisterShutdownHook(func() {
		_ = producer.Close()
	})
}

// SendDelayStockCompensationMessage 锁定库存后，延时检查订单是否创建成功
func SendDelayStockCompensationMessage(ctx context.Context, uuid string, productItems []model.ProductItem) {
	err = producer.BeginTxn()
	if err != nil {
		klog.Errorf("开启事务失败: %v", err)
		return
	}
	msg := &model.DelayStockCompensationMessage{
		Uuid:         uuid,
		ProductItems: productItems,
	}
	delayMsg := &commonModel.DelayMessage{
		Level: constant.DelayTopic30sLevel,
		Topic: constant.DelayStockCompensationTopic,
		Key:   uuid,
		Value: msg.ToJson(),
	}
	bytes, _ := sonic.Marshal(delayMsg)
	err = sendMessage(ctx, constant.DelayTopic, bytes, uuid, producer)
	if err != nil {
		producer.AbortTxn()
		return
	}
	err = producer.CommitTxn()
	if err != nil {
		klog.Errorf("提交事务失败: %v", err)
		return
	}
	klog.Infof("延迟库存补偿消息发送成功，uuid: %v", uuid)
}

func sendMessage(ctx context.Context, topic string, message []byte, key string, producer sarama.SyncProducer) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
		Key:   sarama.StringEncoder(key),
	}

	otel.GetTextMapPropagator().Inject(ctx, tracing.NewProducerMessageCarrier(msg))

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}
	klog.Infof("消息发送成功, topic: %s, partition: %d, offset: %d", topic, partition, offset)
	return nil
}
