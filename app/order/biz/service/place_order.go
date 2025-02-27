package service

import (
	"context"
	commonconstant "douyin_mall/common/constant"
	"douyin_mall/order/biz/dal/mysql"
	"douyin_mall/order/biz/model"
	"douyin_mall/order/infra/kafka/constant"
	"douyin_mall/order/infra/kafka/producer"
	order "douyin_mall/order/kitex_gen/order"
	"douyin_mall/order/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	ctx := s.ctx
	orderId := utils.GetSnowflakeID()
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		newOrder := &model.Order{
			OrderID:       orderId,
			UserID:        req.UserId,
			TotalCost:     req.TotalCost,
			Name:          req.Address.Name,
			PhoneNumber:   req.Address.PhoneNumber,
			Province:      req.Address.Province,
			City:          req.Address.City,
			Region:        req.Address.Region,
			DetailAddress: req.Address.DetailAddress,
		}
		err = model.CreateOrder(ctx, tx, newOrder)
		if err != nil {
			return errors.WithStack(err)
		}
		orderItemList := make([]*model.OrderItem, len(req.OrderItems))
		for i, item := range req.OrderItems {
			orderItemList[i] = &model.OrderItem{
				OrderID:   orderId,
				Cost:      item.Cost,
				ProductID: item.Item.ProductId,
				Quantity:  item.Item.Quantity,
			}
		}
		err = model.CreateOrderItems(ctx, tx, orderItemList)
		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if err != nil {
		klog.CtxErrorf(ctx, "创建订单失败：req: %v, err: %v", req, err)
		return nil, errors.WithStack(err)
	}
	// todo 批量锁定库存

	// 延时取消订单
	producer.SendDelayOrder(ctx, orderId, constant.DelayTopic1mLevel)
	return &order.PlaceOrderResp{
		Order: &order.OrderResult{
			OrderId: orderId,
		},
		StatusCode: 0,
		StatusMsg:  commonconstant.GetMsg(0),
	}, nil
}
