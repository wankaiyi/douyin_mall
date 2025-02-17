package utils

import (
	"context"
	"douyin_mall/common/infra/hot_key_client/model/base"
	"douyin_mall/common/infra/hot_key_client/model/key"
	"github.com/cloudwego/kitex/pkg/klog"
	"sync"
)

// TurnKeyCollector 用两个map轮流存储待测key
type TurnKeyCollector struct {
	Map0  sync.Map
	Map1  sync.Map
	Count base.AtomicCount
}

func NewTurnKeyCollector() *TurnKeyCollector {
	return &TurnKeyCollector{
		Map0:  sync.Map{},
		Map1:  sync.Map{},
		Count: base.AtomicCount{},
	}
}

func (t *TurnKeyCollector) LockAndGetResult() []key.HotkeyModel {
	t.Count.Add(1)
	var result []key.HotkeyModel
	if t.Count.GetCount()%2 == 0 {
		result = get(&t.Map0)
		t.Map0.Clear()
	} else {
		result = get(&t.Map1)
		t.Map1.Clear()
	}
	return result
}
func get(mapper *sync.Map) []key.HotkeyModel {
	var result []key.HotkeyModel
	mapper.Range(func(k, v interface{}) bool {
		klog.CtxInfof(context.Background(), "key:%v, value:%v", k, v)
		result = append(result, v.(key.HotkeyModel))
		return true
	})
	return result
}
func (t *TurnKeyCollector) Collect(keyModel key.HotkeyModel) {
	k := keyModel.Key
	if t.Count.GetCount()%2 == 0 {

		if value, ok := t.Map0.Load(k); ok {
			keyMod := value.(key.HotkeyModel)
			keyMod.Count.Add(1)
			t.Map0.Store(k, keyMod)
			klog.CtxInfof(context.Background(), "Store key:%v, count:%v", k, keyMod.Count)
		} else {
			t.Map0.Store(k, keyModel)
		}

	} else {

		if value, ok := t.Map1.Load(k); ok {
			keyMod := value.(key.HotkeyModel)
			keyMod.Count.Add(1)
			t.Map1.Store(k, keyMod)
		} else {
			t.Map0.Store(k, keyModel)
		}
	}
}
