package producer

import (
	"context"
	"douyin_mall/product/biz/vo"
	"douyin_mall/product/infra/kafka/constant"
	"douyin_mall/product/infra/kafka/model"
	"github.com/IBM/sarama"
	"github.com/bytedance/sonic"
)

func DeleteToKafka(ctx context.Context, product model.DeleteProductSendToKafka) error {
	sonicData, err := sonic.Marshal(vo.ProductKafkaDataVO{
		Type: vo.Type{
			Name: vo.Delete,
		},
		Product: vo.ProductSendToKafka{
			ID: product.ID,
		},
	})
	if err != nil {
		return err
	}
	_, _, err = Producer.SendMessage(&sarama.ProducerMessage{
		Topic: constant.DelTopic,
		Value: sarama.ByteEncoder(sonicData),
	})
	if err != nil {
		return err
	}
	return nil
}
