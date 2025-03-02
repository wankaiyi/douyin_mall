package producer

import (
	"context"
	"douyin_mall/order/conf"
	"douyin_mall/order/infra/kafka/constant"
	"github.com/IBM/sarama"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/google/uuid"
)

var (
	cancelOrderSuccessProducer sarama.SyncProducer
)

func InitCancelOrderSuccessProducer() {
	config := sarama.NewConfig()
	// 保证消息不丢失
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewHashPartitioner
	config.Producer.Idempotent = true
	config.Producer.Retry.Max = 3
	config.Producer.Transaction.ID = uuid.New().String()
	config.Net.MaxOpenRequests = 1

	cancelOrderSuccessProducer, err = sarama.NewSyncProducer(conf.GetConf().Kafka.BizKafka.BootstrapServers, config)
	if err != nil {
		panic(err.Error())
	}

	server.RegisterShutdownHook(func() {
		_ = cancelOrderSuccessProducer.Close()
	})

}

// SendCancelOrderSuccessMessage 取消订单成功后发送事务消息，释放库存
func SendCancelOrderSuccessMessage(ctx context.Context, orderId string) {
	err = producer.BeginTxn()
	if err != nil {
		klog.Errorf("开启事务失败: %v", err)
		return
	}
	err := sendMessage(ctx, constant.CancelOrderSuccessTopic, []byte(orderId), orderId, cancelOrderSuccessProducer)
	if err != nil {
		producer.AbortTxn()
		return
	}
	err = producer.CommitTxn()
	if err != nil {
		klog.Errorf("提交事务失败: %v", err)
	}
}

func SendCancelOrderSuccessMessages(ctx context.Context, orderIds []string) error {
	err = producer.BeginTxn()
	if err != nil {
		klog.Errorf("开启事务失败: %v", err)
		return err
	}
	for _, orderId := range orderIds {
		err := sendMessage(ctx, constant.CancelOrderSuccessTopic, []byte(orderId), orderId, cancelOrderSuccessProducer)
		if err != nil {
			producer.AbortTxn()
			return err
		}
	}
	err = producer.CommitTxn()
	if err != nil {
		klog.Errorf("提交事务失败: %v", err)
		return err
	}
	return nil
}
