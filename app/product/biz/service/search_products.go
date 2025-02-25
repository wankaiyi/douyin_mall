package service

import (
	"bytes"
	"context"
	"crypto/md5"
	"douyin_mall/common/constant"
	"douyin_mall/product/biz/dal/mysql"
	redis "douyin_mall/product/biz/dal/redis"
	"douyin_mall/product/biz/model"
	"douyin_mall/product/biz/vo"
	"douyin_mall/product/infra/elastic"
	product "douyin_mall/product/kitex_gen/product"
	"encoding/json"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"io"
	"strconv"
	"time"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	queryBody := vo.ProductSearchQueryBody{
		Query: &vo.ProductSearchQuery{
			MultiMatch: &vo.ProductSearchMultiMatchQuery{
				Query:  req.Query,
				Fields: []string{"name", "description"},
			},
		},
		Source: &vo.ProductSearchSource{
			"id",
		},
	}
	dslBytes, _ := sonic.Marshal(queryBody)
	//TODO 将dsl计算hashcode
	// 创建一个 MD5 哈希对象
	hasher := md5.New()
	hasher.Write(dslBytes)
	md5Bytes := hasher.Sum(nil)
	//从redis查找该hashcode对应的缓存数据
	dslKey := "product:dslBytes:" + string(md5Bytes)
	dslCache, err := redis.RedisClient.Get(context.Background(), dslKey).Result()
	//搜索返回的id
	var searchIds []int64

	if dslCache != "" && dslCache != "null" {
		err = sonic.UnmarshalString(dslCache, &searchIds)
		if err != nil {
			return
		}
	} else {
		//如果缓存数据存在，则直接返回数据
		//发往elastic
		search, _ := esapi.SearchRequest{
			Index: []string{"product"},
			Body:  bytes.NewReader(dslBytes),
		}.Do(context.Background(), elastic.ElasticClient)
		//解析数据
		searchData, _ := io.ReadAll(search.Body)
		elasticSearchVo := vo.ProductSearchAllDataVo{}
		err = json.Unmarshal(searchData, &elasticSearchVo)
		if err != nil {
			resp = &product.SearchProductsResp{
				StatusCode: 6013,
				StatusMsg:  constant.GetMsg(6013),
			}
			return
		}
		productHitsList := elasticSearchVo.Hits.Hits
		if len(productHitsList) > 0 {
			for i := range productHitsList {
				sourceData := productHitsList[i].Source
				searchIds = append(searchIds, sourceData.ID)
			}

			//将es返回的数据缓存到redis
			dslCacheToRedis, _ := sonic.Marshal(searchIds)
			_, _ = redis.RedisClient.Set(context.Background(), dslKey, dslCacheToRedis, 5*time.Minute).Result()
		}
	}
	var products []*product.Product = make([]*product.Product, 0)
	//根据id从缓存或者数据库钟获取数据
	//keys是redis的key
	var keys []string = make([]string, 0, len(searchIds))
	//根据返回的数据确认是否有缺失数据，有的话把当前的id存进去
	var missingIds []int64
	//将id转换为redis的key
	for i := range searchIds {
		keys = append(keys, "product:"+strconv.FormatInt(searchIds[i], 10))
	}
	//先判断redis是否存在数据，如果存在，则直接返回数据
	if len(keys) > 0 {
		values, err := redis.RedisClient.MGet(context.Background(), keys...).Result()
		for i, value := range values {
			idStr := keys[i][len("product:"):] // 提取ID部分
			if value == nil {
				if id, err := strconv.ParseInt(idStr, 10, 64); err == nil {
					missingIds = append(missingIds, id)
				}
			} else {
				//解析数据
				productData := product.Product{}
				err = sonic.UnmarshalString(value.(string), &productData)
				if err != nil {
					return nil, err
				}
				products = append(products, &productData)
			}
		}
	}

	//如果不存在，则从数据库中获取数据，并存入redis
	if len(missingIds) > 0 {
		//从数据库中获取数据
		list, modelErr := model.SelectProductList(mysql.DB, context.Background(), missingIds)
		pipeline := redis.RedisClient.Pipeline()
		if modelErr == nil {
			missingProducts := make([]*product.Product, len(list))
			for i := range list {
				p := product.Product{
					Id:          list[i].ID,
					Name:        list[i].Name,
					Description: list[i].Description,
					Picture:     list[i].Picture,
					Price:       list[i].Price,
					Stock:       list[i].Stock,
					Sale:        list[i].Sale,
					CategoryId:  list[i].CategoryId,
					BrandId:     list[i].BrandId,
				}
				missingProducts[i] = &p
				s2 := "product:" + strconv.FormatInt(list[i].ID, 10)
				marshalString, err := sonic.MarshalString(p)
				if err != nil {
					return nil, err
				}
				pipeline.Set(context.Background(), s2, marshalString, 1*time.Hour)
			}
			products = append(products, missingProducts...)
			//存入redis
			_, redisErr := pipeline.Exec(context.Background())
			if redisErr != nil {
				err = redisErr
				klog.Error("MSet products err:", err)
				return nil, redisErr
			}
		} else {
			err = modelErr
			return
		}
	}
	//将返回的数据返回到前端
	resp = &product.SearchProductsResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
		Results:    products,
	}
	return
}
