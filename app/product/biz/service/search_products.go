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
	"sync"
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
	if req.Page >= 1 {
		*from = (req.Page - 1) * req.PageSize
	}
	if req.PageSize < 1 {
		*size = 0
	} else {
		size = &req.PageSize
	}
	size = &req.PageSize
	queryBody := vo.ProductSearchQueryBody{
		From: from,
		Size: size,
		Query: &vo.ProductSearchQuery{
			Bool: &vo.ProductSearchBoolQuery{
				Must:   &[]*vo.ProductSearchQuery{},
				Should: &[]*vo.ProductSearchQuery{},
			},
		},
		Source: &vo.ProductSearchSource{
			"id",
		},
	}
	if req.Query != "" {
		v := &vo.ProductSearchQuery{
			MultiMatch: &vo.ProductSearchMultiMatchQuery{
				Query:  req.Query,
				Fields: []string{"name", "description"},
			},
		}
		should := *queryBody.Query.Bool.Should
		should = append(should, v)
		queryBody.Query.Bool.Should = &should
		must := *queryBody.Query.Bool.Must
		must = append(must, v)
		queryBody.Query.Bool.Must = &must
	}
	if req.CategoryId > 0 {
		must := *queryBody.Query.Bool.Must
		must = append(must, &vo.ProductSearchQuery{
			MultiMatch: &vo.ProductSearchMultiMatchQuery{
				Query:  req.CategoryId,
				Fields: []string{"category_id"},
			},
		})
		queryBody.Query.Bool.Must = &must
	}
	dslBytes, _ := sonic.Marshal(queryBody)
	//将dsl计算hashcode
	harsher := md5.New()
	harsher.Write(dslBytes)
	md5Bytes := harsher.Sum(nil)

	//搜索返回的id
	searchIds, err := getSearchIds(s.ctx, dslBytes, md5Bytes)
	klog.CtxInfof(s.ctx, "searchIds: %v", searchIds)
	var products = make([]*product.Product, 0)
	//根据id从缓存或者数据库钟获取数据
	//根据返回的数据确认是否有缺失数据，有的话把当前的id存进去
	var missingIds []int64
	//先判断redis是否存在数据，如果存在，则直接返回数据
	products, missingIds, err = getCache(s.ctx, searchIds, md5Bytes)
	klog.CtxInfof(s.ctx, "products: %v,missingsIds:%v", products, missingIds)
	if err != nil {
		klog.CtxErrorf(s.ctx, "getCache: missingsIds:%v,err:%v", missingIds, err)
		return nil, err
	}

	//如果不存在，则从数据库中获取数据，并存入redis
	missingProduct, err := getMissingProduct(s.ctx, missingIds)
	if err != nil {
		klog.CtxErrorf(s.ctx, "getMissingProduct: err:%v", err)
		return nil, err
	}
	products = append(products, missingProduct...)
	klog.CtxInfof(s.ctx, "搜索的products: %v", products)

	//根据商品id查询库存信息
	productStock, err := getStock(s.ctx, searchIds)
	if err != nil {
		klog.CtxErrorf(s.ctx, "获取库存时, err: %v", err)
		return nil, err
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

func getSearchIds(ctx context.Context, dslBytes []byte, md5bytes []byte) ([]int64, error) {
	var ids = make([]int64, 0)
	dslKey := "product:dslBytes:" + string(md5bytes)
	keyConf := keyModel.NewKeyConf1(constant.ProductService)
	hotkeyModel := keyModel.NewHotKeyModelWithConfig(dslKey, &keyConf)
	var dslCache string
	if dslCache != "" && dslCache != "null" {
		err := sonic.UnmarshalString(dslCache, &ids)
		if err != nil {
			klog.CtxErrorf(ctx, "dsl缓存反序列化失败, err: %v", err)
			return ids, err
		}
	} else {
		//发往elastic
		search, err := esapi.SearchRequest{
			Index: []string{"product"},
			Body:  bytes.NewReader(dslBytes),
		}.Do(ctx, client.ElasticClient)
		if err != nil {
			klog.CtxErrorf(ctx, "es搜索失败, err: %v", err)
			return nil, errors.WithStack(err)
		}
		//解析数据
		searchData, err := io.ReadAll(search.Body)
		if err != nil {
			klog.CtxErrorf(ctx, "es搜索结果解析失败, err: %v", err)
			return nil, errors.WithStack(err)
		}
		elasticSearchVo := vo.ProductSearchAllDataVo{}
		err = json.Unmarshal(searchData, &elasticSearchVo)
		if err != nil {
			klog.CtxErrorf(ctx, "es搜索结果反序列化失败, err: %v", err)
			return nil, errors.WithStack(err)
		}
		productHitsList := elasticSearchVo.Hits.Hits
		if len(productHitsList) > 0 {
			for i := range productHitsList {
				sourceData := productHitsList[i].Source
				ids = append(ids, sourceData.ID)
			}

			//将es返回的数据缓存到redis
			dslCacheToRedis, _ := sonic.Marshal(ids)
			_, err = redisClient.RedisClient.Set(ctx, dslKey, dslCacheToRedis, 5*time.Minute).Result()
			if err != nil {
				klog.CtxErrorf(ctx, "dsl搜索结果缓存到redis失败, err: %v", err)
				return nil, errors.WithStack(err)
			}
			marshalString, err := sonic.MarshalString(ids)
			if err != nil {
				klog.CtxErrorf(ctx, "dsl搜索结果序列化失败, err: %v", err)
				return nil, errors.WithStack(err)
			}
			processor.SmartSet(*hotkeyModel, marshalString)
		}
	}
	return ids, nil
}

func getCache(ctx context.Context, searchIds []int64, md5Bytes []byte) (products []*product.Product, missingIds []int64, err error) {
	//加入hotkey
	products = make([]*product.Product, 0)
	missingIds = make([]int64, 0)
	if len(searchIds) == 0 {
		return products, missingIds, nil
	}
	var keysKey = searchHotkey(string(md5Bytes))
	keysConf := keyModel.NewKeyConf1(constant.ProductService)
	keysHotkeyModel := keyModel.NewHotKeyModelWithConfig(keysKey, &keysConf)
	var values map[int64]map[string]string = make(map[int64]map[string]string)
	localKeysCache := processor.GetValue(*keysHotkeyModel)
	if localKeysCache != nil {
		values = localKeysCache.(map[int64]map[string]string)
	} else {
		for _, id := range searchIds {
			productKey := model.BaseInfoKey(ctx, id)
			result, _ := redisClient.RedisClient.HGetAll(ctx, productKey).Result()
			klog.CtxInfof(ctx, "redis hgetall %v 获取的结果:%v", productKey, result)
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
			picture := valueMap["picture"]
			productData := product.Product{
				Id:            id,
				Name:          valueMap["name"],
				Description:   valueMap["description"],
				Picture:       picture,
				Price:         float32(price),
				Stock:         Stock,
				Sale:          Sale,
				PublishStatus: PublishStatus,
			}
			klog.CtxInfof(ctx, "valueMap的值:%v,productData:%v", valueMap, &productData)
			klog.CtxInfof(ctx, "id:%v,picture:%v", id, picture)
			products = append(products, &productData)
		}
	}
	return products, missingIds, nil
}

func getMissingProduct(ctx context.Context, missingIds []int64) (products []*product.Product, err error) {
	if len(missingIds) > 0 {
		//从数据库中获取数据
		list, err := model.SelectProductList(mysql.DB, context.Background(), missingIds)
		klog.CtxInfof(ctx, "数据库的数据:%v", list)
		if err != nil {
			klog.CtxErrorf(ctx, "从数据库中获取数据失败, err: %v", err)
			return nil, errors.WithStack(err)
		}
		var wg sync.WaitGroup
		var productMap = sync.Map{}
		for i := range list {
			wg.Add(1)
			go func(ctx context.Context, wg *sync.WaitGroup, productMap *sync.Map, pro *model.ProductWithCategory) {
				defer wg.Done()
				uuidString := uuid.New().String()
				lockKey := model.BaseInfoLockKey(ctx, pro.ProductId)
				lock, err := model.SetLock(ctx, redisClient.RedisClient, lockKey, uuidString)
				if err != nil {
					klog.CtxInfof(ctx, "key %v 上锁失败", lockKey)
					return
				}
				if lock {
					p := product.Product{
						Id:            pro.ProductId,
						Name:          pro.ProductName,
						Description:   pro.ProductDescription,
						Picture:       pro.ProductPicture,
						Price:         pro.ProductPrice,
						Stock:         pro.ProductStock,
						Sale:          pro.ProductSale,
						CategoryId:    pro.CategoryID,
						CategoryName:  pro.CategoryName,
						PublishStatus: pro.ProductPublicState,
					}
					productMap.Store(uuidString, &p)
					klog.CtxInfof(ctx, "从数据库内查找的数据:%v", &p)
					productKey := model.BaseInfoKey(ctx, list[i].ProductId)
					err = model.PushToRedisBaseInfo(ctx, model.Product{
						ID:          pro.ProductId,
						Name:        pro.ProductName,
						Description: pro.ProductDescription,
						Picture:     pro.ProductPicture,
						Price:       pro.ProductPrice,
						Stock:       pro.ProductStock,
						LockStock:   pro.ProductLockStock,
						PublicState: pro.ProductPublicState,
						Sale:        pro.ProductSale,
					}, redisClient.RedisClient, productKey)
					if err != nil {
						klog.CtxErrorf(ctx, "product数据缓存到redis失败, err: %v", err)
					}
					err := model.SafeDeleteLock(ctx, redisClient.RedisClient, lockKey, uuidString)
					if err != nil {
						klog.CtxInfof(ctx, "删除锁失败")
						return
					}
				}

			}(ctx, &wg, &productMap, &list[i])
		}
		wg.Wait()
		productMap.Range(func(key, value interface{}) bool {
			products = append(products, value.(*product.Product))
			return true
		})
	}
	return products, nil
}

func getStock(ctx context.Context, searchIds []int64) (productStock map[int64]int64, err error) {
	var wg sync.WaitGroup
	productStock = make(map[int64]int64)
	synMap := sync.Map{}
	for _, id := range searchIds {
		wg.Add(1)
		go func(wg *sync.WaitGroup, ctx context.Context, productStock *sync.Map) {
			defer wg.Done()

			stockKey := model.StockKey(ctx, id)
			exists, err := redisClient.RedisClient.Exists(ctx, stockKey).Result()
			if err != nil {
				klog.CtxErrorf(ctx, "获取库存信息时判断key :%v , err: %v", stockKey, err)
				return
			}
			//如果存在，则从redis中获取数据
			if exists == 1 {
				//连续三次尝试
				maxTry := 3
				for i := 0; i < maxTry; i++ {
					result, err := redisClient.RedisClient.HGetAll(ctx, stockKey).Result()
					if err != nil {
						if i == maxTry-1 {
							klog.CtxErrorf(ctx, "获取库存信息时超过最大重试次数%v key:%v, err: %v", maxTry, stockKey, err)
							return
						}
					} else {
						stock, err := strconv.ParseInt(result["stock"], 10, 64)
						if err != nil {
							klog.CtxErrorf(ctx, "获取库存信息时 库存由string转int64异常,stock:%v ,stock:%v,err: %v", result["stock"], stockKey, err)
						}
						productStock.Store(id, stock)
						break
					}
				}
			} else {
				//不存在
				//则先加锁
				lockKey := model.StockLockKey(ctx, id)
				uuidString := uuid.New().String()
				result, err := redisClient.RedisClient.SetNX(ctx, lockKey, uuidString, 2*time.Second).Result()
				//如果加锁失败
				if err != nil || result == false {
					//连续三次尝试
					maxTry := 3
					for i := 0; i < maxTry; i++ {
						result, err := redisClient.RedisClient.HGetAll(ctx, stockKey).Result()
						if err != nil {
							if i == maxTry-1 {
								return
							}
						} else {
							stock, err := strconv.ParseInt(result["stock"], 10, 64)
							if err != nil {
								klog.CtxErrorf(ctx, "获取库存信息时 库存由string转int64异常,stock:%v ,stock:%v,err: %v", result["stock"], stockKey, err)
							}
							productStock.Store(id, stock)
							break
						}
					}
				} else {
					//如果加锁成功，则从数据库中获取数据
					list, err := model.SelectProductList(mysql.DB, ctx, []int64{id})
					if err != nil {
						klog.CtxErrorf(ctx, "获取库存信息时 从数据库中获取数据失败, err: %v", err)
						return
					}
					if len(list) == 1 {
						productStock.Store(id, list[0].ProductId)
						err := model.PushToRedisStock(ctx, model.Product{
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
							klog.CtxErrorf(ctx, "获取库存信息时 推送到redis时异常,err: %v", err)
							return
						}
					} else if len(list) == 0 {
						klog.CtxErrorf(ctx, "从数据库中获取数据失败, 商品id: %d 不存在", id)
					} else {
						klog.CtxErrorf(ctx, "从数据库中获取数据失败, 商品id: %d 存在多个", id)
					}
					err = model.SafeDeleteLock(ctx, redisClient.RedisClient, lockKey, uuidString)
					if err != nil {
						klog.CtxInfof(ctx, "删除锁失败, err: %v", err)
					}
				}
			}
		}(&wg, ctx, &synMap)
	}
	wg.Wait()
	synMap.Range(func(key, value interface{}) bool {
		productStock[key.(int64)] = value.(int64)
		return true
	})
	return productStock, nil
}
