package add

import (
	"douyin_mall/common/infra/kafka/tracing"
	"douyin_mall/product/infra/kafka/constant"
	"douyin_mall/product/infra/kafka/model"
	"github.com/IBM/sarama"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"go.opentelemetry.io/otel"
)

type UpdateProductHandler struct {
}

func (UpdateProductHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (UpdateProductHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h UpdateProductHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	ctx := session.Context()
	for msg := range claim.Messages() {
		klog.Infof("消费者接受到消息，Received message: Topic=%s, Partition=%d, Offset=%d, Key=%s, Value=%s \n",
			msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))

		session.MarkMessage(msg, "start")
		value := msg.Value //消息内容
		dataVo := model.UpdateProductSendToKafka{}
		//将消息反序列化成Product结构体
		_ = sonic.Unmarshal(value, &dataVo)
		msgCtx := otel.GetTextMapPropagator().Extract(ctx, tracing.NewConsumerMessageCarrier(msg))
		_, span := otel.Tracer(constant.ServiceName).Start(msgCtx, constant.UpdateService)
		otel.GetTextMapPropagator().Extract(ctx, tracing.NewConsumerMessageCarrier(msg))

		//发去redis
		err := UpdateProductToRedis(msgCtx, &dataVo)
		if err != nil {
			session.MarkMessage(msg, err.Error())
			return err
		}
		//发去es
		err = UpdateProductToElasticSearch(msgCtx, &dataVo)
		if err != nil {
			session.MarkMessage(msg, err.Error())
			return err
		}

		span.End()
		session.Commit()
	}

	return nil
}
