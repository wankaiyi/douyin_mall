package order

import (
	"context"

	"douyin_mall/api/biz/service"
	"douyin_mall/api/biz/utils"
	order "douyin_mall/api/hertz_gen/api/order"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// ListOrder .
// @router /order/list [GET]
func ListOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.ListOrderRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &order.ListOrderResponse{}
	resp, err = service.NewListOrderService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
