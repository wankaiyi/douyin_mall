package service

import (
	"context"
	"douyin_mall/checkout/infra/kafka/model"
	"douyin_mall/checkout/infra/kafka/producer"
	"douyin_mall/checkout/infra/rpc"
	checkout "douyin_mall/checkout/kitex_gen/checkout"
	"douyin_mall/checkout/kitex_gen/user"
	"douyin_mall/rpc/kitex_gen/cart"
	"douyin_mall/rpc/kitex_gen/order"
	"douyin_mall/rpc/kitex_gen/payment"
	"douyin_mall/rpc/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
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
	ctx := s.ctx
	klog.CtxInfof(ctx, "收到结算商品请求, req: %v", req)
	getReceiveAddressResp, err := rpc.UserClient.GetReceiveAddress(ctx, &user.GetReceiveAddressReq{
		UserId: req.UserId,
	})
	if err != nil {
		klog.CtxErrorf(ctx, "获取用户地址失败rpc接口调用失败, error: %v", err)
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
		klog.CtxErrorf(ctx, "用户地址不存在, address_id: %d", req.AddressId)
		return nil, errors.New("用户地址不存在")
	}

	// 计算商品的总价
	var productIds []int64
	for _, productItem := range req.Items {
		productIds = append(productIds, int64(productItem.ProductId))
	}
	productListReq := &product.SelectProductListReq{
		Ids: productIds,
	}
	productListResp, err := rpc.ProductClient.SelectProductList(ctx, productListReq)
	if err != nil {
		klog.CtxErrorf(ctx, "获取商品信息失败rpc接口调用失败, req: %v, error: %v", req, err)
		return nil, err
	}
	klog.CtxInfof(ctx, "获取商品信息成功, resp: %v", productListResp)
	productMap := make(map[int64]*product.Product)
	for _, p := range productListResp.Products {
		productMap[p.Id] = p
	}

	var cost float32
	var orderItems []*order.OrderItem
	var productItems []*cart.Product
	for _, productItem := range req.Items {
		p, ok := productMap[int64(productItem.ProductId)]
		if !ok {
			klog.CtxErrorf(ctx, "商品不存在, product_id: %d", productItem.ProductId)
			return nil, errors.New("商品不存在")
		}

		productItems = append(productItems, &cart.Product{
			Id:       productItem.ProductId,
			Quantity: productItem.Quantity,
		})

		cost += float32(productItem.Quantity) * p.Price
		orderItems = append(orderItems, &order.OrderItem{
			Item: &cart.CartItem{
				ProductId: int32(p.Id),
				Quantity:  productItem.Quantity,
			},
			Cost: float64(float32(productItem.Quantity) * p.Price),
		})
	}

	lockQuantityRequest := &product.ProductLockQuantityRequest{Products: convertProductItems(productItems)}
	lockQuantityResponse, err := rpc.ProductClient.LockProductQuantity(ctx, lockQuantityRequest)
	if err != nil {
		klog.CtxErrorf(ctx, "锁定库存失败，req: %v, err: %v", lockQuantityRequest, err)
		return nil, errors.WithStack(err)
	}
	if lockQuantityResponse.GetStatusCode() != 0 {
		klog.CtxInfof(ctx, "锁定库存失败: %v", lockQuantityResponse.GetStatusMsg())
		return &checkout.CheckoutProductItemsResp{
			StatusCode: lockQuantityResponse.GetStatusCode(),
			StatusMsg:  lockQuantityResponse.GetStatusMsg(),
		}, nil
	}
	uuidStr := uuid.New().String()

	var lockProductItems []model.ProductItem
	for _, item := range productItems {
		lockProductItems = append(lockProductItems, model.ProductItem{
			ProductID: item.Id,
			Quantity:  item.Quantity,
		})
	}
	producer.SendDelayStockCompensationMessage(ctx, uuidStr, lockProductItems)

	//创建订单
	placeOrderResp, err := rpc.OrderClient.PlaceOrder(ctx, &order.PlaceOrderReq{
		UserId: req.UserId,
		Address: &order.Address{
			Name:          targetAddress.Name,
			PhoneNumber:   targetAddress.PhoneNumber,
			Province:      targetAddress.Province,
			City:          targetAddress.City,
			Region:        targetAddress.Region,
			DetailAddress: targetAddress.DetailAddress,
		},
		OrderItems: orderItems,
		TotalCost:  float64(cost),
		Uuid:       uuidStr,
	})
	if err != nil {
		klog.CtxErrorf(ctx, "创建订单失败rpc接口调用失败, error: %v", err)
		return nil, errors.WithStack(err)
	}
	orderId := placeOrderResp.Order.OrderId

	chargeResp, err := rpc.PaymentClient.Charge(ctx, &payment.ChargeReq{
		OrderId: orderId,
		Amount:  cost,
	})
	if err != nil {
		klog.CtxErrorf(ctx, "支付失败rpc接口调用失败, error: %v", err)
		return nil, errors.WithStack(err)
	}

	resp = &checkout.CheckoutProductItemsResp{
		StatusCode: chargeResp.StatusCode,
		StatusMsg:  chargeResp.StatusMsg,
		PaymentUrl: chargeResp.PaymentUrl,
	}
	return
}
