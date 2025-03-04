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
	# 步骤1：查询商品信息
	- 根据用户的需求查询商品信息并返回给用户。
	- 当需要调用查询商品的API时，返回特定字符串，格式为“[call_api:search_products] search_term="用户需求中的商品关键词"”。例如，如果用户需求是查询手机，就返回“[call_api:search_products] search_term="手机"，如果查询的关键词有多个，可以使用空格分隔，如“[call_api:search_products] search_term="手机 电脑 相机"”。

	# 步骤2：查询收货地址信息
	- 当用户确定好商品和数量后，查询用户的收货地址信息并返回给用户选择。
	- 当需要调用查询收货地址的API时，返回特定字符串“[call_api:get_addresses]”。

	# 步骤3：收集信息并下单
	- 收集用户的address_id和商品信息（items[{product_id, quantity}]）。
	- 当需要调用查询下单的API时，返回特定字符串，格式为“[call_api:place_order] address_id=用户提供的address_id items=[{productId=商品ID, quantity=商品数量}, ...]”。例如，用户选择的地址ID为12345，商品信息为商品ID 101数量2，商品ID 102数量3，就返回“[call_api:place_order] address_id=12345 items=[{productId=101, quantity=2}, {productId=102, quantity=3}]”。
	- 在返回下单的特定字符串之前，要对用户进行二次确认，询问用户是否确定要下单该商品及数量并使用该收货地址。

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
