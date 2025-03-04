package service

import (
	"context"
	"douyin_mall/checkout/infra/rpc"
	checkout "douyin_mall/checkout/kitex_gen/checkout"
	"douyin_mall/checkout/kitex_gen/user"
	"douyin_mall/rpc/kitex_gen/cart"
	"douyin_mall/rpc/kitex_gen/order"
	"douyin_mall/rpc/kitex_gen/payment"
	"douyin_mall/rpc/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
)

type CheckoutProductItemsService struct {
	ctx context.Context
} // NewCheckoutProductItemsService new CheckoutProductItemsService
func NewCheckoutProductItemsService(ctx context.Context) *CheckoutProductItemsService {
	return &CheckoutProductItemsService{ctx: ctx}
}

// Run create note info
func (s *CheckoutProductItemsService) Run(req *checkout.CheckoutProductItemsReq) (resp *checkout.CheckoutProductItemsResp, err error) {
	//得到用户地址
	getReceiveAddressResp, err := rpc.UserClient.GetReceiveAddress(s.ctx, &user.GetReceiveAddressReq{
		UserId: req.UserId,
	})
	if err != nil {
		klog.CtxErrorf(s.ctx, "获取用户地址失败rpc接口调用失败, error: %v", err)
		return nil, errors.WithStack(err)
	}
	addressList := getReceiveAddressResp.ReceiveAddress
	var targetAddress *user.ReceiveAddress
	for _, address := range addressList {
		if address.Id == req.AddressId {
			targetAddress = address
			break
		}
	}
	if targetAddress == nil {
		klog.CtxErrorf(s.ctx, "用户地址不存在, address_id: %d", req.AddressId)
		return nil, errors.New("用户地址不存在")
	}

	var totalCost float32
	//计算总价
	for _, productItem := range req.Items {
		getProductResp, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{
			Id: uint32(productItem.GetProductId()),
		})
		if err != nil {
			klog.CtxErrorf(s.ctx, "获取商品信息失败rpc接口调用失败, error: %v, product_id: %d", err, productItem.GetProductId())
		}
		totalCost += getProductResp.Product.Price

	}
	//创建订单
	placeOrderResp, err := rpc.OrderClient.PlaceOrder(s.ctx, &order.PlaceOrderReq{
		UserId: req.UserId,
		Address: &order.Address{
			Name:          targetAddress.Name,
			PhoneNumber:   targetAddress.PhoneNumber,
			Province:      targetAddress.Province,
			City:          targetAddress.City,
			Region:        targetAddress.Region,
			DetailAddress: targetAddress.DetailAddress,
		},
		OrderItems: convertProductItem2OrderItem(s.ctx, req.Items),
		TotalCost:  float64(totalCost),
	})
	if err != nil {
		klog.CtxErrorf(s.ctx, "创建订单失败rpc接口调用失败, error: %v", err)
		return nil, errors.WithStack(err)
	}
	orderId := placeOrderResp.Order.OrderId

	//得到订单金额
	orderResp, err := rpc.OrderClient.GetOrder(s.ctx, &order.GetOrderReq{
		OrderId: orderId,
	})
	if err != nil {
		klog.CtxErrorf(s.ctx, "获取订单失败rpc接口调用失败, error: %v", err)
		return nil, errors.WithStack(err)
	}
	cost := orderResp.Order.Cost
	chargeResp, err := rpc.PaymentClient.Charge(s.ctx, &payment.ChargeReq{
		OrderId: orderId,
		Amount:  float32(cost),
	})
	if err != nil {
		klog.CtxErrorf(s.ctx, "支付失败rpc接口调用失败, error: %v", err)
		return nil, errors.WithStack(err)
	}

	resp = &checkout.CheckoutProductItemsResp{
		StatusCode: chargeResp.StatusCode,
		StatusMsg:  chargeResp.StatusMsg,
		PaymentUrl: chargeResp.PaymentUrl,
	}
	return
}

func convertProductItem2OrderItem(ctx context.Context, productItems []*checkout.ProductItem) []*order.OrderItem {
	var orderItems []*order.OrderItem
	for _, productItem := range productItems {
		productResp, err := rpc.ProductClient.GetProduct(ctx, &product.GetProductReq{
			Id: uint32(productItem.ProductId),
		})
		if err != nil {
			klog.CtxErrorf(ctx, "获取商品信息失败rpc接口调用失败, error: %v", err)

		}
		if productResp.Product == nil {
			klog.CtxErrorf(ctx, "商品不存在, product_id: %d", productItem.ProductId)
		}
		//计算商品的总价
		cost := float32(productItem.Quantity) * productResp.Product.Price

		item := &order.OrderItem{
			Item: &cart.CartItem{
				ProductId: productItem.ProductId,
				Quantity:  productItem.Quantity,
			},
			Cost: float64(cost),
		}
		orderItems = append(orderItems, item)

	}
	return orderItems
}
