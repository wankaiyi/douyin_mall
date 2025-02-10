package order

import (
	"context"

	"douyin_mall/api/biz/service"
	"douyin_mall/api/biz/utils"
	order "douyin_mall/api/hertz_gen/api/order"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Search .
// @router /order/search [POST]
func Search(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.OrderRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &order.OrderResponse{}
	resp, err = service.NewSearchService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// OrderInsert .
// @router /order/insert [POST]
func OrderInsert(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.OrderInsertRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &order.OrderInsertResponse{}
	resp, err = service.NewOrderInsertService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// OrderSelect .
// @router /order/select [POST]
func OrderSelect(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.OrderSelectRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &order.OrderSelectResponse{}
	resp, err = service.NewOrderSelectService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// OrderDelete .
// @router /order/delete [POST]
func OrderDelete(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.OrderDeleteRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &order.OrderDeleteResponse{}
	resp, err = service.NewOrderDeleteService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// OrderUpdate .
// @router /order/update [POST]
func OrderUpdate(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.OrderUpdateRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &order.OrderUpdateResponse{}
	resp, err = service.NewOrderUpdateService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
