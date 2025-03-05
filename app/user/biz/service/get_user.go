package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/user/biz/dal/mysql"
	"douyin_mall/user/biz/dal/redis"
	"douyin_mall/user/biz/model"
	user "douyin_mall/user/kitex_gen/user"
	redisUtils "douyin_mall/user/utils/redis"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

type GetUserService struct {
	ctx context.Context
} // NewGetUserService new GetUserService
func NewGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{ctx: ctx}
}

// Run create note info
func (s *GetUserService) Run(req *user.GetUserReq) (resp *user.GetUserResp, err error) {
	ctx := s.ctx
	userId := req.UserId
	redisKey := redisUtils.GetUserKey(userId)
	redisClient := redis.RedisClient

	redisClient.HIncrBy(ctx, redisUtils.GetUserInfoHitRateKey(), "visit", 1)
	cachedData, err := redisClient.HGetAll(ctx, redisKey).Result()
	if err == nil && len(cachedData) > 0 {
		klog.Infof("缓存命中用户信息: %v", cachedData)
		redisClient.HIncrBy(ctx, redisUtils.GetUserInfoHitRateKey(), "hit", 1)
		redisClient.Expire(ctx, redisKey, time.Hour*2)

		userInfo := &user.User{
			Username:    cachedData["username"],
			Email:       cachedData["email"],
			Sex:         cachedData["sex"],
			Description: cachedData["description"],
			Avatar:      cachedData["avatar"],
			CreatedAt:   cachedData["createdAt"],
		}

		return &user.GetUserResp{
			StatusCode: 0,
			StatusMsg:  "success",
			User:       userInfo,
		}, nil
	}

	u, err := s.SelectAndCacheUserInfo(ctx, userId)
	if err != nil {
		return nil, err
	}
	resp.User = &user.User{
		Username:    u.Username,
		Email:       u.Email,
		Sex:         model.SexToString(u.Sex),
		Description: u.Description,
		Avatar:      u.Avatar,
		CreatedAt:   u.CreatedAt,
	}
	resp.StatusCode = 0
	resp.StatusMsg = constant.GetMsg(0)
	return resp, nil
}

func (s *GetUserService) SelectAndCacheUserInfo(ctx context.Context, userId int32) (*model.UserBasicInfo, error) {
	u, err := model.GetUserBasicInfoById(mysql.DB, ctx, userId)
	if err != nil {
		return nil, err
	}
	err = model.CacheUserInfo(ctx, u, userId)
	if err != nil {
		return nil, err
	}
	return u, nil
}
