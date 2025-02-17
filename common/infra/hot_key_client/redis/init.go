package redis

import (
	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func Init(rdb *redis.Client) {
	Rdb = rdb
}
