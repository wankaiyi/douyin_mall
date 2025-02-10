package order

import (
	"context"

	"douyin_mall/api/biz/service"
	"douyin_mall/api/biz/utils"
	order "douyin_mall/api/hertz_gen/api/order"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CartInsert .
// @router /cart/insert [POST]
func CartInsert(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.CartInsertRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &order.CartInsertResponse{}
	resp, err = service.NewCartInsertService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// CartSelect .
// @router /cart/select [POST]
func CartSelect(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.CartSelectRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &order.CartSelectResponse{}
	resp, err = service.NewCartSelectService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// CartDelete .
// @router /cart/delete [POST]
func CartDelete(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.CartDeleteRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &order.CartDeleteResponse{}
	resp, err = service.NewCartDeleteService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// CartUpdate .
// @router /cart/update [POST]
func CartUpdate(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.CartUpdateRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &order.CartUpdateResponse{}
	resp, err = service.NewCartUpdateService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
