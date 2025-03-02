package service

import (
	"bytes"
	"context"
	"crypto/md5"
	"douyin_mall/common/constant"
	keyModel "douyin_mall/common/infra/hot_key_client/model/key"
	"douyin_mall/common/infra/hot_key_client/processor"
	"douyin_mall/product/biz/dal/mysql"
	redisClient "douyin_mall/product/biz/dal/redis"
	"douyin_mall/product/biz/model"
	"douyin_mall/product/biz/vo"
	"douyin_mall/product/infra/elastic/client"
	product "douyin_mall/product/kitex_gen/product"
	"encoding/json"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/google/uuid"
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
							Fields: []string{"name", "description"},
						},
					},
				},
			},
		},
		Source: &vo.ProductSearchSource{
			"id",
		},
	}
	if req.Query == "" {
		queryBody.Query = &vo.ProductSearchQuery{
			MatchAll: &vo.All{},
		}
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
			_, err = redisClient.RedisClient.Set(s.ctx, dslKey, dslCacheToRedis, 5*time.Minute).Result()
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
	//根据返回的数据确认是否有缺失数据，有的话把当前的id存进去
	var missingIds []int64
	//先判断redis是否存在数据，如果存在，则直接返回数据
	if len(searchIds) > 0 {
		//加入hotkey
		var keysKey = searchHotkey(string(md5Bytes))
		keysConf := keyModel.NewKeyConf1(constant.ProductService)
		keysHotkeyModel := keyModel.NewHotKeyModelWithConfig(keysKey, &keysConf)
		var values map[int64]map[string]string = make(map[int64]map[string]string)
		localKeysCache := processor.GetValue(*keysHotkeyModel)
		if localKeysCache != nil {
			values = localKeysCache.(map[int64]map[string]string)
		} else {
			for _, id := range searchIds {
				productKey := model.BaseInfoKey(s.ctx, id)
				result, _ := redisClient.RedisClient.HGetAll(context.Background(), productKey).Result()
				if len(result) != 0 {
					values[id] = result
				} else {
					missingIds = append(missingIds, id)
				}
			}
			processor.SmartSet(*keysHotkeyModel, values)
		}
		for id, value := range values {
			if value == nil {
				missingIds = append(missingIds, id)
			} else {
				//解析数据
				valueMap := value
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
			productKey := model.BaseInfoKey(s.ctx, list[i].ProductId)
			err := model.PushToRedisBaseInfo(context.Background(), model.Product{
				ID:          list[i].ProductId,
				Name:        list[i].ProductName,
				Description: list[i].ProductDescription,
				Picture:     list[i].ProductPicture,
				Price:       list[i].ProductPrice,
				Stock:       list[i].ProductStock,
				LockStock:   list[i].ProductLockStock,
				PublicState: list[i].ProductPublicState,
				Sale:        list[i].ProductSale,
			}, redisClient.RedisClient, productKey)
			if err != nil {
				klog.CtxErrorf(s.ctx, "product数据缓存到redis失败, err: %v", err)
				return nil, err
			}
		}
		products = append(products, missingProducts...)
	}

	//根据商品id查询库存信息
	var productStock map[int64]int64 = make(map[int64]int64)
	for _, id := range searchIds {
		stockKey := model.StockKey(s.ctx, id)
		exists, err := redisClient.RedisClient.Exists(s.ctx, stockKey).Result()
		if err != nil {
			return nil, err
		}
		//如果存在，则从redis中获取数据
		if exists == 1 {
			//连续三次尝试
			maxTry := 3
			for i := 0; i < maxTry; i++ {
				result, err := redisClient.RedisClient.HGetAll(s.ctx, stockKey).Result()
				if err != nil {
					if i == maxTry-1 {
						return nil, errors.New("redis获取库存失败")
					}
				} else {
					productStock[id], _ = strconv.ParseInt(result["stock"], 10, 64)
					break
				}
			}
		} else {
			//不存在
			//则先加锁
			lockKey := model.StockLockKey(s.ctx, id)
			uuidString := uuid.New().String()
			result, err := redisClient.RedisClient.SetNX(s.ctx, lockKey, uuidString, 1000*time.Microsecond).Result()
			//如果加锁失败
			if err != nil || result == false {
				//连续三次尝试
				maxTry := 3
				for i := 0; i < maxTry; i++ {
					result, err := redisClient.RedisClient.HGetAll(s.ctx, stockKey).Result()
					if err != nil {
						if i == maxTry-1 {
							return nil, errors.New("redis获取库存失败")
						}
					} else {
						productStock[id], _ = strconv.ParseInt(result["stock"], 10, 64)
						break
					}
				}
			} else {
				//如果加锁成功，则从数据库中获取数据
				list, err := model.SelectProductList(mysql.DB, s.ctx, []int64{id})
				if err != nil {
					klog.CtxErrorf(s.ctx, "从数据库中获取数据失败, err: %v", err)
					return nil, err
				}
				if len(list) == 1 {
					productStock[id] = list[0].ProductStock
					err := model.PushToRedisStock(s.ctx, model.Product{
						ID:          id,
						Name:        list[0].ProductName,
						Description: list[0].ProductDescription,
						Picture:     list[0].ProductPicture,
						Price:       list[0].ProductPrice,
						Stock:       list[0].ProductStock,
						LockStock:   list[0].ProductLockStock,
						PublicState: list[0].ProductPublicState,
						Sale:        list[0].ProductSale,
					}, redisClient.RedisClient, stockKey)
					if err != nil {
						return nil, err
					}
				} else if len(list) == 0 {
					klog.CtxErrorf(s.ctx, "从数据库中获取数据失败, 商品id: %d 不存在", id)
				} else {
					klog.CtxErrorf(s.ctx, "从数据库中获取数据失败, 商品id: %d 存在多个", id)
				}
				err = model.SafeDeleteLock(s.ctx, redisClient.RedisClient, lockKey, uuidString)
				if err != nil {
					klog.CtxErrorf(s.ctx, "删除锁失败, err: %v", err)
				}
			}
		}
	}
	for _, pro := range products {
		pro.Stock = productStock[pro.Id]
	}

	//将返回的数据返回到前端
	resp = &product.SearchProductsResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
		Results:    products,
	}
	return
}

func searchHotkey(md string) string {
	return "product:hotkey:" + md
}
