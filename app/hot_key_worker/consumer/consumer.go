package consumer

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"hotkey/model/cache"
	"hotkey/model/key"
	"hotkey/model/util"
	"hotkey/redis"
	"hotkey/tool"
	"time"
)

var (
	hotKeyCache *bigcache.BigCache
)

func init() {
	hotKeyCache = cache.NewRecentKeyCache()
}

func Consume() {
	for {
		hotkeyModel := util.BlQueue.Take()
		if hotkeyModel.Remove {
			err := removeKey(hotkeyModel)
			if err != nil {
				klog.CtxErrorf(context.Background(), "remove key error:%s", err)
			}
		} else {

			err := newKey(hotkeyModel)
			if err != nil {
				klog.CtxErrorf(context.Background(), "new key error:%s", err)
			}
		}
	}

}

func removeKey(hotkeyModel key.HotkeyModel) (err error) {
	buildKey := BuildKey(hotkeyModel)
	//从热键缓存中删除
	err = hotKeyCache.Delete(buildKey)
	if err != nil {
		return err
	}
	//通知所有的client集群删除
	err = redis.PublishClientChannel(hotkeyModel)
	if err != nil {
		return err
	}
	return nil
}

func newKey(hotkeyModel key.HotkeyModel) (err error) {
	buildKey := BuildKey(hotkeyModel)
	//判断key是否刚热过
	_, err = hotKeyCache.Get(buildKey)
	if err == nil {
		return nil
	}
	//if exist != nil {
	//	return nil
	//}
	slidingWindow := getWindow(hotkeyModel, buildKey)
	hot := slidingWindow.AddCount(hotkeyModel.Count.GetCount())

	//不热放进所有键的缓存
	if !hot {
		marshal, _ := sonic.Marshal(slidingWindow)
		cache.GetAllKeyCache(hotkeyModel.ServiceName).Set(buildKey, marshal)
		return nil
	}
	//热放进热键的缓存
	hotKeyCache.Set(buildKey, []byte("hot"))
	hotkeyModel.CreateAt = time.Now().UnixMilli()
	err = redis.PublishClientChannel(hotkeyModel)
	if err != nil {
		return err
	}
	return nil

}

func BuildKey(hotkeyModel key.HotkeyModel) string {
	return hotkeyModel.ServiceName + "+" + hotkeyModel.Key
}

// 获取窗口
func getWindow(hotkeyModel key.HotkeyModel, key string) *tool.SlidingWindow {
	bigCache := cache.GetAllKeyCache(hotkeyModel.ServiceName)
	window, emptyError := bigCache.Get(key)
	if emptyError != nil {
		slideWindow := tool.NewSlidingWindow(hotkeyModel.Interval, hotkeyModel.Threshold)
		windowJson, _ := sonic.Marshal(slideWindow)
		bigCache.Set(key, windowJson)
		return slideWindow
	}
	var slideWindow tool.SlidingWindow
	sonic.Unmarshal(window, &slideWindow)
	return &slideWindow

}
