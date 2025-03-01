package check

import (
	"bytes"
	"context"
	"douyin_mall/product/biz/dal/mysql"
	"douyin_mall/product/biz/model"
	"douyin_mall/product/biz/vo"
	"douyin_mall/product/infra/elastic/client"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

func ProduceIndicesInit() {
	ctx := context.Background()
	// 构建请求
	productIndicesExist, err := esapi.IndicesExistsRequest{
		Index: []string{"product"},
	}.Do(ctx, client.ElasticClient)
	if err != nil {
		klog.Errorf("查询索引库product失败,err:%v", err)
		return
	}
	//如果product不存在，就创建这个索引库
	if productIndicesExist.StatusCode != 200 {

		SettingData, err := sonic.Marshal(vo.ProductSearchMappingSetting)
		if err != nil {
			klog.Errorf("序列化ProductSearchMappingSetting失败,err:%v", err)
			return
		}
		create, err := esapi.IndicesCreateRequest{
			Index: "product",
			Body:  bytes.NewReader(SettingData),
		}.Do(ctx, client.ElasticClient)
		if err != nil {
			klog.Errorf("创建索引库product失败,err:%v", err)
		}
		body := create.Body
		fmt.Printf("%v", body)
		if create.StatusCode != 200 {
			klog.Error("create product indices failed")
			return
		}
		//将数据导入到product索引库中
		//1 从数据库中获取数据
		products, err := model.SelectProductAllWithoutCondition(mysql.DB, ctx)
		if err != nil {
			klog.Errorf("查询数据库失败,err:%v", err)
			return
		}
		//2 遍历数据，将数据转换为sonic格式
		for i := range products {
			pro := products[i]
			dataVo := vo.ProductSearchDataVo{
				Name:         pro.ProductName,
				Description:  pro.ProductDescription,
				ID:           pro.ProductId,
				CategoryName: pro.CategoryName,
			}
			sonicData, _ := sonic.Marshal(dataVo)
			//3 调用esapi.BulkRequest将数据导入到product索引库中
			_, err = esapi.IndexRequest{
				Index:   "product",
				Body:    bytes.NewReader(sonicData),
				Refresh: "true",
			}.Do(ctx, client.ElasticClient)
			if err != nil {
				klog.Errorf("导入数据到索引库product失败,err:%v", err)
				return
			}
		}
	}
}
