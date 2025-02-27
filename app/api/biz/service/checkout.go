package service

import (
	"context"
	"douyin_mall/api/infra/rpc"
	"douyin_mall/rpc/kitex_gen/user"
	"errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	checkout "douyin_mall/api/hertz_gen/api/checkout"
	rpcCheckout "douyin_mall/rpc/kitex_gen/checkout"
	"github.com/cloudwego/hertz/pkg/app"
)

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	address := &user.AddReceiveAddressReq{
		UserId:        h.Context.Value("user_id").(int32),
		Name:          req.Address.GetName(),
		PhoneNumber:   req.Address.GetPhoneNumber(),
		City:          req.Address.GetCity(),
		Province:      req.GetAddress().GetProvince(),
		Region:        req.Address.Region,
		DefaultStatus: req.Address.GetDefaultStatus(),
		DetailAddress: req.Address.GetDetailAddress(),
	}
	checkoutResp, err := rpc.CheckoutClient.Checkout(h.Context, &rpcCheckout.CheckoutReq{
		UserId:  uint32(h.Context.Value("user_id").(int32)),
		Address: address,
	})
	if err != nil {
		hlog.CtxErrorf(h.Context, "结算失败接口rpc调用失败，err = %v", err)
		return nil, errors.New("结算失败,请稍后再试")
	}

	return &checkout.CheckoutResp{
		StatusCode: checkoutResp.GetStatusCode(),
		StatusMsg:  checkoutResp.GetStatusMsg(),
		PaymentUrl: checkoutResp.GetPaymentUrl(),
	}, nil
}
