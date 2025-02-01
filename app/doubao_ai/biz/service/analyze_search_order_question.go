package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/common/utils"
	"douyin_mall/doubao_ai/biz/dal/mysql"
	"douyin_mall/doubao_ai/biz/dal/redis"
	"douyin_mall/doubao_ai/biz/model"
	"douyin_mall/doubao_ai/conf"
	"douyin_mall/doubao_ai/kitex_gen/doubao_ai"
	redisUtils "douyin_mall/doubao_ai/utils/redis"
	"encoding/json"
	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type AnalyzeSearchOrderQuestionService struct {
	ctx context.Context
} // NewAnalyzeSearchOrderQuestionService new AnalyzeSearchOrderQuestionService
func NewAnalyzeSearchOrderQuestionService(ctx context.Context) *AnalyzeSearchOrderQuestionService {
	return &AnalyzeSearchOrderQuestionService{ctx: ctx}
}

// Run create note info
func (s *AnalyzeSearchOrderQuestionService) Run(req *doubao_ai.SearchOrderQuestionReq) (resp *doubao_ai.SearchOrderQuestionResp, err error) {
	ctx := s.ctx
	question := req.Question

	if question == "" {
		klog.Info("智能购物助手：用户查询订单输入为空：", req)
		return &doubao_ai.SearchOrderQuestionResp{
			StatusCode: 2000,
			StatusMsg:  constant.GetMsg(2000),
		}, nil
	}

	userId := req.UserId
	uuid := req.Uuid

	// 获取这个会话中用户的历史对话
	exist, err := model.ConversionExist(mysql.DB, ctx, uuid)
	if err != nil {
		klog.Errorf("智能购物助手：查询会话是否存在失败，会话id: %s, err: %v", uuid, err)
		return nil, errors.WithStack(err)
	}
	optional := true
	var chatHistory []*schema.Message
	var historyMessages []model.Message
	if exist {
		// 历史对话存在
		optional = false
		cacheMessages, err := redis.RedisClient.LRange(ctx, redisUtils.GetChatHistoryKey(uuid), 0, 9).Result()
		if err != nil {
			klog.Errorf("智能购物助手：查询会话历史对话失败，会话id: %s, err: %v", uuid, err)
		}
		if len(cacheMessages) > 0 {
			for _, value := range cacheMessages {
				message := &schema.Message{}
				err = json.Unmarshal([]byte(value), &message)
				if err != nil {
					klog.Errorf("智能购物助手：json解析会话历史对话失败，value: %s, err: %v", value, err)
					return nil, errors.WithStack(err)
				}
				chatHistory = append(chatHistory, message)
			}
		} else {
			// 从数据库中获取历史对话
			historyMessages, err = model.GetChatHistoryByUuid(mysql.DB, ctx, uuid)
			for _, message := range historyMessages {
				chatHistory = append(chatHistory, &schema.Message{
					Role:    schema.RoleType(message.Role),
					Content: message.Content,
				})
			}
			if err != nil {
				klog.Errorf("智能购物助手：查询会话历史对话失败，会话id: %s, err: %v", uuid, err)
				return nil, errors.WithStack(err)
			}
		}
	}

	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		userMessage := &model.Message{
			UserId:   userId,
			Role:     model.RoleUser,
			Content:  question,
			Uuid:     uuid,
			Scenario: model.OrderInquiry,
		}

		// 先插数据库再存缓存，防止事务回滚导致redis缓存不一致
		err = model.CreateMessage(tx, ctx, userMessage)
		if err != nil {
			klog.Errorf("智能购物助手：用户消息写入数据库失败，message: %s, err: %v", userMessage, err)
			return errors.WithStack(err)
		}

		searchOrderPrompt := "现在的时间是{{.datetime}}，我想让你作为一个智能购物助手的信息提取专家，专注于从用户的查询语句中精准提取与商品有关的信息，并以JSON格式呈现。JSON的格式需为 {start_time: \"yyyy - MM - dd HH:mm:ss\", end_time: \"yyyy - MM - dd HH:mm:ss\", search_terms: [\"item1\", \"item2\"]}，属性的值可以为空字符串或空数组。你的任务是：\n\n1. 仔细阅读并理解用户的查询和对话的上下文，忽略与商品名称和时间无关的修饰信息。\n2. 提取出与商品相关的关键词或类型描述，以及提到的下单时间范围（如果有）。如果在之前的对话中有明确提及商品名称且未被新的对话内容否定，应将其包含在search_terms中。\n3. 将提取的信息组织成简洁、准确的JSON格式，确保没有多余内容或信息丢失。\n\n请确保输出的JSON准确反映用户的查询和对话上下文需求，并在解析时考虑不同表述的变体和模糊性。"
		template := prompt.FromMessages(schema.GoTemplate,
			schema.SystemMessage(searchOrderPrompt),
			// optional=false 表示必需的消息列表，在模版输入中找不到对应变量会报错
			schema.MessagesPlaceholder("chat_history", optional),
			schema.UserMessage("{{.query}}"),
		)
		messages, err := template.Format(ctx, map[string]any{
			"datetime":     utils.GetCurrentFormattedTime(),
			"query":        question,
			"chat_history": chatHistory,
		})
		if err != nil {
			klog.Errorf("智能购物助手：查询订单格式化对话信息失败，会话id: %s, err: %v", uuid, err)
			return errors.WithStack(err)
		}

		chatModel, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
			Model:  conf.GetConf().Ark.Model,
			APIKey: conf.GetConf().Ark.ApiKey,
		})
		result, err := chatModel.Generate(ctx, messages)
		if err != nil {
			klog.Errorf("智能购物助手：查询订单生成对话失败：result: %s, err: %v", result, err)
			return errors.WithStack(err)
		}

		err = json.Unmarshal([]byte(result.Content), &resp)
		if err != nil {
			klog.Errorf("智能购物助手：json解析查询订单对话结果失败，result: %s, err: %v", result, err)
			return errors.WithStack(err)
		}

		aiMessage := &model.Message{
			UserId:   userId,
			Role:     model.RoleAssistant,
			Content:  result.Content,
			Uuid:     uuid,
			Scenario: model.OrderInquiry,
		}
		err = model.CreateMessage(mysql.DB, ctx, aiMessage)
		if err != nil {
			klog.Errorf("智能购物助手：AI回复消息写入数据库失败，message: %s, err: %v", aiMessage, err)
			return errors.WithStack(err)
		}

		// 将用户消息和AI回复消息存入缓存
		userMsgStr, err := json.Marshal(userMessage)
		if err != nil {
			klog.Errorf("智能购物助手：json序列化用户消息失败，message: %s, err: %v", userMessage, err)
			return errors.WithStack(err)
		}
		aiMsgStr, err := json.Marshal(aiMessage)
		if err != nil {
			klog.Errorf("智能购物助手：json序列化AI回复消息失败，message: %s, err: %v", aiMessage, err)
			return errors.WithStack(err)
		}

		var preparedCacheMessages []string
		for _, historyMessage := range historyMessages {
			historyMsgStr, _ := json.Marshal(historyMessage)
			preparedCacheMessages = append(preparedCacheMessages, string(historyMsgStr))
		}
		preparedCacheMessages = append(preparedCacheMessages, string(userMsgStr))
		preparedCacheMessages = append(preparedCacheMessages, string(aiMsgStr))

		err = redis.RedisClient.LPush(ctx, redisUtils.GetChatHistoryKey(uuid), preparedCacheMessages).Err()
		if err != nil {
			klog.Errorf("智能购物助手：保存消息到缓存失败，message: %s, err: %v", userMessage, err)
			return errors.WithStack(err)
		}
		redis.RedisClient.Expire(ctx, redisUtils.GetChatHistoryKey(uuid), time.Minute*10)

		return nil
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if err != nil {
		return nil, errors.WithStack(err)
	}
	resp.StatusCode = 0
	resp.StatusMsg = constant.GetMsg(0)
	return resp, nil
}
