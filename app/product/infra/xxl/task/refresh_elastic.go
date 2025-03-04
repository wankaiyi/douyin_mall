package task

import (
	"bytes"
	"context"
	"douyin_mall/product/biz/dal/mysql"
	redisClient "douyin_mall/product/biz/dal/redis"
	"douyin_mall/product/biz/model"
	"douyin_mall/product/biz/vo"
	"douyin_mall/product/infra/elastic/client"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/pkg/errors"
	"github.com/xxl-job/xxl-job-executor-go"
	"io"
	"strings"
	"time"
)

func RefreshElastic(cxt context.Context, param *xxl.RunReq) string {
	index := param.BroadcastIndex
	total := param.BroadcastTotal

	klog.CtxInfof(cxt, "刷新Elastic开始 CheckAccountTask start")
	err := refresh(cxt, index, total)
	if err != nil {
		klog.Errorf("刷新Elastic失败 CheckAccountTask failed, err: %v", err)
		return err.Error()
	}
	return "刷新Elastic成功"
}

func refresh(ctx context.Context, index int64, total int64) (err error) {
	lockKey := "product_refresh_lock"
	//先创建索引库
	indexName := fmt.Sprintf("product_v%d", time.Now().Unix())
	klog.CtxInfof(ctx, "预期创建索引库%v", indexName)
	err = createIndex(ctx, lockKey, indexName)
	if err != nil {
		return err
	}
	//睡眠十秒
	time.Sleep(10 * time.Second)
	//从redis获取索引名称
	indexName, err = redisClient.RedisClient.Get(ctx, lockKey).Result()
	if err != nil {
		klog.CtxErrorf(ctx, "从数据库获取es新索引名失败,err:%v", err)
		return err
	}
	klog.CtxInfof(ctx, "获取到索引库%v", indexName)
	err = importDataToIndex(ctx, indexName, index, total)
	if err != nil {
		klog.CtxErrorf(ctx, "导入数据到索引库失败,err:%v", err)
		return err
	}
	klog.CtxInfof(ctx, "导入数据到索引库成功")
	klog.CtxInfof(ctx, "设置索引库别名")
	err = alias(ctx, indexName, total)
	if err != nil {
		klog.CtxErrorf(ctx, "设置索引库别名失败,err:%v", err)
		return err
	}
	return nil
}

func setAlias(ctx context.Context, indexName string) (err error) {
	aliasBody := map[string]interface{}{
		"actions": []map[string]interface{}{
			{
				"add": map[string]interface{}{
					"index": indexName,
					"alias": "product",
				},
			},
		},
	}
	aliasBytes, err := sonic.Marshal(aliasBody)
	if err != nil {
		klog.CtxErrorf(ctx, "序列化设置别名的配置失败,err:%v", err)
		return err
	}
	aliasUpdate, err := esapi.IndicesUpdateAliasesRequest{
		Body: bytes.NewReader(aliasBytes),
	}.Do(ctx, client.ElasticClient)
	if err != nil {
		klog.CtxErrorf(ctx, "更新%v索引库别名失败,err:%v", indexName, err)
		return err
	}
	if aliasUpdate.StatusCode == 200 {
		klog.CtxInfof(ctx, "更新%v索引库别名成功", indexName)
	}
	return nil
}

func unlinkAlias(ctx context.Context, indexName string) (err error) {
	aliasBody := map[string]interface{}{
		"actions": []map[string]interface{}{
			{
				"remove": map[string]interface{}{
					"index": indexName,
					"alias": "product",
				},
			},
		},
	}
	aliasBytes, err := sonic.Marshal(aliasBody)
	if err != nil {
		klog.CtxErrorf(ctx, "序列化删除别名的配置失败,err:%v", err)
		return err
	}
	aliasUpdate, err := esapi.IndicesUpdateAliasesRequest{
		Body: bytes.NewReader(aliasBytes),
	}.Do(ctx, client.ElasticClient)
	if err != nil {
		klog.CtxErrorf(ctx, "删除%v索引库的别名失败,err:%v", indexName, err)
		return err
	}
	if aliasUpdate.StatusCode != 200 {
		klog.CtxErrorf(ctx, "删除%v索引库的别名失败,statusCode:%v", indexName, aliasUpdate.Body)
		return errors.New("删除索引库的别名失败")
	}
	klog.CtxInfof(ctx, "删除%v索引库的别名成功", indexName)
	return nil
}

