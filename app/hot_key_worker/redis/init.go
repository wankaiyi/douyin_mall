package redis

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/redis/go-redis/v9"
	"hotkey/conf"
	"hotkey/constant"
	"hotkey/model/key"
)

var Rdb *redis.Client

func Init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     conf.GetConf().Redis.Address,
		Username: conf.GetConf().Redis.Username,
		Password: conf.GetConf().Redis.Password,
		DB:       conf.GetConf().Redis.DB, // 默认DB 0
	})
}

func PublishClientChannel(hotKeyModel key.HotkeyModel) (err error) {
	marshal, _ := sonic.Marshal(hotKeyModel)
	err = Rdb.Publish(context.Background(), constant.ClientChannel, marshal).Err()
	if err != nil {
		return err
	}
	return nil
}
