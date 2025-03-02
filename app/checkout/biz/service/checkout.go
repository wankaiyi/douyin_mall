package service

import (
	"context"
	redsync "douyin_mall/checkout/biz/dal/red_sync"
	"douyin_mall/checkout/infra/rpc"
	checkout "douyin_mall/checkout/kitex_gen/checkout"
	"douyin_mall/common/constant"
	"douyin_mall/rpc/kitex_gen/cart"
	"douyin_mall/rpc/kitex_gen/order"
	"douyin_mall/rpc/kitex_gen/payment"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
	"strconv"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	//获得userId
	userId := req.UserId
	//得到互斥锁
	rsync := redsync.GetRedsync()
	mutexName := "checkout_mutex_" + strconv.FormatInt(int64(userId), 10)
	mutex := rsync.NewMutex(mutexName)
	//加锁
	if err := mutex.Lock(); err != nil {
		klog.CtxErrorf(s.ctx, "获取互斥锁失败，lock failed, :%d,err:%s", userId, err.Error())
		return nil, err
	}
	defer mutex.Unlock()

	//获得购物车信息
	cartResp, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: int32(userId)})
	if err != nil {
		klog.CtxErrorf(s.ctx, "获取购物车信息失败:%v,参数:req:%+v", err, req)
		return nil, errors.WithStack(err)
	}
	productItems := cartResp.Products

	//得到用户地址
	address := &order.Address{
		Name:          req.Address.ReceiveAddress.GetName(),
		PhoneNumber:   req.Address.ReceiveAddress.GetPhoneNumber(),
		DetailAddress: req.Address.ReceiveAddress.GetDetailAddress(),
		City:          req.Address.ReceiveAddress.GetCity(),
		Province:      req.Address.ReceiveAddress.GetProvince(),
		Region:        req.Address.ReceiveAddress.GetRegion(),
	}

	//创建订单信息
	placeOrderResp, err := rpc.OrderClient.PlaceOrder(s.ctx, &order.PlaceOrderReq{UserId: int32(userId), Address: address, OrderItems: convertCartProductItems2OrderItems(productItems)})
	if err != nil {
		klog.CtxErrorf(s.ctx, "创建订单失败: %v ,参数:req:%+v", err, req)
		return nil, errors.WithStack(err)
	}

	//调用支付接口
	//计算总价
	totalCost := float32(0)
	for _, item := range productItems {
		totalCost += item.GetPrice() * float32(item.GetQuantity())
	}
	chargeResp, err := rpc.PaymentClient.Charge(s.ctx, &payment.ChargeReq{
		UserId:  int32(userId),
		Amount:  totalCost,
		OrderId: placeOrderResp.GetOrder().OrderId,
	})
	if err != nil {
		klog.CtxErrorf(s.ctx, "调用支付接口失败，错误: %v,参数:req:%+v", err, req)
		return nil, errors.WithStack(err)
	}
	//调用成功返回结果
	resp = &checkout.CheckoutResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
		PaymentUrl: chargeResp.GetPaymentUrl(),
	}
	return resp, nil
}

func convertCartProductItems2OrderItems(productItems []*cart.Product) []*order.OrderItem {
	var orderItems []*order.OrderItem
	for _, item := range productItems {
		orderItem := &order.OrderItem{
			Item: &cart.CartItem{
				ProductId: item.GetId(),
				Quantity:  item.GetQuantity(),
			},
			Cost: float64(item.GetPrice() * float32(item.GetQuantity())),
		}
		orderItems = append(orderItems, orderItem)
	}
	return orderItems
}
