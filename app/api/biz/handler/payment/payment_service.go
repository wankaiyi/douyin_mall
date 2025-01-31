package payment

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	"douyin_mall/api/biz/service"
	"douyin_mall/api/biz/utils"
	payment "douyin_mall/api/hertz_gen/api/payment"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Charge .
// @router /payment/charge [POST]
func Charge(ctx context.Context, c *app.RequestContext) {
	var err error
	var req payment.ChargeRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &payment.ChargeResponse{}
	resp, err = service.NewChargeService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// Notify .
// @router /payment/notify [POST]
func Notify(ctx context.Context, c *app.RequestContext) {
	var err error
	var req payment.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		return
	}

	err = service.NewNotifyService(ctx, c).Run(c)

	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		return
	}

}

// Cancel .
// @router /payment/cancel [POST]
func Cancel(ctx context.Context, c *app.RequestContext) {
	var err error
	var req payment.CancelRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewCancelService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
