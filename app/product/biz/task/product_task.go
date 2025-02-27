package task

import (
	"context"
	"douyin_mall/product/biz/dal/redis"
	"douyin_mall/product/biz/model"
	"douyin_mall/product/biz/vo"
	"douyin_mall/product/conf"
	"douyin_mall/product/infra/elastic"
	kf "douyin_mall/product/infra/kafka"
	"github.com/IBM/sarama"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"strings"
	"time"
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
			Price:       product.Price,
			Picture:     product.Picture,
		},
	})
	if err != nil {
		return err
	}
	_, _, err = kf.Producer.SendMessage(&sarama.ProducerMessage{
		Topic: kf.Topic,
		Value: sarama.ByteEncoder(sonicData),
	})
	if err != nil {
		return err
	}
	return nil
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
	if err != nil {
		return err
	}
	_, _, err = kf.Producer.SendMessage(&sarama.ProducerMessage{
		Topic: kf.Topic,
		Value: sarama.ByteEncoder(sonicData),
	})
	if err != nil {
		return err
	}
	return nil
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
			Price:       product.Price,
			Picture:     product.Picture,
		},
	})
	if err != nil {
		return err
	}
	_, _, err = kf.Producer.SendMessage(&sarama.ProducerMessage{
		Topic: kf.Topic,
		Value: sarama.ByteEncoder(sonicData),
	})
	if err != nil {
		return err
	}
	return nil
}

func DeleteProductToElasticSearch(id int64) {
	body := vo.ProductSearchQueryBody{
		Query: &vo.ProductSearchQuery{
			Term: &vo.ProductSearchTermQuery{
				"id": id,
			},
		},
	}
	sonicData, err := sonic.Marshal(body)
	if err != nil {
		klog.Errorf("[DeleteProductToElasticSearch]%v 序列化失败, err: %v", body, err)
		return
	}
	request, err := esapi.DeleteByQueryRequest{
		Index: []string{"product"},
		Body:  strings.NewReader(string(sonicData)),
	}.Do(context.Background(), elastic.ElasticClient)
	//1 调用esapi.DeleteRequest删除product索引库中id为product.ID的数据
	if err != nil {
		klog.Errorf("es删除产品数据失败,err:%v", err)
		return
	}
	if request.StatusCode != http.StatusOK {
		klog.Errorf("es删除产品数据失败, StatusCode:%d", request.StatusCode)
	}
}
func DeleteProductToRedis(id int64) {
	key := "product:" + strconv.FormatInt(id, 10)
	_, err := redis.RedisClient.Del(context.Background(), key).Result()
	if err != nil {
		klog.Error("redis delete error", err)
		return
	}
	return
}
func UpdateProductToElasticSearch(p *vo.ProductSendToKafka) {
	body := vo.ProductSearchQueryBody{
		Query: &vo.ProductSearchQuery{
			Term: &vo.ProductSearchTermQuery{
				"id": p.ID,
			},
		},
		Doc: &vo.ProductSearchDoc{
			Name:        p.Name,
			Description: p.Description,
		},
	}
	sonicData, err := sonic.Marshal(body)
	if err != nil {
		klog.Errorf("[UpdateProductToElasticSearch]%v 序列化失败, err: %v", body, err)
		return
	}
	request, err := esapi.DeleteByQueryRequest{
		Index: []string{"product"},
		Body:  strings.NewReader(string(sonicData)),
	}.Do(context.Background(), elastic.ElasticClient)
	//1 调用esapi.DeleteRequest删除product索引库中id为product.ID的数据
	if err != nil || request.StatusCode != http.StatusOK {
		klog.Errorf("es更新产品数据失败")
	}
}
func UpdateProductToRedis(product *vo.ProductSendToKafka) {
	pro := vo.ProductRedisDataVo{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Picture:     product.Picture,
	}
	key := "product:" + strconv.FormatInt(product.ID, 10)
	//4 调用redis的set方法将数据导入到redis缓存中
	marshal, err := sonic.MarshalString(pro)
	if err != nil {
		klog.Errorf("序列化失败,err:%v", errors.WithStack(err))
		return
	}
	_, err = redis.RedisClient.Set(context.Background(), key, marshal, 1*time.Hour).Result()
	if err != nil {
		klog.Error("redis set error", err)
		return
	}
	return
}

func AddProductToElasticSearch(product *vo.ProductSendToKafka) {
	pro := vo.ProductSearchDataVo{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
	}
	sonicData, err := sonic.Marshal(pro)
	if err != nil {
		klog.Error("序列化失败", errors.WithStack(err))
	}
	//3 调用esapi.BulkRequest将数据导入到product索引库中
	response, _ := esapi.IndexRequest{
		Index: "product",
		Body:  strings.NewReader(string(sonicData)),
	}.Do(context.Background(), elastic.ElasticClient)
	print(response.StatusCode)
	return
}
func AddProductToRedis(product *vo.ProductSendToKafka) {
	pro := vo.ProductRedisDataVo{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Picture:     product.Picture,
	}
	key := "product:" + strconv.FormatInt(product.ID, 10)
	//4 调用redis的set方法将数据导入到redis缓存中
	marshal, err := sonic.MarshalString(pro)
	if err != nil {
		klog.Error("序列化失败", err)
		return
	}
	_, err = redis.RedisClient.Set(context.Background(), key, marshal, 1*time.Hour).Result()
	if err != nil {
		klog.Error("redis set error", err)
		return
	}
	return
}

type ProductKafkaHandler struct {
}

func (ProductKafkaHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (ProductKafkaHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h ProductKafkaHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	count := 0
	batchSize := 10
	for msg := range claim.Messages() {
		klog.Infof("消费者接受到消息，Received message: Topic=%s, Partition=%d, Offset=%d, Key=%s, Value=%s \n",
			msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
		session.MarkMessage(msg, "start")
		value := msg.Value //消息内容
		dataVo := vo.ProductKafkaDataVO{}
		_ = sonic.Unmarshal(value, &dataVo)
		//将消息反序列化成Product结构体
		switch dataVo.Type.Name {
		case vo.Add:
			AddProductToElasticSearch(&dataVo.Product)
			AddProductToRedis(&dataVo.Product)
		case vo.Update:
			UpdateProductToElasticSearch(&dataVo.Product)
			UpdateProductToRedis(&dataVo.Product)
		case vo.Delete:
			DeleteProductToElasticSearch(dataVo.Product.ID)
			DeleteProductToRedis(dataVo.Product.ID)
		}
		count++
		session.MarkMessage(msg, "done")
		if count >= batchSize {
			count = 0
			session.Commit()
		}
	}
	session.Commit()
	return nil
}

func Consumer() {
	config := sarama.NewConfig()
	strategies := make([]sarama.BalanceStrategy, 1)
	strategies[0] = sarama.NewBalanceStrategyRange()
	config.Consumer.Group.Rebalance.GroupStrategies = strategies
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// 创建消费者
	brokers := strings.Split(conf.GetConf().Kafka.BizKafka.BootstrapServers, ",")
	consumer, err := sarama.NewConsumerGroup(brokers, "product_group", config)
	handler := ProductKafkaHandler{}
	for {
		err = consumer.Consume(
			context.Background(),
			[]string{conf.GetConf().Kafka.BizKafka.ProductTopicId},
			handler,
		)
		if err != nil {
			klog.Error("Error from consumer: ", err)
		}
	}
}
