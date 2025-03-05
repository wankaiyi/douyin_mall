package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/common/utils"
	"douyin_mall/doubao_ai/biz/dal/mysql"
	"douyin_mall/doubao_ai/biz/dal/redis"
	"douyin_mall/doubao_ai/biz/model"
	"douyin_mall/doubao_ai/conf"
	redisUtils "douyin_mall/doubao_ai/utils/redis"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type BaseAiService struct {
	ctx context.Context
	db  *gorm.DB
}

type transactionParams struct {
	UserId          int32
	Question        string
	Uuid            string
	ChatHistory     []*schema.Message
	HistoryMessages []model.Message
	Optional        bool
}

type requestAiParams struct {
	Optional    bool
	ChatHistory []*schema.Message
	Question    string
	Uuid        string
	UserId      int32
}

func NewBaseAiService(ctx context.Context) *BaseAiService {
	return &BaseAiService{ctx: ctx}
}

func (s *BaseAiService) Run(req *model.BaseRequest, handler model.AIConversationHandler) (resp *model.BaseResponse, err error) {
	if req.Question == "" {
		klog.Info("智能助手：用户输入为空：", req)
		return &model.BaseResponse{
			StatusCode: 2000,
			StatusMsg:  constant.GetMsg(2000),
		}, nil
	}

	userId := req.UserId
	uuid := req.Uuid

	exist, err := model.ConversionExist(mysql.DB, s.ctx, userId, uuid)
	if err != nil {
		klog.CtxErrorf(s.ctx, "智能助手：查询会话是否存在失败，会话id: %s, err: %v", uuid, err)
		return nil, errors.WithStack(err)
	}

	chatHistory, historyMessages, optional, err := s.getChatHistory(exist, uuid, userId)
	if err != nil {
		return nil, err
	}

	txParams := &transactionParams{
		UserId:          userId,
		Question:        req.Question,
		Uuid:            uuid,
		ChatHistory:     chatHistory,
		HistoryMessages: historyMessages,
		Optional:        optional,
	}

	aiResponseContent, err := s.processTransaction(txParams, handler, req.Scenario)
	if err != nil {
		return nil, err
	}

	data, err := handler.ProcessResponse(aiResponseContent)
	if err != nil {
		klog.Errorf("智能助手：处理AI响应失败，AI回复内容: %s, err: %v", aiResponseContent, err)
		return nil, errors.WithStack(err)
	}

	return &model.BaseResponse{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
		Data:       data,
	}, nil
}

func (s *BaseAiService) getChatHistory(exist bool, uuid string, userId int32) (chatHistory []*schema.Message, historyMessages []model.Message, optional bool, err error) {
	ctx := s.ctx
	optional = true
	if exist {
		optional = false
		cacheMessages, err := redis.RedisClient.LRange(ctx, redisUtils.GetChatHistoryKey(uuid), 0, 19).Result()
		if err != nil {
			klog.CtxErrorf(ctx, "智能助手：查询会话历史对话失败，会话id: %s, err: %v", uuid, err)
		}
		if len(cacheMessages) > 0 {
			for _, value := range cacheMessages {
				message := &schema.Message{}
				err = sonic.Unmarshal([]byte(value), &message)
				if err != nil {
					klog.CtxErrorf(ctx, "智能助手：json解析会话历史对话失败，value: %s, err: %v", value, err)
					err = errors.WithStack(err)
					return nil, nil, false, err
				}
				chatHistory = append(chatHistory, message)
			}
		} else {
			historyMessages, err = model.GetChatHistoryByUuid(mysql.DB, ctx, uuid, userId)
			if err != nil {
				klog.CtxErrorf(ctx, "智能助手：查询会话历史对话失败，会话id: %s, err: %v", uuid, err)
				return nil, nil, false, errors.WithStack(err)
			}
			chatHistory = make([]*schema.Message, len(historyMessages))
			for i, message := range historyMessages {
				chatHistory[i] = &schema.Message{
					Role:    schema.RoleType(message.Role),
					Content: message.Content,
				}
			}
		}
	}
	return
}

