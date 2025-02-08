package category

import (
	"context"

	"douyin_mall/api/biz/service"
	"douyin_mall/api/biz/utils"
	category "douyin_mall/api/hertz_gen/api/category"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CategorySelect .
// @router /category/select [POST]
func CategorySelect(ctx context.Context, c *app.RequestContext) {
	var err error
	var req category.CategorySelectRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &category.CategorySelectResponse{}
	resp, err = service.NewCategorySelectService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
