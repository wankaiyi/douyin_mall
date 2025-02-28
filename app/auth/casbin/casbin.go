package casbin

import (
	"context"
	myredis "douyin_mall/auth/biz/dal/redis"
	"douyin_mall/auth/conf"
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

var Enforcer *casbin.Enforcer

func InitCasbin() {
	dsn := conf.GetConf().Mysql.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		klog.Errorf("连接数据库失败: %v", err)
		panic(err)
	}
	if err = db.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		panic(err)
	}
	if conf.GetConf().Env == "dev" {
		db = db.Debug()
	}

	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		klog.Errorf("创建casbin适配器失败: %v", err)
		panic(err)
	}

	Enforcer, err = casbin.NewEnforcer(conf.GetConf().Casbin.ModelPath, adapter)
	if err != nil {
		klog.Errorf("创建casbin执行器失败: %v", err)
		panic(err)
	}

	printAllPolicies()

	subscribeToRedisChannel(myredis.RedisClient, context.Background())

}

func subscribeToRedisChannel(redisClient *redis.Client, ctx context.Context) {
	pubsub := redisClient.Subscribe(ctx, "casbin_policy_updates")
	go func() {
		defer pubsub.Close()
		for {
			msg, err := pubsub.ReceiveMessage(ctx)
			if err != nil {
				klog.Errorf("接收 Redis 订阅消息失败: %v", err)
				continue
			}

			klog.Infof("接收到 Redis 订阅消息: %s", msg.Payload)

			// 重新加载 Casbin 权限
			err = Enforcer.LoadPolicy()
			if err != nil {
				klog.Errorf("重新加载 Casbin 权限失败: %v", err)
			} else {
				klog.Info("成功重新加载 Casbin 权限")
				printAllPolicies()
			}
		}
	}()

}

func printAllPolicies() {
	policies, _ := Enforcer.GetFilteredPolicy(0) // 0 表示不过滤任何字段
	fmt.Println("当前所有权限策略：")
	for _, policy := range policies {
		fmt.Printf("%v\n", policy)
	}
}

func AddPolicy(sub, obj, act string) error {
	_, err := Enforcer.AddPolicy(sub, obj, act)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
