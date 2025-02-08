package red_sync

import (
	"douyin_mall/api/conf"
	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

var (
	redisAddr = conf.GetConf().Redsync.Addr
)

// GetRedsync 返回一个redsync 实例
func GetRedsync() (rsync *redsync.Redsync) {
	client := goredislib.NewClient(&goredislib.Options{
		Addr: redisAddr,
	})

	pool := goredis.NewPool(client)

	rsync = redsync.New(pool)
	return
}
