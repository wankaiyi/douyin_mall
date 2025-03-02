package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/order/infra/rpc"
	"douyin_mall/order/kitex_gen/cart"
	order "douyin_mall/order/kitex_gen/order"
	"douyin_mall/rpc/kitex_gen/doubao_ai"
	"douyin_mall/rpc/kitex_gen/product"
	"douyin_mall/rpc/kitex_gen/user"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
	"regexp"
	"strconv"
	"strings"
)

const (
	specificPlaceOrderPrefix   = "[call_api:place_order]"
	specificSearchOrderPrefix  = "[call_api:search_products]"
	specificGetAddressesPrefix = "[call_api:get_addresses]"
	addressIdPattern           = `address_id=(\d+)`
	itemsPattern               = `items=\[({.*?})\]`
	itemPattern                = `productId=(\d+), quantity=(\d+)`
	searchTermsPattern         = `search_term="([^"]+)"`
)

var (
	addressIdRegex   = regexp.MustCompile(addressIdPattern) // 示例：[call_api:place_order] address_id=12345 items=[{productId=101, quantity=2}, {productId=102, quantity=3}]
	itemsRegex       = regexp.MustCompile(itemsPattern)
	itemRegex        = regexp.MustCompile(itemPattern)
	searchTermsRegex = regexp.MustCompile(searchTermsPattern)
)

