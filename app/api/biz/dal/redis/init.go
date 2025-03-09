package redis

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"io"
	"os"

	"douyin_mall/api/conf"
	"github.com/redis/go-redis/v9"
)

var (
	RedisClient    *redis.Client
	LimitScriptSha string
)

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     conf.GetConf().Redis.Address,
		Username: conf.GetConf().Redis.Username,
		Password: conf.GetConf().Redis.Password,
		DB:       conf.GetConf().Redis.DB,
	})
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
	if err := initScript(); err != nil {
		hlog.CtxErrorf(context.Background(), "初始化lua脚本失败: %v", err)
	}
}

func initScript() (err error) {
	open, err := os.Open(conf.GetConf().LuaScript.LimitScript)
	if err != nil {
		return err
	}

	defer open.Close()
	script, err := io.ReadAll(open)
	if err != nil {
		return err
	}

	LimitScriptSha, err = RedisClient.ScriptLoad(context.Background(), string(script)).Result()
	if err != nil {
		return err
	}
	return nil
}
