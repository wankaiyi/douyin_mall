package kafka

import (
	"context"
	"douyin_mall/product/biz/model"
	"douyin_mall/product/biz/vo"
	"douyin_mall/product/infra/elastic"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/bytedance/sonic"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"strings"
)

func AddProduct(product *model.Product) (err error) {
	//TODO 用专门的结构体来封装发送上去的数据
	sonicData, err := sonic.Marshal(vo.ProductKafkaDataVO{
		Type: vo.Type{
			Name: vo.Add,
		},
		Product: vo.ProductSendToKafka{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
		},
	})
	Producer.Input() <- &sarama.ProducerMessage{
		Topic: Topic,
		Value: sarama.ByteEncoder(sonicData),
	}
	return
}

func DeleteProduct(product *model.Product) (err error) {
	//TODO 用专门的结构体来封装发送上去的数据
	sonicData, err := sonic.Marshal(vo.ProductKafkaDataVO{
		Type: vo.Type{
			Name: vo.Delete,
		},
		Product: vo.ProductSendToKafka{
			ID: product.ID,
		},
	})
	Producer.Input() <- &sarama.ProducerMessage{
		Topic: Topic,
		Value: sarama.ByteEncoder(sonicData),
	}
	return
}

func UpdateProduct(product *model.Product) (err error) {
	//TODO 用专门的结构体来封装发送上去的数据
	sonicData, err := sonic.Marshal(vo.ProductKafkaDataVO{
		Type: vo.Type{
			Name: vo.Update,
		},
		Product: vo.ProductSendToKafka{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
		},
	})
	Producer.Input() <- &sarama.ProducerMessage{
		Topic: Topic,
		Value: sarama.ByteEncoder(sonicData),
	}
	return
}
func DeleteProductToElasticSearch(id int64) {
	body := vo.ProductSearchQueryBody{
		Query: vo.ProductSearchQuery{
			Term: vo.ProductSearchTermQuery{
				"id": id,
			},
		},
	}
	sonicData, err := sonic.Marshal(body)
	if err != nil {
		return
	}
	request, _ := esapi.DeleteByQueryRequest{
		Index: []string{"product"},
		Body:  strings.NewReader(string(sonicData)),
	}.Do(context.Background(), elastic.ElasticClient)
	//1 调用esapi.DeleteRequest删除product索引库中id为product.ID的数据
	print(request.StatusCode) //2 打印返回的状态码
}

func UpdateProductToElasticSearch(p *vo.ProductSendToKafka) {
	body := vo.ProductSearchQueryBody{
		Query: vo.ProductSearchQuery{
			Term: vo.ProductSearchTermQuery{
				"id": p.ID,
			},
		},
		Doc: vo.ProductSearchDoc{
			Name:        p.Name,
			Description: p.Description,
		},
	}
	sonicData, err := sonic.Marshal(body)
	if err != nil {
		return
	}
	request, _ := esapi.DeleteByQueryRequest{
		Index: []string{"product"},
		Body:  strings.NewReader(string(sonicData)),
	}.Do(context.Background(), elastic.ElasticClient)
	//1 调用esapi.DeleteRequest删除product索引库中id为product.ID的数据
	print(request.StatusCode) //2 打印返回的状态码
}

func AddProductToElasticSearch(product *vo.ProductSendToKafka) {
	pro := vo.ProductSearchDataVo{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
	}
	sonicData, _ := sonic.Marshal(pro)
	//3 调用esapi.BulkRequest将数据导入到product索引库中
	response, _ := esapi.IndexRequest{
		Index: "product",
		Body:  strings.NewReader(string(sonicData)),
	}.Do(context.Background(), elastic.ElasticClient)
	print(response.StatusCode)
	return
}

type ProductKafkaConsumer struct {
}

func (ProductKafkaConsumer) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (ProductKafkaConsumer) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h ProductKafkaConsumer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		value := msg.Value //消息内容
		dataVo := vo.ProductKafkaDataVO{}
		_ = sonic.Unmarshal(value, &dataVo)
		//将消息反序列化成Product结构体
		switch dataVo.Type.Name {
		case vo.Add:
			AddProductToElasticSearch(&dataVo.Product)
		case vo.Update:
			UpdateProductToElasticSearch(&dataVo.Product)
		case vo.Delete:
			DeleteProductToElasticSearch(dataVo.Product.ID)
		}
		sess.MarkMessage(msg, "")
		sess.Commit()
	}
	return nil
}
