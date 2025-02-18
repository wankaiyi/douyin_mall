package listener

import (
	"context"
	"douyin_mall/common/infra/hot_key_client/constants"
	"douyin_mall/common/infra/hot_key_client/model/cache"
	"douyin_mall/common/infra/hot_key_client/model/key"
	"douyin_mall/common/infra/hot_key_client/redis"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/klog"

	"log"
)

func ListenStarter() {
	pubsub := redis.Rdb.Subscribe(context.Background(), constants.ClientChannel)
	ch := pubsub.Channel()

	for msg := range ch {
		klog.CtxInfof(context.Background(), "receive message: %s", msg.Payload)
		var hotKeyModel key.HotkeyModel
		sonic.Unmarshal([]byte(msg.Payload), &hotKeyModel)
		if hotKeyModel.ServiceName != constants.ClientServiceName {
			break
		}
		if hotKeyModel.Remove {
			removeKey(hotKeyModel.Key)
		}
		err := newKey(hotKeyModel)
		if err != nil {
			log.Printf(err.Error())
		}

	}
	defer pubsub.Close()
}

func newKey(hotKeyModel key.HotkeyModel) (err error) {
	hlog.CtxInfof(context.Background(), "receive new key: %s", hotKeyModel.Key)
	if cache.IsHot(hotKeyModel.Key) {
		log.Printf("repeat receive hot key, key: %s", hotKeyModel.Key)
	}
	return addKey(hotKeyModel, err)
}

func addKey(hotKeyModel key.HotkeyModel, err error) error {
	valueModel, ok := cache.LocalStore.GetDefaultValue(hotKeyModel.Key)
	if ok {
		// key already exist, update duration已存在则重置
		klog.CtxInfof(context.Background(), "key already exist,will update duration, key: %s", hotKeyModel.Key)
		valueModel.Duration = hotKeyModel.Duration
		err = cache.LocalStore.PutInLocalCacheStore(hotKeyModel.Key, valueModel)
		if err != nil {
			return err
		}
		return nil
	}
	valueModel.Duration = hotKeyModel.Duration
	err = cache.LocalStore.PutInLocalCacheStore(hotKeyModel.Key, valueModel)
	if err != nil {
		return err
	}
	return nil
}

// removeKey remove key from local cache
func removeKey(key string) {
	cache.LocalStore.Remove(key)
}
