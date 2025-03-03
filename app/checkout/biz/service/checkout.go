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
	"douyin_mall/rpc/kitex_gen/product"
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
	ctx := s.ctx
	userId := req.UserId
	//得到互斥锁
	rsync := redsync.GetRedsync()
	mutexName := "checkout_mutex_" + strconv.FormatInt(int64(userId), 10)
	mutex := rsync.NewMutex(mutexName)
	//加锁
	if err := mutex.Lock(); err != nil {
		klog.CtxErrorf(ctx, "获取互斥锁失败，lock failed, :%d,err:%s", userId, err.Error())
		return nil, err
	}
	defer mutex.Unlock()

	//获得购物车信息
	getCartReq := &cart.GetCartReq{UserId: int32(userId)}
	cartResp, err := rpc.CartClient.GetCart(ctx, getCartReq)
	if err != nil {
		klog.CtxErrorf(ctx, "获取购物车信息失败:%v,参数:req:%+v", err, getCartReq)
		return nil, errors.WithStack(err)
	}
	klog.CtxInfof(ctx, "获取购物车信息成功: req: %v, resp: %v", getCartReq, cartResp)
	productItems := cartResp.Products

	//判断库存是否充足
	//锁定库存
	lockQuantityResponse, _ := rpc.ProductClient.LockProductQuantity(ctx, &product.ProductLockQuantityRequest{Products: convertProductItems(productItems)})

	if lockQuantityResponse.GetStatusCode() != 0 {
		klog.CtxErrorf(ctx, "锁定库存失败: %v", lockQuantityResponse.GetStatusMsg())
		return &checkout.CheckoutResp{
			StatusCode: lockQuantityResponse.GetStatusCode(),
			StatusMsg:  lockQuantityResponse.GetStatusMsg(),
		}, nil
	}

	//得到用户地址
	address := &order.Address{
		Name:          req.Address.ReceiveAddress.GetName(),
		PhoneNumber:   req.Address.ReceiveAddress.GetPhoneNumber(),
		DetailAddress: req.Address.ReceiveAddress.GetDetailAddress(),
		City:          req.Address.ReceiveAddress.GetCity(),
		Province:      req.Address.ReceiveAddress.GetProvince(),
		Region:        req.Address.ReceiveAddress.GetRegion(),
	}

	//计算总价
	totalCost := float32(0)
	for _, item := range productItems {
		totalCost += item.GetPrice() * float32(item.GetQuantity())
	}

	//创建订单信息
	placeOrderReq := &order.PlaceOrderReq{
		UserId:     int32(userId),
		Address:    address,
		OrderItems: convertCartProductItems2OrderItems(productItems),
		TotalCost:  float64(totalCost),
	}
	placeOrderResp, err := rpc.OrderClient.PlaceOrder(ctx, placeOrderReq)
	if err != nil {
		klog.CtxErrorf(ctx, "创建订单失败: %v ,参数:req:%+v", err, placeOrderResp)
		return nil, errors.WithStack(err)
	}
	klog.CtxInfof(ctx, "创建订单成功: req: %v, resp: %v", placeOrderResp, placeOrderResp)

	//调用支付接口
	chargeReq := &payment.ChargeReq{
		UserId:  int32(userId),
		Amount:  totalCost,
		OrderId: placeOrderResp.GetOrder().OrderId,
	}
	chargeResp, err := rpc.PaymentClient.Charge(ctx, chargeReq)
	if err != nil {
		klog.CtxErrorf(ctx, "调用支付接口失败，错误: %v,参数:req:%+v", err, chargeReq)
		return nil, errors.WithStack(err)
	}
	klog.CtxInfof(ctx, "调用支付接口成功: req: %v, resp: %v", chargeReq, chargeResp)
	//清空购物车
	emptyCartReq := &cart.EmptyCartReq{UserId: int32(userId)}
	_, err = rpc.CartClient.EmptyCart(ctx, emptyCartReq)
	if err != nil {
		klog.CtxErrorf(ctx, "清空购物车失败: %v,参数:req:%+v", err, emptyCartReq)
		return nil, errors.WithStack(err)
	}
	klog.CtxInfof(ctx, "清空购物车成功: req: %v", emptyCartReq)

	//调用成功返回结果
	resp = &checkout.CheckoutResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
		PaymentUrl: chargeResp.GetPaymentUrl(),
	}
	return resp, nil
}

func convertProductItems(productItems []*cart.Product) (productsLockQuantity []*product.ProductLockQuantity) {
	productsLockQuantity = make([]*product.ProductLockQuantity, len(productItems))
	for _, item := range productItems {
		productLockQuantity := &product.ProductLockQuantity{
			Id:       int64(item.GetId()),
			Quantity: int64(item.GetQuantity()),
		}
		productsLockQuantity = append(productsLockQuantity, productLockQuantity)
	}
	return

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
