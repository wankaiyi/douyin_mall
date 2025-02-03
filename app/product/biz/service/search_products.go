package service

import (
	"context"
	"douyin_mall/product/biz/vo"
	"douyin_mall/product/infra/elastic"
	product "douyin_mall/product/kitex_gen/product"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"io"
	"strings"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	query := req.Query
	queryBody := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  query,
				"fields": []string{"Name", "Description"},
			},
		},
	}
	jsonData, _ := json.Marshal(queryBody)
	//发往elastic
	//TODO 将关键词发往elastic，检索数据
	search, _ := esapi.SearchRequest{
		Index: []string{"product"},
		Body:  strings.NewReader(string(jsonData)),
	}.Do(context.Background(), &elastic.ElasticClient)
	//解析数据
	searchData, _ := io.ReadAll(search.Body)
	elasticSearchVo := vo.ProductSearchAllDataVo{}
	convertErr := json.Unmarshal(searchData, &elasticSearchVo)
	if convertErr != nil {
		return nil, convertErr
	}
	productHitsList := elasticSearchVo.Hits.Hits
	var products = []*product.Product{}
	for i := range productHitsList {
		productData := productHitsList[i].Source
		pro := product.Product{
			Name:        productData.Name,
			Description: productData.Description,
		}
		products = append(products, &pro)
	}
	//TODO 将返回的数据返回到前端
	resp = &product.SearchProductsResp{
		StatusCode: 0,
		StatusMsg:  "success",
		Results:    products,
	}
	return resp, nil
}
