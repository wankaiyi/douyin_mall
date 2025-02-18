package hotKeyClient

import (
	"douyin_mall/common/infra/hot_key_client/constants"
	"douyin_mall/common/infra/hot_key_client/listener"
	"douyin_mall/common/infra/hot_key_client/publisher"
	cientRedis "douyin_mall/common/infra/hot_key_client/redis"

	"github.com/redis/go-redis/v9"
)

// Start 启动hotkey客户端,参数分别是redis客户端和服务名称
func Start(redisClient *redis.Client, serviceName string) {

	cientRedis.Init(redisClient)
	constants.Init(serviceName)
	go publisher.PublishStarter()
	go listener.ListenStarter()

}
