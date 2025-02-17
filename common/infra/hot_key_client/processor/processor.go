package processor

import (
	"context"
	"douyin_mall/common/infra/hot_key_client/model/cache"
	"douyin_mall/common/infra/hot_key_client/model/holder"
	keyModel "douyin_mall/common/infra/hot_key_client/model/key"
	"douyin_mall/common/infra/hot_key_client/model/value"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

// IsHotKey 检查是否热点,不是则上报
func IsHotKey(hotKeyModel keyModel.HotkeyModel) bool {
	isHot := cache.IsHot(hotKeyModel.Key)
	if !isHot {
		//收集key数据并发送

		push(hotKeyModel, 1)
	} else {
		valueModel, _ := cache.LocalStore.GetDefaultValue(hotKeyModel.Key)
		if isNearExpire(valueModel) {
			// 临近过期也发送
			push(hotKeyModel, 1)
		}
	}
	return isHot
}

// SmartSet 判断是否是热key，是则给value赋值
func SmartSet(hotKeyModel keyModel.HotkeyModel, value any) {
	if IsHotKey(hotKeyModel) {
		valueModel, _ := cache.LocalStore.GetDefaultValue(hotKeyModel.Key)
		valueModel.Value = value
		err := cache.LocalStore.PutInLocalCacheStore(hotKeyModel.Key, valueModel)
		if err != nil {
			klog.CtxErrorf(context.Background(), "put hotkey error: %v ,key: %s", err, hotKeyModel.Key)
		}
		return

	}
}

// SetDirectly 设置本地缓存，不做任何判断
func SetDirectly(key string, valueModel value.ValueModel) (err error) {
	err = cache.LocalStore.PutInLocalCacheStore(key, valueModel)
	return
}

// Get 获取本地换出key的value(无论是否热key)
func Get(key string) any {
	hotKeyValueModel, ok := cache.LocalStore.GetDefaultValue(key)
	if !ok {
		//非热key直接返回nil
		return nil
	}
	//热key但没有设置value则返回nil
	return hotKeyValueModel.Value

}

// GetValue 相当于isHotKey和get两个方法的整合，该方法直接返回本地缓存的value。
// 如果是热key，则存在两种情况，1是返回value，2是返回null。返回null是因为尚未给它set真正的value，
// 返回非null说明已经调用过set方法了，本地缓存value有值了。
// 如果不是热key，则返回null，并且将key上报到探测集群进行数量探测
func GetValue(hotKeyModel keyModel.HotkeyModel) any {
	hotKeyValueModel, ok := cache.LocalStore.GetDefaultValue(hotKeyModel.Key)
	if !ok {

		push(hotKeyModel, 1)
		return nil
	}
	return hotKeyValueModel.Value

}

// 检查是否临近过期
func isNearExpire(valueModel value.ValueModel) bool {
	return (valueModel.CreatedAt+valueModel.Duration)-time.Now().UnixMilli() <= 2000

}

// Remove 从本地缓存中删除key,并且通知整个集群删除该key
func Remove(key string) {
	//删除本地缓存
	cache.LocalStore.Remove(key)
	//通知集群删除
	hotkeyModel := keyModel.NewDefaultHotkeyModel1(key)
	push(*hotkeyModel, -1)

}
func push(hotKeyModel keyModel.HotkeyModel, count int) {
	if count <= 0 {
		count = 1
	}
	hotKeyModel.Count.Add(int64(count))
	// 收集积攒发送
	holder.TurnKeyCollector.Collect(hotKeyModel)

}