type SmartPlaceOrderService struct {
	ctx context.Context
} // NewSmartPlaceOrderService new SmartPlaceOrderService
func NewSmartPlaceOrderService(ctx context.Context) *SmartPlaceOrderService {
	return &SmartPlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *SmartPlaceOrderService) Run(req *order.SmartPlaceOrderReq) (resp *order.SmartPlaceOrderResp, err error) {
	ctx := s.ctx
	// AI对话
	placeOrderQuestionReq := &doubao_ai.PlaceOrderQuestionReq{
		Question: req.Question,
		Uuid:     req.Uuid,
		UserId:   req.UserId,
	}
	placeOrderQuestionResp, err := rpc.DoubaoClient.AnalyzePlaceOrderQuestion(ctx, placeOrderQuestionReq)
	if err != nil {
		klog.CtxErrorf(ctx, "rpc调用AI分析用户下单问题失败，req：%v, err：%v", placeOrderQuestionReq, err)
		return nil, errors.WithStack(err)
	}
	klog.CtxInfof(ctx, "AI分析用户下单问题成功，req：%v, resp：%v", placeOrderQuestionReq, placeOrderQuestionResp)

	aiResponse := placeOrderQuestionResp.Response
	if placeOrderQuestionResp.StatusCode != 0 {
		resp = &order.SmartPlaceOrderResp{
			StatusCode: placeOrderQuestionResp.StatusCode,
			StatusMsg:  placeOrderQuestionResp.StatusMsg,
			Response:   aiResponse,
		}
		return resp, nil
	}

	// 如果返回特定的结果（正则匹配），则调用对应的接口
	var customerChatMessage string
	if strings.HasPrefix(aiResponse, specificPlaceOrderPrefix) {
		addressId, cartItems, err := parsePlaceOrderResponse(aiResponse)
		if err != nil {
			klog.CtxErrorf(ctx, "匹配到特定结果，但解析AI返回结果失败，ai结果：%v, err：%v", aiResponse, err)
			return nil, err
		}
		// 请求结算接口进行下单
		fmt.Println("addressId:", addressId)
		fmt.Println("cartItems:", cartItems)
		// todo 调用结算接口
		customerChatMessage = "下单成功，请点击链接进行支付：" + "https://www.baidu.com"
	} else if strings.HasPrefix(aiResponse, specificSearchOrderPrefix) {
		// 搜索商品
		searchTerm, err := parseSearchTermResponse(aiResponse)
		if err != nil {
			klog.CtxErrorf(ctx, "匹配到特定结果，但解析AI返回结果失败，ai结果：%v, err：%v", aiResponse, err)
			return nil, err
		}
		searchProductsReq := &product.SearchProductsReq{
			Query: searchTerm,
		}
		searchProductsResp, err := rpc.ProductClient.SearchProducts(ctx, searchProductsReq)
		if err != nil {
			klog.CtxErrorf(ctx, "rpc调用搜索商品失败，req：%v, err：%v", searchProductsReq, err)
			return nil, err
		}
		klog.CtxInfof(ctx, "搜索商品成功，req：%v, resp：%v", searchProductsReq, searchProductsResp)
		customerChatMessage = fmt.Sprintf("为您找到以下商品：%v", searchProductsResp.Results)
	} else if strings.HasPrefix(aiResponse, specificGetAddressesPrefix) {
		getReceiveAddressReq := &user.GetReceiveAddressReq{
			UserId: req.UserId,
		}
		getReceiveAddressResp, err := rpc.UserClient.GetReceiveAddress(ctx, getReceiveAddressReq)
		if err != nil {
			klog.CtxErrorf(ctx, "rpc调用获取用户收货地址失败，req：%v, err：%v", getReceiveAddressReq, err)
			return nil, err
		}
		klog.CtxInfof(ctx, "获取用户收货地址成功，req：%v, resp：%v", getReceiveAddressReq, getReceiveAddressResp)
		customerChatMessage = fmt.Sprintf("请选择收货地址：%v", getReceiveAddressResp.ReceiveAddress)
	} else {
		return &order.SmartPlaceOrderResp{
			StatusCode: 0,
			StatusMsg:  constant.GetMsg(0),
			Response:   aiResponse,
		}, nil
	}

	_, err = rpc.DoubaoClient.AddChatMessage(ctx, &doubao_ai.AddChatMessageReq{
		Content:  customerChatMessage,
		Uuid:     req.Uuid,
		UserId:   req.UserId,
		Role:     "assistant",
		Scenario: 2,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &order.SmartPlaceOrderResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
		Response:   customerChatMessage,
	}, nil
}

func parseSearchTermResponse(aiResponse string) (string, error) {
	if !regexp.MustCompile(regexp.QuoteMeta(specificSearchOrderPrefix)).MatchString(aiResponse) {
		return "", errors.New("未检测到特定字符串: " + specificSearchOrderPrefix)
	}
	searchTermMatch := searchTermsRegex.FindString(aiResponse)
	if len(searchTermMatch) < 2 {
		return "", errors.New("未找到 search_term")
	}
	return searchTermMatch, nil
}

func parsePlaceOrderResponse(aiResponse string) (int, []cart.CartItem, error) {
	// 检查是否包含特定字符串
	if !regexp.MustCompile(regexp.QuoteMeta(specificPlaceOrderPrefix)).MatchString(aiResponse) {
		return 0, nil, errors.Errorf("未检测到特定字符串: %s", specificPlaceOrderPrefix)
	}

	addressIDMatch := addressIdRegex.FindStringSubmatch(aiResponse)
	if len(addressIDMatch) < 2 {
		return 0, nil, errors.New("未找到 address_id")
	}
	addressID, err := strconv.Atoi(addressIDMatch[1])
	if err != nil {
		return 0, nil, errors.Errorf("address_id 格式错误: %v", err)
	}

	itemsMatch := itemsRegex.FindStringSubmatch(aiResponse)
	if len(itemsMatch) < 2 {
		return 0, nil, errors.New("未找到 items 列表")
	}
	itemsStr := itemsMatch[1]

	itemMatches := itemRegex.FindAllStringSubmatch(itemsStr, -1)
	if len(itemMatches) == 0 {
		return 0, nil, errors.New("未找到有效的 items")
	}

	var items []cart.CartItem
	for _, match := range itemMatches {
		productId, _ := strconv.ParseInt(match[1], 10, 32)
		quantity, _ := strconv.ParseInt(match[2], 10, 32)
		items = append(items, cart.CartItem{ProductId: (int32)(productId), Quantity: (int32)(quantity)})
	}

	return addressID, items, nil
}
