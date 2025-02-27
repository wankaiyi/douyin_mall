package elastic

import (
	"douyin_mall/product/conf"
	"github.com/cloudwego/kitex/pkg/klog"
	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	"sync"
)

var (
	ElasticClient *elasticsearch7.Client
	once          sync.Once
)

func InitClient() {
	elasticsearchConf := conf.GetConf().Elasticsearch
	var err error
	ElasticClient, err = elasticsearch7.NewClient(elasticsearch7.Config{
		Addresses: []string{"http://" + elasticsearchConf.Host + ":" + elasticsearchConf.Port},
		Username:  elasticsearchConf.Username,
		Password:  elasticsearchConf.Password,
	})
	if err != nil {
		klog.Errorf("es客户端初始化失败, err: %v", err)
		return
	}
}
