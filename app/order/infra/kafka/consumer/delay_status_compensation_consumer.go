package consumer

import (
	"context"
	"douyin_mall/common/infra/kafka/tracing"
	"douyin_mall/order/biz/dal/mysql"
	model3 "douyin_mall/order/biz/model"
	"douyin_mall/order/conf"
	"douyin_mall/order/infra/kafka/constant"
	model2 "douyin_mall/order/infra/kafka/model"
	"douyin_mall/order/infra/rpc"
	"douyin_mall/rpc/kitex_gen/product"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"go.opentelemetry.io/otel"
)

type delayStatusCompensationMsgConsumerGroup struct{}

func (delayStatusCompensationMsgConsumerGroup) Setup(_ sarama.ConsumerGroupSession) error { return nil }
func (delayStatusCompensationMsgConsumerGroup) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}
func (h delayStatusCompensationMsgConsumerGroup) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	ctx := sess.Context()
	for msg := range claim.Messages() {
		topic := msg.Topic
		klog.Infof("收到消息，topic:%q partition:%d offset:%d  value:%s\n", topic, msg.Partition, msg.Offset, string(msg.Value))

		var delayStockCompensationMessage model2.DelayStockCompensationMessage
		err := sonic.Unmarshal(msg.Value, &delayStockCompensationMessage)
		if err != nil {
			klog.Errorf("解析消息失败，topic:%q partition:%d offset:%d  value:%s\n", topic, msg.Partition, msg.Offset, string(msg.Value))
			sess.MarkMessage(msg, "")
			sess.Commit()
			continue
		}

		msgCtx := otel.GetTextMapPropagator().Extract(ctx, tracing.NewConsumerMessageCarrier(msg))
		_, span := otel.Tracer("delay-status-compensation-consumer").Start(msgCtx, "consume-delay-status-compensation-message")
		otel.GetTextMapPropagator().Extract(ctx, tracing.NewConsumerMessageCarrier(msg))

		uuid := delayStockCompensationMessage.Uuid
		exist, err := model3.CheckOrderExist(ctx, mysql.DB, uuid)
		if err != nil {
			klog.Errorf("根据检查订单是否存在失败，uuid:%s", uuid)
			return err
		}
		if !exist {
			klog.Infof("订单创建失败，补偿释放库存，uuid:%s", uuid)
			var productItems []*product.ProductUnLockQuantity
			for _, productItem := range delayStockCompensationMessage.ProductItems {
				productItems = append(productItems, &product.ProductUnLockQuantity{
					ProductId: int64(productItem.ProductID),
					Quantity:  int64(productItem.Quantity),
				})
			}
			unlockResp, err := rpc.ProductClient.UnlockProductQuantity(ctx, &product.ProductUnLockQuantityRequest{
				Products: productItems,
			})
			if err != nil {
				klog.Errorf("根据订单uuid:%s释放库存失败，原因:%s", uuid, err.Error())
				return err
			}
			if unlockResp.StatusCode != 0 {
				klog.Infof("根据订单uuid:%s释放库存失败，原因:%s", uuid, unlockResp.StatusMsg)
			}
		}

		span.End()
		sess.MarkMessage(msg, "")
		sess.Commit()
	}
	return nil
}

func InitDelayStatusCompensationConsumer() {
	consumerConfig := sarama.NewConfig()
	consumerConfig.Version = sarama.V3_5_0_0
	consumerConfig.Consumer.Offsets.AutoCommit.Enable = false
	consumerConfig.Consumer.Offsets.Initial = sarama.OffsetNewest
	consumerConfig.Consumer.Offsets.Retry.Max = 3

	groupId := constant.DelayStockCompensationGroupId
	if conf.GetConf().Env == "dev" {
		groupId += "-dev"
	}
	cGroup, err := sarama.NewConsumerGroup(conf.GetConf().Kafka.BizKafka.BootstrapServers, groupId, consumerConfig)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			err = cGroup.Consume(
				context.Background(),
				[]string{constant.DelayStockCompensationTopic},
				delayStatusCompensationMsgConsumerGroup{},
			)
			if err != nil {
				fmt.Println(err.Error())
				break
			}
		}
	}()

	server.RegisterShutdownHook(func() {
		_ = cGroup.Close()
	})

}
