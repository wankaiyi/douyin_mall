package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/common/utils"
	"douyin_mall/order/biz/dal/mysql"
	"douyin_mall/order/biz/model"
	"douyin_mall/order/infra/rpc"
	order "douyin_mall/order/kitex_gen/order"
	"douyin_mall/rpc/kitex_gen/doubao_ai"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
)

type SmartSearchOrderService struct {
	ctx context.Context
} // NewSmartSearchOrderService new SmartSearchOrderService
func NewSmartSearchOrderService(ctx context.Context) *SmartSearchOrderService {
	return &SmartSearchOrderService{ctx: ctx}
}

// Run create note info
func (s *SmartSearchOrderService) Run(req *order.SmartSearchOrderReq) (resp *order.SmartSearchOrderResp, err error) {
	ctx := s.ctx
	searchOrderQuestionReq := &doubao_ai.SearchOrderQuestionReq{
		Uuid:     req.Uuid,
		UserId:   req.UserId,
		Question: req.Question,
	}
	analyzeResp, err := rpc.DoubaoClient.AnalyzeSearchOrderQuestion(ctx, searchOrderQuestionReq)
	if err != nil {
		klog.CtxErrorf(ctx, "rpc调用AI分析用户问题失败，req：%v, err：%v", searchOrderQuestionReq, err)
		return nil, errors.WithStack(err)
	}

	if analyzeResp.StatusCode != 0 {
		return &order.SmartSearchOrderResp{
			StatusCode: analyzeResp.StatusCode,
			StatusMsg:  analyzeResp.StatusMsg,
		}, nil
	}

	orders, err := model.SmartSearchOrder(ctx, mysql.DB, req.UserId, analyzeResp.SearchTerms, analyzeResp.StartTime, analyzeResp.EndTime)
	if err != nil {
		klog.CtxErrorf(ctx, "数据库查询订单失败，err：%v", err)
		return nil, errors.WithStack(err)
	}

	var orderList []*order.Order
	for _, o := range orders {
		var productList []*order.Product
		for _, p := range o.OrderItems {
			productList = append(productList, &order.Product{
				Id:          p.ProductID,
				Name:        p.ProductName,
				Description: p.ProductDescription,
				Picture:     p.ProductPicture,
				Price:       p.ProductPrice,
				Quantity:    p.Quantity,
			})
		}
		orderList = append(orderList, &order.Order{
			OrderId: o.OrderID,
			Address: &order.Address{
				Name:          o.Name,
				PhoneNumber:   o.PhoneNumber,
				Province:      o.Province,
				City:          o.City,
				Region:        o.Region,
				DetailAddress: o.DetailAddress,
			},
			Products:  productList,
			Cost:      o.TotalCost,
			CreatedAt: utils.GetFormattedDateTime(o.CreatedAt),
			Status:    o.Status,
		})
	}

	resp = &order.SmartSearchOrderResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
		Orders:     orderList,
	}

	return resp, nil
}
