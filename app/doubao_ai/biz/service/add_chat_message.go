package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/doubao_ai/biz/dal/mysql"
	"douyin_mall/doubao_ai/biz/dal/redis"
	"douyin_mall/doubao_ai/biz/model"
	doubao_ai "douyin_mall/doubao_ai/kitex_gen/doubao_ai"
	redisUtils "douyin_mall/doubao_ai/utils/redis"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
)

type AddChatMessageService struct {
	ctx context.Context
} // NewAddChatMessageService new AddChatMessageService
func NewAddChatMessageService(ctx context.Context) *AddChatMessageService {
	return &AddChatMessageService{ctx: ctx}
}

// Run create note info
func (s *AddChatMessageService) Run(req *doubao_ai.AddChatMessageReq) (resp *doubao_ai.AddChatMessageResp, err error) {
	ctx := s.ctx
	userId := req.UserId
	role := req.Role
	content := req.Content
	uuid := req.Uuid
	scenario := req.Scenario
	aiMessage := &model.Message{
		UserId:   userId,
		Role:     model.Role(role),
		Content:  content,
		Uuid:     uuid,
		Scenario: model.Scenario(scenario),
	}
	err = model.CreateMessage(mysql.DB, ctx, aiMessage)
	if err != nil {
		klog.CtxErrorf(ctx, "智能助手：对话消息写入数据库失败，message: %s, err: %v", aiMessage, err)
		return nil, errors.WithStack(err)
	}
	jsonStrBytes, _ := sonic.Marshal(aiMessage)
	redis.RedisClient.RPush(ctx, redisUtils.GetChatHistoryKey(uuid), string(jsonStrBytes))
	return &doubao_ai.AddChatMessageResp{StatusCode: 0, StatusMsg: constant.GetMsg(0)}, nil
}
