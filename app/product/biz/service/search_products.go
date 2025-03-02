package service

import (
	"bytes"
	"context"
	"crypto/md5"
	"douyin_mall/common/constant"
	keyModel "douyin_mall/common/infra/hot_key_client/model/key"
	"douyin_mall/common/infra/hot_key_client/processor"
	"douyin_mall/product/biz/dal/mysql"
	redis "douyin_mall/product/biz/dal/redis"
	"douyin_mall/product/biz/model"
	"douyin_mall/product/biz/vo"
	"douyin_mall/product/infra/elastic/client"
	product "douyin_mall/product/kitex_gen/product"
	"encoding/json"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/pkg/errors"
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
	from := new(int64)
	size := new(int64)
	var p int64 = 0
	if req.Page >= 1 {
		p = (req.Page - 1) * req.PageSize
	}
	from = &p
	size = &req.PageSize
	queryBody := vo.ProductSearchQueryBody{
		From: from,
		Size: size,
		Query: &vo.ProductSearchQuery{
			Bool: &vo.ProductSearchBoolQuery{
				Should: &[]*vo.ProductSearchQuery{
					{
						MultiMatch: &vo.ProductSearchMultiMatchQuery{
							Query:  req.Query,
							Fields: []string{"name", "description", "category_name"},
						},
					},
					//{
					//	MultiMatch: &vo.ProductSearchMultiMatchQuery{
					//		Query:  req.CategoryName,
					//		Fields: []string{"category_name"},
					//	},
					//},
				},
			},
			//MultiMatch: &vo.ProductSearchMultiMatchQuery{
			//	Query:  req.Query,
			//	Fields: []string{"name", "description"},
			//},
		},
		Source: &vo.ProductSearchSource{
			"id",
		},
	}
	if req.Query == "" {
		queryBody.Query.Bool = nil
		queryBody.Query.MatchAll = &vo.All{}
	}
	dslBytes, _ := sonic.Marshal(queryBody)
	//将dsl计算hashcode
	harsher := md5.New()
	harsher.Write(dslBytes)
	md5Bytes := harsher.Sum(nil)
	//从redis查找该hashcode对应的缓存数据
	dslKey := "product:dslBytes:" + string(md5Bytes)
	keyConf := keyModel.NewKeyConf1(constant.ProductService)
	hotkeyModel := keyModel.NewHotKeyModelWithConfig(dslKey, &keyConf)
	var dslCache string
	localDslCache := processor.GetValue(*hotkeyModel)
	if localDslCache != nil {
		dslCache = localDslCache.(string)
	}

	//搜索返回的id
	var searchIds []int64

	if dslCache != "" && dslCache != "null" {
		err = sonic.UnmarshalString(dslCache, &searchIds)
		if err != nil {
			klog.CtxErrorf(s.ctx, "dsl缓存反序列化失败, err: %v", err)
			return nil, errors.WithStack(err)
		}
	} else {
		//发往elastic
		search, err := esapi.SearchRequest{
			Index: []string{"product"},
			Body:  bytes.NewReader(dslBytes),
		}.Do(s.ctx, client.ElasticClient)
		if err != nil {
			klog.CtxErrorf(s.ctx, "es搜索失败, err: %v", err)
			return nil, errors.WithStack(err)
		}
		//解析数据
		searchData, err := io.ReadAll(search.Body)
		if err != nil {
			klog.CtxErrorf(s.ctx, "es搜索结果解析失败, err: %v", err)
			return nil, errors.WithStack(err)
		}
		elasticSearchVo := vo.ProductSearchAllDataVo{}
		err = json.Unmarshal(searchData, &elasticSearchVo)
		if err != nil {
			klog.CtxErrorf(s.ctx, "es搜索结果反序列化失败, err: %v", err)
			return nil, errors.WithStack(err)
		}
		productHitsList := elasticSearchVo.Hits.Hits
		if len(productHitsList) > 0 {
			for i := range productHitsList {
				sourceData := productHitsList[i].Source
				searchIds = append(searchIds, sourceData.ID)
			}

			//将es返回的数据缓存到redis
			dslCacheToRedis, _ := sonic.Marshal(searchIds)
			_, err = redis.RedisClient.Set(s.ctx, dslKey, dslCacheToRedis, 5*time.Minute).Result()
			if err != nil {
				klog.CtxErrorf(s.ctx, "dsl搜索结果缓存到redis失败, err: %v", err)
				return nil, errors.WithStack(err)
			}
			marshalString, err := sonic.MarshalString(searchIds)
			if err != nil {
				klog.CtxErrorf(s.ctx, "dsl搜索结果序列化失败, err: %v", err)
				return nil, errors.WithStack(err)
			}
			processor.SmartSet(*hotkeyModel, marshalString)
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
		//加入hotkey
		var keysKey = "product:keys:" + string(md5Bytes)
		keysConf := keyModel.NewKeyConf1(constant.ProductService)
		keysHotkeyModel := keyModel.NewHotKeyModelWithConfig(keysKey, &keysConf)
		var values []interface{}
		localKeysCache := processor.GetValue(*keysHotkeyModel)
		if localKeysCache != nil {
			values = localKeysCache.([]interface{})
		} else {
			for i := range keys {
				result, _ := redis.RedisClient.HGetAll(context.Background(), keys[i]).Result()
				if len(result) != 0 {
					values = append(values, result)
				} else {
					idStr := keys[i][len("product:"):] // 提取ID部分
					if id, err := strconv.ParseInt(idStr, 10, 64); err == nil {
						missingIds = append(missingIds, id)
					}
				}
			}
			//values, err = redis.RedisClient.MGet(context.Background(), keys...).Result()
			//if err != nil {
			//	klog.CtxErrorf(s.ctx, "redis获取dsl搜索结果失败, err: %v", err)
			//	return nil, errors.WithStack(err)
			//}
			processor.SmartSet(*keysHotkeyModel, values)
		}
		for i, value := range values {
			idStr := keys[i][len("product:"):] // 提取ID部分
			if value == nil {
				if id, err := strconv.ParseInt(idStr, 10, 64); err == nil {
					missingIds = append(missingIds, id)
				} else {
					klog.CtxErrorf(s.ctx, "productid转换失败, err: %v", err)
					return nil, errors.WithStack(err)
				}
			} else {
				//解析数据
				valueMap := value.(map[string]string)
				id, _ := strconv.ParseInt(valueMap["id"], 10, 64)
				Stock, _ := strconv.ParseInt(valueMap["stock"], 10, 64)
				Sale, _ := strconv.ParseInt(valueMap["sale"], 10, 64)
				PublishStatus, _ := strconv.ParseInt(valueMap["publish_status"], 10, 64)
				price, _ := strconv.ParseFloat(valueMap["price"], 64)
				productData := product.Product{
					Id:            id,
					Name:          valueMap["name"],
					Description:   valueMap["description"],
					Picture:       valueMap["picture"],
					Price:         float32(price),
					Stock:         Stock,
					Sale:          Sale,
					PublishStatus: PublishStatus,
				}
				products = append(products, &productData)
			}
		}
	}

	//如果不存在，则从数据库中获取数据，并存入redis
	if len(missingIds) > 0 {
		//从数据库中获取数据
		list, err := model.SelectProductList(mysql.DB, context.Background(), missingIds)
		if err != nil {
			klog.CtxErrorf(s.ctx, "从数据库中获取数据失败, err: %v", err)
			return nil, errors.WithStack(err)
		}
		missingProducts := make([]*product.Product, len(list))
		for i := range list {
			p := product.Product{
				Id:            list[i].ProductId,
				Name:          list[i].ProductName,
				Description:   list[i].ProductDescription,
				Picture:       list[i].ProductPicture,
				Price:         list[i].ProductPrice,
				Stock:         list[i].ProductStock,
				Sale:          list[i].ProductSale,
				CategoryId:    list[i].CategoryID,
				CategoryName:  list[i].CategoryName,
				PublishStatus: list[i].ProductPublicState,
			}
			missingProducts[i] = &p
			productKey := "product:" + strconv.FormatInt(list[i].ProductId, 10)
			err := model.PushToRedis(context.Background(), model.Product{
				ID:          list[i].ProductId,
				Name:        list[i].ProductName,
				Description: list[i].ProductDescription,
				Picture:     list[i].ProductPicture,
				Price:       list[i].ProductPrice,
				Stock:       list[i].ProductStock,
				LockStock:   list[i].ProductLockStock,
				PublicState: list[i].ProductPublicState,
			}, redis.RedisClient, productKey)
			if err != nil {
				klog.CtxErrorf(s.ctx, "product数据缓存到redis失败, err: %v", err)
				return nil, err
			}
			//_, err := redis.RedisClient.HSet(s.ctx, productKey, map[string]interface{}{
			//	"id":          list[i].ProductId,
			//	"name":        list[i].ProductName,
			//	"description": list[i].ProductDescription,
			//	"picture":     list[i].ProductPicture,
			//	"stock":       list[i].ProductStock,
			//}).Result()
			//if err != nil {
			//	return nil, err
			//}
		}
		products = append(products, missingProducts...)
	}
	//将返回的数据返回到前端
	resp = &product.SearchProductsResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
		Results:    products,
	}
	return
}
