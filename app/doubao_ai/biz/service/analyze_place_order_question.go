package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/doubao_ai/biz/model"
	doubao_ai "douyin_mall/doubao_ai/kitex_gen/doubao_ai"
)

type SmartPlaceOrderHandler struct{}

func (h *SmartPlaceOrderHandler) GetPrompt() string {
	return `
	我想让你作为一个温柔且乐于助人的抖音商城的智能购物助手，根据用户的需求帮助用户对商品下单。首先，请仔细阅读用户的需求，包括对话的上下文，并从上下文中分析出需要的数据。
	接下来，请按照以下步骤完成任务：

	步骤1：查询商品信息
	- 任务：根据用户的需求查询商品信息。
	- 操作：
	- - 当用户提出商品需求时，直接返回特定字符串，格式为：
	[call_api:search_products] search_term="用户需求中的商品关键词"。
	- - 示例：
	- - - 用户输入：“我想买手机。”
	输出：[call_api:search_products] search_term="手机"。
	- - - 用户输入：“帮我找一下手机和电脑。”
	输出：[call_api:search_products] search_term="手机 电脑"。

	步骤2：查询收货地址信息
	- 任务：当用户确定商品和数量后，查询用户的收货地址信息。
	- 操作：
	- - 直接返回特定字符串：[call_api:get_addresses]。
	- - 示例：
	- - - 用户输入：“我确定要买这个商品。”
	输出：[call_api:get_addresses]。

	步骤3：收集信息并下单
	- 任务：收集用户的地址信息和商品信息（items[{product_id, quantity}]），并进行二次确认。
	- 操作：
	1. 1. 当用户提供地址信息和商品信息后，直接返回特定字符串，格式为：
	[call_api:place_order] address_id=用户提供的address_id items=[{productId=商品ID, quantity=商品数量}, ...]。
	1. 2. 示例：
	- - - 用户输入：“使用猎德村这个地址”
	输出：[call_api:place_order] address_id=12345 items=[{productId=101, quantity=2}, {productId=102, quantity=3}]。

	额外规则
	1. 禁止返回多余内容：
	- - 仅输出特定字符串或必要信息，不添加任何解释、确认或其他多余内容。
	2. 严格遵循格式：
	- - 确保输出的字符串格式完全符合要求，包括大小写、标点符号和空格。

	请按照上述步骤和要求进行操作，准确输出相应的特定字符串和进行二次确认。`
}

func (h *SmartPlaceOrderHandler) ProcessResponse(content string) (interface{}, error) {
	return content, nil
}

type AnalyzePlaceOrderQuestionService struct {
	ctx context.Context
} // NewAnalyzePlaceOrderQuestionService new AnalyzePlaceOrderQuestionService
func NewAnalyzePlaceOrderQuestionService(ctx context.Context) *AnalyzePlaceOrderQuestionService {
	return &AnalyzePlaceOrderQuestionService{ctx: ctx}
}

// Run create note info
func (s *AnalyzePlaceOrderQuestionService) Run(req *doubao_ai.PlaceOrderQuestionReq) (resp *doubao_ai.PlaceOrderQuestionResp, err error) {
	//ctx := s.ctx
	baseReq := &model.BaseRequest{
		UserId:   req.UserId,
		Question: req.Question,
		Uuid:     req.Uuid,
		Scenario: model.MockPlaceOrder,
	}
	handler := &SmartPlaceOrderHandler{}
	baseResp, err := NewBaseAiService(s.ctx).Run(baseReq, handler)
	if err != nil {
		return nil, err
	}
	return &doubao_ai.PlaceOrderQuestionResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
		Response:   baseResp.Data.(string),
	}, nil
}
