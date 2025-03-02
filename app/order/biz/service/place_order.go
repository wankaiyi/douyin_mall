package service

import (
	"context"
	commonconstant "douyin_mall/common/constant"
	"douyin_mall/order/biz/dal/mysql"
	"douyin_mall/order/biz/model"
	"douyin_mall/order/infra/kafka/constant"
	"douyin_mall/order/infra/kafka/producer"
	"douyin_mall/order/infra/rpc"
	order "douyin_mall/order/kitex_gen/order"
	"douyin_mall/order/utils"
	"douyin_mall/rpc/kitex_gen/product"
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
	var productIds []int64
	for _, item := range req.OrderItems {
		productIds = append(productIds, int64(item.Item.ProductId))
	}
	productListReq := &product.SelectProductListReq{
		Ids: productIds,
	}
	productListResp, err := rpc.ProductClient.SelectProductList(ctx, productListReq)
	if err != nil {
		klog.CtxErrorf(ctx, "rpc获取商品信息失败：req: %v, err: %v", productListReq, err)
		return nil, errors.WithStack(err)
	}
	productIdToObj := make(map[int64]*product.Product)
	for _, productInfo := range productListResp.Products {
		productIdToObj[productInfo.Id] = productInfo
	}
	orderItemList := make([]*model.OrderItem, len(req.OrderItems))
	for i, item := range req.OrderItems {
		orderItemList[i] = &model.OrderItem{
			OrderID: orderId,
			Cost:    item.Cost,
			Product: model.Product{
				ProductID:          item.Item.ProductId,
				ProductName:        productIdToObj[int64(item.Item.ProductId)].Name,
				ProductPrice:       float64(productIdToObj[int64(item.Item.ProductId)].Price),
				ProductDescription: productIdToObj[int64(item.Item.ProductId)].Description,
				ProductPicture:     productIdToObj[int64(item.Item.ProductId)].Picture,
			},
			Quantity: item.Item.Quantity,
		}
	}

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
