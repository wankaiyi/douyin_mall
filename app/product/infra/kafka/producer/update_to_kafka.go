package producer

import (
	"context"
	"douyin_mall/common/infra/kafka/tracing"
	"douyin_mall/product/infra/kafka/constant"
	"douyin_mall/product/infra/kafka/model"
	"github.com/IBM/sarama"
	"github.com/bytedance/sonic"
	"go.opentelemetry.io/otel"
)

func UpdateToKafka(ctx context.Context, product model.UpdateProductSendToKafka) error {
	//TODO 用专门的结构体来封装发送上去的数据
	sonicData, err := sonic.Marshal(model.AddProductSendToKafka{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Picture:     product.Picture,
	})
	if err != nil {
		return err
	}
	msg := &sarama.ProducerMessage{
		Topic: constant.UpdateTopic,
		Value: sarama.ByteEncoder(sonicData),
	}
	otel.GetTextMapPropagator().Inject(ctx, tracing.NewProducerMessageCarrier(msg))
	_, _, err = Producer.SendMessage(msg)
	if err != nil {
		return err
	}
	return nil
}
