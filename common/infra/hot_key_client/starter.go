package hotKeyClient

import (
	"douyin_mall/common/infra/hot_key_client/listener"
	"douyin_mall/common/infra/hot_key_client/publisher"
	cientRedis "douyin_mall/common/infra/hot_key_client/redis"

	"github.com/redis/go-redis/v9"
)

// Start 启动hotkey客户端
func Start(redisClient *redis.Client) {

	cientRedis.Init(redisClient)
	go publisher.PublishStarter()
	go listener.ListenStarter()

}
