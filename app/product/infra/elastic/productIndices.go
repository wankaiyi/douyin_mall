package elastic

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"strings"
)

var mapping = `{
  "mappings": {
    "properties": {
      "name": {
        "type": "text",
        "analyzer": "ik_smart"
      },
      "description": {
        "type": "text"
      }
    }
  }
}
`

func ProduceIndicesInit() {
	// 构建请求
	productIndicesExist, err := esapi.IndicesExistsRequest{
		Index: []string{"product"},
	}.Do(context.Background(), &ElasticClient)
	if err != nil {
		hlog.Error(err)
		return
	}
	//如果product不存在，就创建这个索引库
	if productIndicesExist.StatusCode != 200 {
		create, err := esapi.IndicesCreateRequest{
			Index: "product",
			Body:  strings.NewReader(mapping),
		}.Do(context.Background(), &ElasticClient)
		if err != nil {
			hlog.Info(err)
		}
		if create.StatusCode != 200 {
			hlog.Error("create product indices failed")
		}
	}
}
