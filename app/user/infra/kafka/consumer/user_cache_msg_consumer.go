package consumer

import (
	"context"
	"douyin_mall/user/biz/dal/mysql"
	"douyin_mall/user/biz/dal/redis"
	"douyin_mall/user/biz/model"
	"douyin_mall/user/conf"
	redisUtils "douyin_mall/user/utils/redis"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"time"
)

type msgConsumerGroup struct{}

func (msgConsumerGroup) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (msgConsumerGroup) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h msgConsumerGroup) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		klog.Infof("收到消息，topic:%q partition:%d offset:%d  value:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))

		userCacheMsg := UserCacheMessage{}
		err := sonic.Unmarshal(msg.Value, &userCacheMsg)
		if err != nil {
			klog.Errorf("反序列化消息失败，err：%v", err)
			return err
		}

		err = selectAndCacheUserInfo(sess.Context(), userCacheMsg.UserId)
		if err != nil {
			klog.Errorf("缓存用户信息失败，err：%v", err)
			return err
		}

		err = selectAndCacheUserAddresses(sess.Context(), userCacheMsg.UserId)
		if err != nil {
			klog.Errorf("缓存用户地址失败，err：%v", err)
			return err
		}

		sess.MarkMessage(msg, "")
		sess.Commit()
	}
	return nil
}

func selectAndCacheUserAddresses(ctx context.Context, userId int32) error {
	addresses, err := model.GetAddressList(ctx, mysql.DB, userId)
	if err != nil {
		return err
	}
	if len(addresses) == 0 {
		return nil
	}
	luaScript := `
		if redis.call('EXISTS', KEYS[1]) == 0 then
			return redis.call('RPUSH', KEYS[1], unpack(ARGV))
		else
			return 0
		end
	`
	key := redisUtils.GetUserAddressKey(userId)
	redisClient := redis.RedisClient
	addressStrs := make([]string, len(addresses))
	for i, address := range addresses {
		addressStr, err := sonic.Marshal(address)
		if err != nil {
			return err
		}
		addressStrs[i] = string(addressStr)
		//addressStrs = append(addressStrs, string(addressStr))
	}
	//bytes, err := sonic.Marshal(addresses)
	err = redisClient.Eval(ctx, luaScript, []string{key}, addressStrs).Err()
	if err != nil {
		return err
	}
	//err = redisClient.RPush(ctx, key, addressStrs).Err()
	//if err != nil {
	//	return err
	//}
	// 设置过期时间和access token的过期时间一致
	err = redisClient.Expire(ctx, key, time.Hour*2).Err()
	if err != nil {
		return err
	}
	return nil
}

func selectAndCacheUserInfo(ctx context.Context, userId int32) error {
	user, err := model.GetUserBasicInfoById(mysql.DB, ctx, userId)
	if err != nil {
		return err
	}
	key := redisUtils.GetUserKey(userId)
	redisClient := redis.RedisClient
	err = redisClient.HSet(ctx, key, user.ToMap()).Err()
	if err != nil {
		return err
	}
	// 设置过期时间和access token的过期时间一致
	err = redisClient.Expire(ctx, key, time.Hour*2).Err()
	if err != nil {
		return err
	}
	return nil
}

type UserCacheMessage struct {
	UserId int32 `json:"user_id"`
}

func InitUserCacheMessageConsumer() {
	consumerConfig := sarama.NewConfig()
	consumerConfig.Version = sarama.V3_5_0_0
	consumerConfig.Consumer.Offsets.AutoCommit.Enable = false
	consumerConfig.Consumer.Offsets.Initial = sarama.OffsetNewest
	consumerConfig.Consumer.Offsets.Retry.Max = 3

	gtoupId := "cache-user-info"
	if conf.GetConf().Env == "dev" {
		gtoupId += "-dev"
	}
	cGroup, err := sarama.NewConsumerGroup(conf.GetConf().Kafka.BizKafka.BootstrapServers, "cache-user-info-dev", consumerConfig)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			err = cGroup.Consume(context.Background(), []string{"auth_service_deliver_token"}, msgConsumerGroup{})
			if err != nil {
				fmt.Println(err.Error())
				break
			}
		}
	}()

	server.RegisterShutdownHook(func() {
		_ = cGroup.Close()
	})

}
