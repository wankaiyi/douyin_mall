package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/doubao_ai/biz/model"
	"douyin_mall/doubao_ai/kitex_gen/doubao_ai"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
)

type SmartSearchOrderHandler struct{}

func (h *SmartSearchOrderHandler) GetPrompt() string {
	return `现在的时间是{{.datetime}}，我想让你作为一个抖音商城智能购物助手的信息提取专家，专注于从用户的查询语句中精准提取与商品有关的信息，并以JSON格式呈现。JSON的格式需为 {"start_time": "yyyy - MM - dd HH:mm:ss", "end_time": "yyyy - MM - dd HH:mm:ss", "search_terms": ["item1", "item2"]}，属性的值可以为空字符串或空数组。你的任务是：

	1. 仔细阅读并理解用户的查询和对话的上下文，忽略与商品名称、描述和时间无关的修饰信息。
	2. 提取出与商品相关的关键词或类型描述，以及提到的下单时间范围（如果有）。如果在之前的对话中有明确提及商品名称且未被新的对话内容否定，应将其包含在 search_terms 中。
	3. 对于时间范围，只有当用户明确提及了具体的下单时间范围时，才将对应的时间填入 start_time 和 end_time 字段。若用户未提及时间范围，start_time 和 end_time 都应设置为空字符串。
	4. 将提取的信息组织成简洁、准确的JSON格式，确保没有多余内容或信息丢失。

	请确保输出的JSON准确反映用户的查询和对话上下文需求，并在解析时考虑不同表述的变体和模糊性。`
}

func (h *SmartSearchOrderHandler) ProcessResponse(content string) (interface{}, error) {
	var result *doubao_ai.SearchOrderQuestionResp
	err := sonic.Unmarshal([]byte(content), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type AnalyzeSearchOrderQuestionService struct {
	ctx context.Context
} // NewAnalyzeSearchOrderQuestionService new AnalyzeSearchOrderQuestionService
func NewAnalyzeSearchOrderQuestionService(ctx context.Context) *AnalyzeSearchOrderQuestionService {
	return &AnalyzeSearchOrderQuestionService{ctx: ctx}
}

// Run create note info
func (s *AnalyzeSearchOrderQuestionService) Run(req *doubao_ai.SearchOrderQuestionReq) (resp *doubao_ai.SearchOrderQuestionResp, err error) {
	ctx := s.ctx
	baseReq := &model.BaseRequest{
		UserId:   req.UserId,
		Question: req.Question,
		Uuid:     req.Uuid,
		Scenario: model.OrderInquiry,
	}
	handler := &SmartSearchOrderHandler{}
	baseResp, err := NewBaseAiService(s.ctx).Run(baseReq, handler)
	if err != nil {
		return nil, err
	}
	resp, ok := baseResp.Data.(*doubao_ai.SearchOrderQuestionResp)
	if !ok {
		klog.CtxErrorf(ctx, "类型转换错误，data：%v", baseResp.Data)
		return nil, errors.New("智能购物助手返回结果异常")
	}
	resp.StatusCode = 0
	resp.StatusMsg = constant.GetMsg(0)
	return resp, nil
}
