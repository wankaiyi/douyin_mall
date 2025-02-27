package add

import (
	"bytes"
	"context"
	"douyin_mall/product/biz/vo"
	"douyin_mall/product/infra/elastic/client"
	"douyin_mall/product/infra/kafka/model"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/pkg/errors"
)

func AddProductToElasticSearch(ctx context.Context, product *model.AddProductSendToKafka) (err error) {
	pro := vo.ProductSearchDataVo{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
	}
	sonicData, err := sonic.Marshal(pro)
	if err != nil {
		klog.Error("序列化失败", errors.WithStack(err))
		return errors.WithStack(err)
	}
	//3 调用esapi.BulkRequest将数据导入到product索引库中
	response, _ := esapi.IndexRequest{
		Index: "product",
		Body:  bytes.NewReader(sonicData),
	}.Do(ctx, client.ElasticClient)
	print(response.StatusCode)
	return nil
}
