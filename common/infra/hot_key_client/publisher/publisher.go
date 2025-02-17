package publisher

import (
	"context"
	"douyin_mall/common/infra/hot_key_client/constants"
	"douyin_mall/common/infra/hot_key_client/model/holder"
	"douyin_mall/common/infra/hot_key_client/model/key"
	"douyin_mall/common/infra/hot_key_client/redis"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/spaolacci/murmur3"

	"time"
)

func PublishStarter() {
	//每隔500毫秒发布一次热点key
	rateChannel := time.Tick(500 * time.Millisecond)
	for {
		<-rateChannel
		err := publish()
		if err != nil {
			klog.CtxErrorf(context.Background(), "publish hotkey failed:%s", err)
		}

	}

}

func publish() error {
	hotkeyModels := holder.TurnKeyCollector.LockAndGetResult()
	for _, hotkeyModel := range hotkeyModels {
		go func() {
			err := push(hotkeyModel)
			if err != nil {
				klog.CtxErrorf(context.Background(), "push hotkey failed:%s", err)
			}
		}()
	}

	return nil
}

func push(hotkeyModel key.HotkeyModel) error {
	hlog.CtxInfof(context.Background(), "publish hotkey:%s", hotkeyModel.Key)
	marshal, _ := sonic.Marshal(hotkeyModel)
	err := redis.Rdb.Publish(context.Background(), getWorkerChannelId(hotkeyModel.Key), marshal).Err()
	return err
}

func getWorkerChannelId(key string) string {
	length := redis.Rdb.LLen(context.Background(), constants.WorkerChannelIdList)
	if length.Val() == 0 {
		klog.CtxErrorf(context.Background(), "worker channel id list is empty, please starter worker first!")
		return ""
	}
	sum32 := murmur3.Sum32([]byte(key))
	workerIdIndex := sum32 % uint32(length.Val())

	workerChannelId := redis.Rdb.LRange(context.Background(), constants.WorkerChannelIdList, int64(workerIdIndex), int64(workerIdIndex))
	return workerChannelId.Val()[0]
}