func (s *BaseAiService) processTransaction(params *transactionParams, handler model.AIConversationHandler, scenario model.Scenario) (aiResponseContent string, err error) {
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		userMessage := &model.Message{
			UserId:   params.UserId,
			Role:     model.RoleUser,
			Content:  params.Question,
			Uuid:     params.Uuid,
			Scenario: model.OrderInquiry,
		}
		if err = model.CreateMessage(tx, s.ctx, userMessage); err != nil {
			klog.CtxErrorf(s.ctx, "智能助手：用户消息写入数据库失败，message: %s, err: %v", userMessage, err)
			return errors.WithStack(err)
		}

		requestAiParams := &requestAiParams{
			Optional:    params.Optional,
			ChatHistory: params.ChatHistory,
			Question:    params.Question,
			Uuid:        params.Uuid,
			UserId:      params.UserId,
		}
		aiMessage, err := s.generateAiResponse(requestAiParams, handler, scenario)
		if err != nil {
			return err
		}
		aiResponseContent = aiMessage.Content
		err = model.CreateMessage(mysql.DB, s.ctx, aiMessage)
		if err != nil {
			klog.CtxErrorf(s.ctx, "智能助手：AI回复消息写入数据库失败，message: %s, err: %v", aiMessage, err)
			return errors.WithStack(err)
		}

		err = s.cacheChatMessages(userMessage, aiMessage, params.HistoryMessages, params.Uuid)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return "", errors.WithStack(err)
	}
	return aiResponseContent, nil
}

func (s *BaseAiService) generateAiResponse(params *requestAiParams, handler model.AIConversationHandler, scenario model.Scenario) (*model.Message, error) {
	template := prompt.FromMessages(schema.GoTemplate,
		schema.SystemMessage(handler.GetPrompt()),
		schema.MessagesPlaceholder("chat_history", params.Optional),
		schema.UserMessage("{{.query}}"),
	)
	messages, err := template.Format(s.ctx, map[string]any{
		"datetime":     utils.GetCurrentFormattedDateTime(),
		"query":        params.Question,
		"chat_history": params.ChatHistory,
	})
	if err != nil {
		klog.Errorf("智能助手：格式化对话信息失败，会话id: %s, err: %v", params.Uuid, err)
		return nil, errors.WithStack(err)
	}

	chatModel, err := ark.NewChatModel(s.ctx, &ark.ChatModelConfig{
		Model:  conf.GetConf().Ark.Model,
		APIKey: conf.GetConf().Ark.ApiKey,
	})
	maxRetries := 3
	for i := 0; i < maxRetries; i++ {
		result, err := chatModel.Generate(s.ctx, messages)
		if err != nil {
			klog.CtxErrorf(s.ctx, "智能助手：生成对话失败，尝试第%d次，err: %v", i+1, err)
			time.Sleep(time.Millisecond * 500)
			continue
		}
		return &model.Message{
			UserId:   params.UserId,
			Role:     model.RoleAssistant,
			Content:  result.Content,
			Uuid:     params.Uuid,
			Scenario: scenario,
		}, nil
	}
	return nil, err
}

func (s *BaseAiService) cacheChatMessages(userMessage *model.Message, aiMessage *model.Message, historyMessages []model.Message, uuid string) error {
	userMsgStr, err := sonic.Marshal(userMessage)
	if err != nil {
		klog.CtxErrorf(s.ctx, "智能助手：json序列化用户消息失败，message: %s, err: %v", userMessage, err)
		return errors.WithStack(err)
	}
	aiMsgStr, err := sonic.Marshal(aiMessage)
	if err != nil {
		klog.CtxErrorf(s.ctx, "智能助手：json序列化AI回复消息失败，message: %s, err: %v", aiMessage, err)
		return errors.WithStack(err)
	}

	preparedCacheMessages := make([]string, 0, len(historyMessages)+2)
	for _, historyMessage := range historyMessages {
		historyMsgStr, _ := sonic.Marshal(historyMessage)
		preparedCacheMessages = append(preparedCacheMessages, string(historyMsgStr))
	}
	preparedCacheMessages = append(preparedCacheMessages, string(userMsgStr), string(aiMsgStr))

	err = redis.RedisClient.RPush(s.ctx, redisUtils.GetChatHistoryKey(uuid), preparedCacheMessages).Err()
	if err != nil {
		klog.CtxErrorf(s.ctx, "智能助手：保存消息到缓存失败，message: %s, err: %v", userMessage, err)
		return errors.WithStack(err)
	}
	redis.RedisClient.Expire(s.ctx, redisUtils.GetChatHistoryKey(uuid), time.Minute*10)
	return nil
}

func PutAiResponse(ctx context.Context, userId int32, uuid string, content string, role string, scenario int8) error {
	aiMessage := &model.Message{
		UserId:   userId,
		Role:     model.Role(role),
		Content:  content,
		Uuid:     uuid,
		Scenario: model.Scenario(scenario),
	}
	err := model.CreateMessage(mysql.DB, ctx, aiMessage)
	if err != nil {
		klog.CtxErrorf(ctx, "智能助手：AI回复消息写入数据库失败，message: %s, err: %v", aiMessage, err)
		return errors.WithStack(err)
	}
	jsonStrBytes, _ := sonic.Marshal(aiMessage)
	redis.RedisClient.RPush(ctx, redisUtils.GetChatHistoryKey(uuid), string(jsonStrBytes))
	return nil
}