func alias(ctx context.Context, indexName string, total int64) (err error) {
	luaScript := `
		local key = KEYS[1]
		local total = ARGV[1]

		if redis.call("EXISTS", key) == 0 then
			redis.call("SET", key, total)
		end
		local result = redis.call("DECRBY", key)
		return result
`
	keys := []string{indexName}
	args := []interface{}{total}
	result, err := redisClient.RedisClient.Eval(ctx, luaScript, keys, args).Result()
	if err != nil {
		klog.CtxErrorf(ctx, "获取分片结果值失败,key:%v,err: %v", indexName, err)
		return err
	}
	//如果是0，说明已经全部导入完成
	if result.(int64) == 0 {
		//为表设置别名
		aliasGet, err := esapi.IndicesGetAliasRequest{
			Index: []string{"product"},
		}.Do(ctx, client.ElasticClient)
		if err != nil {
			klog.CtxErrorf(ctx, "获取product索引库别名失败,err:%v", err)
			return err
		}
		//如果没有索引库的别名叫product，那么直接把新索引的别名改为product
		if aliasGet.StatusCode == 404 {
			err := setAlias(ctx, indexName)
			if err != nil {
				klog.CtxErrorf(ctx, "设置%v索引库别名失败,err:%v", indexName, err)
				return err
			}
		} else {
			bodyBytes, err := io.ReadAll(aliasGet.Body)
			if err != nil {
				klog.CtxErrorf(ctx, "读取product指向的索引库失败,err:%v", err)
				return err
			}
			aliasMap := map[string]interface{}{}
			err = sonic.Unmarshal(bodyBytes, &aliasMap)
			if err != nil {
				klog.CtxErrorf(ctx, "反序列化product指向的索引库失败,err:%v", err)
				return err
			}
			for index, _ := range aliasMap {
				if index != "product" {
					err := unlinkAlias(ctx, index)
					if err != nil {
						klog.CtxErrorf(ctx, "删除%v索引库别名失败,err:%v", index, err)
						return err
					}
				}
			}
			err = setAlias(ctx, indexName)
			if err != nil {
				klog.CtxErrorf(ctx, "设置%v索引库别名失败,err:%v", indexName, err)
				return err
			}
		}
	}
	return nil
}

func createIndex(ctx context.Context, lockKey string, indexName string) (err error) {
	luaScript := `
		local key = KEYS[1]
		local index = ARGV[1]
		local result = redis.call("SETNX", key,index)
		redis.call("EXPIRE", key, 300)
		return result
	`
	keys := []string{lockKey}
	args := []interface{}{indexName}
	result, err := redisClient.RedisClient.Eval(ctx, luaScript, keys, args).Result()
	if err != nil {
		klog.CtxErrorf(ctx, "获取锁失败,锁key:%v,err: %v", lockKey, err)
		return err
	}
	//如果result为1，那么需要创建新索引库product_V{timestamp}
	if result.(int64) != 1 {
		klog.CtxInfof(ctx, "创建新索引库 %s", indexName)
		marshal, err := sonic.Marshal(vo.ProductSearchMappingSetting)
		if err != nil {
			klog.CtxErrorf(ctx, "序列化新索引的配置失败,err:%v", err)
			return err
		}
		create, err := esapi.IndicesCreateRequest{
			Index: indexName,
			Body:  bytes.NewReader(marshal),
		}.Do(ctx, client.ElasticClient)
		if err != nil {
			return err
		}
		if create.StatusCode != 200 {
			klog.CtxErrorf(ctx, "新索引库创建失败,statusCode:%v", create.Body)
		}
	}
	return nil
}

func importDataToIndex(ctx context.Context, indexName string, index int64, total int64) (err error) {
	klog.CtxInfof(ctx, "开始导入数据到索引库%v,index:%v,total:%v", indexName, index, total)
	//从数据库获取分片数据
	allProduct, err := model.SelectProductAll(mysql.DB, ctx, index, total)
	if err != nil {
		klog.CtxErrorf(ctx, "从数据库获取分片数据失败,err:%v", err)
		return err
	}
	klog.CtxInfof(ctx, "获取到%v条数据", len(allProduct))
	productMap := map[int64]model.ProductWithCategory{}
	for i := range allProduct {
		productMap[allProduct[i].ProductId] = allProduct[i]
	}

	bulkBody := strings.Builder{}
	typeBody := map[string]map[string]string{
		"create": {
			"_index": indexName,
		},
	}
	typeBytes, err := sonic.Marshal(typeBody)
	if err != nil {
		return err
	}
	for _, pro := range allProduct {
		dataVo := vo.ProductSearchDataVo{
			Name:         pro.ProductName,
			Description:  pro.ProductDescription,
			ID:           pro.ProductId,
			CategoryName: pro.CategoryName,
			CategoryID:   pro.CategoryID,
		}
		voBytes, err := sonic.Marshal(dataVo)
		if err != nil {
			return err
		}
		bulkBody.Write(typeBytes)
		bulkBody.WriteString("\n")
		bulkBody.Write(voBytes)
		bulkBody.WriteString("\n")
	}
	klog.CtxInfof(ctx, "开始导入数据到索引库%v", indexName)
	klog.CtxInfof(ctx, "数据:%v", bulkBody.String())
	bulkResponse, err := esapi.BulkRequest{
		Index: indexName,
		Body:  strings.NewReader(bulkBody.String()),
	}.Do(ctx, client.ElasticClient)
	if err != nil {
		return err
	}
	if bulkResponse.StatusCode != 200 {
		klog.CtxErrorf(ctx, "导入数据到索引库失败,statusCode:%v", bulkResponse.Body)
		return errors.New("导入数据到索引库失败,es状态码不为200")
	}
	return nil
}
