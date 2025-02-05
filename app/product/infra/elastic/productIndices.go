package elastic

import (
	"context"
	"douyin_mall/product/biz/dal/mysql"
	"douyin_mall/product/biz/model"
	"douyin_mall/product/biz/vo"
	"encoding/json"
	"github.com/bytedance/sonic"
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
		SettingData, err := sonic.Marshal(vo.ProductSearchMappingSetting)
		if err != nil {
			return
		}
		create, err := esapi.IndicesCreateRequest{
			Index: "product",
			Body:  strings.NewReader(string(SettingData)),
		}.Do(context.Background(), &ElasticClient)
		if err != nil {
			hlog.Info(err)
		}
		if create.StatusCode != 200 {
			hlog.Error("create product indices failed")
			return
		}
		//将数据导入到product索引库中
		//1 从数据库中获取数据
		var products []model.Product
		result := mysql.DB.Table("tb_product").Select("*").Find(&products)
		if result.Error != nil {
			hlog.Error(result.Error)
			return
		}
		//2 遍历数据，将数据转换为json格式
		for i := range products {
			pro := products[i]
			dataVo := vo.ProductSearchDataVo{
				Name:        pro.Name,
				Description: pro.Description,
			}
			jsonData, _ := json.Marshal(dataVo)
			//3 调用esapi.BulkRequest将数据导入到product索引库中
			_, _ = esapi.IndexRequest{
				Index:   "product",
				Body:    strings.NewReader(string(jsonData)),
				Refresh: "true",
			}.Do(context.Background(), &ElasticClient)
		}
	}
}
