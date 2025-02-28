package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/common/utils"
	"douyin_mall/order/biz/dal/mysql"
	"douyin_mall/order/biz/model"
	order "douyin_mall/order/kitex_gen/order"
	"github.com/pkg/errors"
)

type GetOrderService struct {
	ctx context.Context
} // NewGetOrderService new GetOrderService
func NewGetOrderService(ctx context.Context) *GetOrderService {
	return &GetOrderService{ctx: ctx}
}

// Run create note info
func (s *GetOrderService) Run(req *order.GetOrderReq) (resp *order.GetOrderResp, err error) {
	ctx := s.ctx
	o, err := model.GetOrder(ctx, mysql.DB, req.UserId, req.OrderId)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var products []*order.Product
	for _, item := range o.OrderItems {
		products = append(products, &order.Product{
			Id:          item.ProductID,
			Name:        item.ProductName,
			Price:       item.ProductPrice,
			Quantity:    item.Quantity,
			Picture:     item.ProductPicture,
			Description: item.ProductDescription,
		})
	}
	orderInfo := &order.Order{
		OrderId: o.OrderID,
		Address: &order.Address{
			Name:          o.Name,
			PhoneNumber:   o.PhoneNumber,
			Province:      o.Province,
			City:          o.City,
			Region:        o.Region,
			DetailAddress: o.DetailAddress,
		},
		Products:  products,
		Cost:      o.TotalCost,
		CreatedAt: utils.GetFormattedDateTime(o.CreatedAt),
	}

	return &order.GetOrderResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
		Order:      orderInfo,
	}, nil
}
