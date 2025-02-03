package service

import (
	"context"
	"douyin_mall/product/biz/dal/mysql"
	"douyin_mall/product/biz/model"
	"douyin_mall/product/infra/elastic"
	product "douyin_mall/product/kitex_gen/product"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"io"
	"log"
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
				"fields": []string{"description", "name"},
			},
		},
	}
	jsonData, err := json.Marshal(queryBody)
	//发往elastic
	//TODO 将关键词发往elastic，检索数据
	search, _ := esapi.SearchRequest{
		Index: []string{"product"},
		Body:  strings.NewReader(string(jsonData)),
	}.Do(context.Background(), &elastic.ElasticClient)
	hlog.Info(search)
	//解析数据
	searchData, _ := io.ReadAll(search.Body)
	elasticSearchVo := map[string]interface{}{}
	convertErr := json.Unmarshal(searchData, &elasticSearchVo)
	if convertErr != nil {
		return nil, convertErr
	}
	hits := elasticSearchVo["hits"].(map[string]interface{})
	productHitsList := hits["hits"].([]interface{})
	hlog.Info(elasticSearchVo, hits, productHitsList)
	var products []*product.Product
	for i := range productHitsList {
		productData := productHitsList[i].(map[string]interface{})
		source := productData["_source"].(map[string]interface{})
		description := source["description"]
		name := source["name"]
		pro := product.Product{
			Name:        name.(string),
			Description: description.(string),
		}
		products = append(products, &pro)
	}
	//TODO 将返回的数据返回到前端
	resp = &product.SearchProductsResp{
		StatusCode: 0,
		StatusMsg:  "success",
		Results:    products,
	}
	return
}

func sss() {
	var products model.Product
	result := mysql.DB.Table("tb_product").Select("*").Limit(1).Find(&products)
	if result.Error != nil {
		hlog.Error(result.Error)
		return
	}
	data := map[string]string{
		"name":        products.Name,
		"description": products.Description,
	}
	// 将 map 转换为 JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error marshaling data: %s", err)
	}
	add, _ := esapi.IndexRequest{
		Index:   "product",
		Body:    strings.NewReader(string(jsonData)),
		Refresh: "true",
	}.Do(context.Background(), &elastic.ElasticClient)
	if err != nil {
		hlog.Error(add, err)
	}
}
