package elastic

import (
	"douyin_mall/product/conf"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
)

var (
	ElasticClient elasticsearch7.Client
)

func InitClient() {
	elasticsearchConf := conf.GetConf().Elasticsearch
	var client *elasticsearch7.Client
	client, err := elasticsearch7.NewClient(elasticsearch7.Config{
		Addresses: []string{"http://" + elasticsearchConf.Host + ":" + elasticsearchConf.Port},
		Username:  elasticsearchConf.Username,
		Password:  elasticsearchConf.Password,
	})
	ElasticClient = *client
	if err != nil {
		hlog.Error(err)
		return
	}
	ProduceIndicesInit()

}
