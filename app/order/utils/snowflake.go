package utils

import (
	"context"
	myredis "douyin_mall/order/biz/dal/redis"
	rediskeys "douyin_mall/order/utils/redis"
	"github.com/bwmarrin/snowflake"
)

var (
	node *snowflake.Node
	err  error
	ctx  = context.Background()
)

func InitSnowflake() {
	nodeId := allocateNodeID()
	node, err = snowflake.NewNode(nodeId)
	if err != nil {
		panic(err)
	}
}

func GetSnowflakeID() string {
	return node.Generate().String()
}

func allocateNodeID() int64 {
	nodeId, err := myredis.RedisClient.Incr(ctx, rediskeys.OrderServiceNodeIdKey).Result()
	if err != nil {
		panic(err)
	}
	return nodeId
}
