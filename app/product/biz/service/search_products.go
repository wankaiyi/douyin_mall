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
	klog.CtxInfof(s.ctx, "开始搜索,参数为: dsl:%v", queryBody)
	searchIds, err := getSearchIds(s.ctx, dslBytes, md5Bytes)
	klog.CtxInfof(s.ctx, "搜索完成,dsl:%v,搜索结果为%v", queryBody, searchIds)

	var products = make([]*product.Product, 0)
	var missingIds []int64

	klog.CtxInfof(s.ctx, "开始获取缓存,参数:searchIds:%v,md5Bytes:%v", searchIds, md5Bytes)
	products, missingIds, err = GetCache(s.ctx, searchIds, md5Bytes)
	klog.CtxInfof(s.ctx, "获取缓存结果,products: %v,缺失的商品数据:%v", products, missingIds)
	if err != nil {
		klog.CtxErrorf(s.ctx, "获取缓存信息失败, err: %v", err)
		return nil, err
	}

	klog.CtxInfof(s.ctx, "开始获取缺失的商品数据")
	missingProduct, err := GetMissingProduct(s.ctx, missingIds)
	if err != nil {
		klog.CtxErrorf(s.ctx, "获取缺失的商品数据失败, err:%v", err)
		return nil, err
	}
	klog.CtxInfof(s.ctx, "获取缺失的商品数据成功, missingProduct: %v", missingProduct)

	products = append(products, missingProduct...)
	klog.CtxInfof(s.ctx, "组装的的products: %v", products)

	//根据商品id查询库存信息
	klog.CtxInfof(s.ctx, "开始获取商品的库存信息")
	productStock, err := GetStock(s.ctx, searchIds)
	if err != nil {
		klog.CtxErrorf(s.ctx, "获取库存时, err: %v", err)
		return nil, err
	}
	klog.CtxInfof(s.ctx, "获取商品的库存信息成功")
	for _, pro := range products {
		pro.Stock = productStock[pro.Id]
	}

	klog.CtxInfof(s.ctx, "最终返回的products: %v", products)
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
	DetectKey := "search:id"
	var ids = make([]int64, 0)
	dslKey := "product:dslBytes:" + string(md5bytes)
	keyConf := keyModel.NewKeyConf1(constant.ProductService)
	hotkeyModel := keyModel.NewHotKeyModelWithConfig(dslKey, &keyConf)
	localCache := processor.GetValue(*hotkeyModel)
	var dslCache string
	if localCache == nil || localCache.(string) == "" || localCache.(string) == "null" {
		redisCache, err := redisClient.RedisClient.Get(ctx, dslKey).Result()
		if err == nil {
			dslCache = redisCache
		}
		klog.CtxInfof(ctx, "命中redis缓存: key:%v,value:%v", DetectKey, redisCache)
		go addRedisHit(ctx, DetectKey)
	} else {
		go addLocalHit(ctx, DetectKey)
		klog.CtxInfof(ctx, "命中本地缓存: key:%v,value:%v", DetectKey, localCache)
		dslCache = localCache.(string)
	}
	if dslCache != "" && dslCache != "null" {
		err := sonic.UnmarshalString(dslCache, &ids)
		if err != nil {
			klog.CtxErrorf(ctx, "dsl缓存反序列化失败, err: %v", err)
			return ids, err
		}
	} else {
		klog.CtxInfof(ctx, "未命中缓存, key:%v", DetectKey)
		go notHit(ctx, DetectKey)
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

func GetCache(ctx context.Context, searchIds []int64, md5Bytes []byte) (products []*product.Product, missingIds []int64, err error) {
	detectKey := "search:base"
	//加入hotkey
	products = make([]*product.Product, 0)
	missingIds = make([]int64, 0)
	if len(searchIds) == 0 {
		return products, missingIds, nil
	}
	var keysKey = searchHotkey(string(md5Bytes))
	keysConf := keyModel.KeyConf{
		ServiceName: constant.ProductService,
		Threshold:   10,
		Interval:    10,
		Duration:    10,
	}
	keysHotkeyModel := keyModel.NewHotKeyModelWithConfig(keysKey, &keysConf)
	var values = make(map[int64]map[string]string)
	localKeysCache := processor.GetValue(*keysHotkeyModel)
	if localKeysCache != nil {
		go addLocalHit(ctx, detectKey)
		klog.CtxInfof(ctx, "本地缓存命中,key:%v,value: %v", detectKey, localKeysCache)
		err := sonic.UnmarshalString(localKeysCache.(string), &values)
		if err != nil {
			klog.CtxErrorf(ctx, "本地缓存反序列化缓存失败, err: %v", err)
			return nil, nil, err
		}
		processor.SmartSet(*keysHotkeyModel, localKeysCache)
	} else {
		for _, id := range searchIds {
			productKey := model.BaseInfoKey(ctx, id)
			result, _ := redisClient.RedisClient.HGetAll(ctx, productKey).Result()
			klog.CtxInfof(ctx, "id:%v,productKey:%v执行hgetall命令获取的结果为:%v", id, productKey, result)
			if len(result) != 0 {
				values[id] = result
			} else {
				missingIds = append(missingIds, id)
			}
		}
		marshal, err := sonic.MarshalString(values)
		if err != nil {
			klog.CtxErrorf(ctx, "values：%v,redis序列化缓存失败, err: %v", values, err)
			return nil, nil, err
		}
		processor.SmartSet(*keysHotkeyModel, marshal)
	}
	for _, id := range searchIds {
		if values[id] == nil {
			klog.CtxInfof(ctx, "未命中缓存,key:%v", detectKey)
			go notHit(ctx, "search:base")
			missingIds = append(missingIds, id)
		} else {
			//解析数据
			valueMap := values[id]
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
			products = append(products, &productData)
		}
	}
	return products, missingIds, nil
}

func GetMissingProduct(ctx context.Context, missingIds []int64) (products []*product.Product, err error) {
	klog.CtxInfof(ctx, "开始获取缺失的商品数据,参数:missingIds:%v", missingIds)
	products = make([]*product.Product, 0)
	if len(missingIds) > 0 {

		//从数据库中获取数据
		klog.CtxInfof(ctx, "开始从数据库中获取数据，根据missingIds:%v", missingIds)
		list, err := model.SelectProductList(mysql.DB, context.Background(), missingIds)
		if err != nil {
			klog.CtxErrorf(ctx, "从数据库中获取数据失败, err: %v", err)
			return nil, errors.WithStack(err)
		}

		klog.CtxInfof(ctx, "根据缺失的id列表查询数据库,结果:%v", list)
		for i := range list {

			func(ctx context.Context, pro *model.ProductWithCategory) {
				uuidString := uuid.New().String()
				lockKey := model.BaseInfoLockKey(ctx, pro.ProductId)
				klog.CtxInfof(ctx, "请求分布式锁,key:%v,uuid:%v", lockKey, uuidString)
				lock, err := model.SetLock(ctx, redisClient.RedisClient, lockKey, uuidString)
				if err != nil {
					klog.CtxInfof(ctx, "key %v 上锁失败", lockKey)
					return
				}

				klog.CtxInfof(ctx, "key %v 上锁状态:%v", lockKey, lock)
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
					products = append(products, &p)
					productKey := model.BaseInfoKey(ctx, list[i].ProductId)
					err = model.PushToRedisBaseInfo(ctx, model.Product{
						Base: model.Base{
							ID: pro.ProductId,
						},
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
					} else {
						klog.CtxInfof(ctx, "product数据缓存到redis成功")
					}
					err := model.SafeDeleteLock(ctx, redisClient.RedisClient, lockKey, uuidString)
					if err != nil {
						klog.CtxInfof(ctx, "删除锁失败")
						return
					} else {
						klog.CtxInfof(ctx, "安全删除锁成功,key:%v", lockKey)
					}
				}

			}(ctx, &list[i])
		}
	}
	return products, nil
}

func GetStock(ctx context.Context, searchIds []int64) (productStock map[int64]int64, err error) {
	detectKey := "search:stock"
	productStock = make(map[int64]int64)

	for _, id := range searchIds {
		//库存的key
		stockKey := model.StockKey(ctx, id)
		klog.CtxInfof(ctx, "id:%v,stockKey:%v", id, stockKey)
		//判断是否存在
		exists, err := redisClient.RedisClient.Exists(ctx, stockKey).Result()
		if err != nil {
			klog.CtxErrorf(ctx, "获取库存信息时判断key :%v , err: %v", stockKey, err)
			return nil, errors.WithStack(err)
		}

		klog.CtxInfof(ctx, "stockKey是否存在:%v", exists == 1)
		//如果存在，则从redis中获取数据
		if exists == 1 {
			klog.CtxInfof(ctx, "命中redis缓存,key:%v", stockKey)
			go addRedisHit(ctx, detectKey)
			//连续三次尝试
			maxTry := 3
			for i := 0; i < maxTry; i++ {
				result, err := redisClient.RedisClient.HGetAll(ctx, stockKey).Result()
				klog.CtxInfof(ctx, "result:%v,err:%v", result, err)
				if err != nil {
					if i == maxTry-1 {
						klog.CtxErrorf(ctx, "获取库存信息时超过最大重试次数%v key:%v, err: %v", maxTry, stockKey, err)
						return nil, errors.WithStack(err)
					}
				} else {
					stock, err := strconv.ParseInt(result["stock"], 10, 64)
					if err != nil {
						klog.CtxErrorf(ctx, "获取库存信息时 库存由string转int64异常,stock:%v ,stock:%v,err: %v", result["stock"], stockKey, err)
					}
					klog.CtxInfof(ctx, "成功获取redis上的stock:%v", stock)
					productStock[id] = stock
					break
				}
			}
		} else {
			klog.CtxInfof(ctx, "未命中缓存,key:%v", detectKey)
			go notHit(ctx, detectKey)
			//不存在
			//则先加锁
			lockKey := model.StockLockKey(ctx, id)
			uuidString := uuid.New().String()
			klog.CtxInfof(ctx, "lockKey:%v,uuidString:%v", lockKey, uuidString)
			result, err := redisClient.RedisClient.SetNX(ctx, lockKey, uuidString, 2*time.Second).Result()

			//如果加锁失败
			if err != nil || result == false {
				klog.CtxInfof(ctx, "lockKey:%v,uuidString:%v加锁失败", lockKey, uuidString)
				//连续三次尝试
				maxTry := 3
				for i := 0; i < maxTry; i++ {
					result, err := redisClient.RedisClient.HGetAll(ctx, stockKey).Result()
					if err != nil {
						if i == maxTry-1 {
							return nil, errors.WithStack(err)
						}
					} else {
						klog.CtxInfof(ctx, "redis中获取库存信息成功,key:%v,result:%v", stockKey, result)
						stock, err := strconv.ParseInt(result["stock"], 10, 64)
						if err != nil {
							klog.CtxErrorf(ctx, "获取库存信息时 库存由string转int64异常,stock:%v ,stock:%v,err: %v", result["stock"], stockKey, err)
						}
						productStock[id] = stock
						break
					}
				}
			} else {
				klog.CtxInfof(ctx, "lockKey:%v,uuidString:%v加锁成功", lockKey, uuidString)
				//如果加锁成功，则从数据库中获取数据
				list, err := model.SelectProductList(mysql.DB, ctx, []int64{id})
				if err != nil {
					klog.CtxErrorf(ctx, "获取库存信息时 从数据库中获取数据失败, err: %v", err)
					return nil, errors.WithStack(err)
				}
				klog.CtxInfof(ctx, "数据库中的库存数据:%v", list)
				if len(list) == 1 {
					productStock[id] = list[0].ProductStock
					err = model.PushToRedisStock(ctx, model.Product{
						Base: model.Base{
							ID: list[0].ProductId,
						},
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
						return nil, errors.WithStack(err)
					}
					klog.CtxInfof(ctx, "库存信息推送成功")
				} else if len(list) == 0 {
					klog.CtxErrorf(ctx, "从数据库中获取数据失败, 商品id: %d 不存在", id)
				} else {
					klog.CtxErrorf(ctx, "从数据库中获取数据失败, 商品id: %d 存在多个", id)
				}
				err = model.SafeDeleteLock(ctx, redisClient.RedisClient, lockKey, uuidString)
				if err != nil {
					klog.CtxInfof(ctx, "删除锁失败,lockKey:%v,uuidString:%v, err: %v", lockKey, uuidString, err)
				}
				klog.CtxInfof(ctx, "删除锁成功,lockKey:%v,uuidString:%v", lockKey, uuidString)
			}
		}
	}
	return productStock, nil
}

func addRedisHit(ctx context.Context, key string) {
	luaScript := `
local key = KEYS[1]
      redis.call('HINCRBY', key,'total_queries',1)
redis.call('HINCRBY', key,'redis_cache_hit',1)
`
	redisClient.RedisClient.Eval(ctx, luaScript, []string{key})
}
func notHit(ctx context.Context, key string) {
	luaScript := `
local key = KEYS[1]
redis.call('HINCRBY', key,'total_queries',1)
`
	redisClient.RedisClient.Eval(ctx, luaScript, []string{key})
}

func addLocalHit(ctx context.Context, key string) {
	luaScript := `
local key = KEYS[1]
redis.call('HINCRBY', key,'total_queries',1)
redis.call('HINCRBY', key,'local_cache_hit',1)
`
	redisClient.RedisClient.Eval(ctx, luaScript, []string{key})
}
