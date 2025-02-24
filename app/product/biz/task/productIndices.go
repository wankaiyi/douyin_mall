package task

import (
	"context"
	"douyin_mall/product/biz/dal/mysql"
	"douyin_mall/product/biz/model"
	"douyin_mall/product/biz/vo"
	"douyin_mall/product/infra/elastic"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"strings"
)

func ProduceIndicesInit() {
	// 构建请求
	productIndicesExist, err := esapi.IndicesExistsRequest{
		Index: []string{"product"},
	}.Do(nil, elastic.ElasticClient)
	if err != nil {
		klog.Error(err)
		return
	}
	//如果product不存在，就创建这个索引库
	if productIndicesExist.StatusCode != 200 {
		SettingData, err := sonic.Marshal(vo.ProductSearchMappingSetting)
		s := string(SettingData)
		fmt.Printf("%v", s)
		if err != nil {
			return
		}
		create, err := esapi.IndicesCreateRequest{
			Index: "product",
			Body:  strings.NewReader(s),
		}.Do(context.Background(), elastic.ElasticClient)
		if err != nil {
			klog.Info(err)
		}
		body := create.Body
		fmt.Printf("%v", body)
		if create.StatusCode != 200 {
			klog.Error("create product indices failed")
			return
		}
		//将数据导入到product索引库中
		//1 从数据库中获取数据
		var products []model.Product
		result := mysql.DB.Table("tb_product").Select("*").Find(&products)
		if result.Error != nil {
			klog.Error(result.Error)
			return
		}
		//2 遍历数据，将数据转换为sonic格式
		for i := range products {
			pro := products[i]
			dataVo := vo.ProductSearchDataVo{
				Name:        pro.Name,
				Description: pro.Description,
				ID:          pro.ID,
			}
			sonicData, _ := sonic.Marshal(dataVo)
			//3 调用esapi.BulkRequest将数据导入到product索引库中
			_, _ = esapi.IndexRequest{
				Index:   "product",
				Body:    strings.NewReader(string(sonicData)),
				Refresh: "true",
			}.Do(context.Background(), elastic.ElasticClient)
		}
	}
}
