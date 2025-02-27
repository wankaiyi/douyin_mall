package service

import (
	"context"
	"douyin_mall/api/infra/rpc"
	"douyin_mall/common/constant"
	rpcUser "douyin_mall/rpc/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"

	user "douyin_mall/api/hertz_gen/api/user"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddReceiveAddressService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddReceiveAddressService(Context context.Context, RequestContext *app.RequestContext) *AddReceiveAddressService {
	return &AddReceiveAddressService{RequestContext: RequestContext, Context: Context}
}

func (h *AddReceiveAddressService) Run(req *user.AddReceiveAddressRequest) (resp *user.AddReceiveAddressResponse, err error) {
	ctx := h.Context
	addReceiveAddressResp, err := rpc.UserClient.AddReceiveAddress(ctx, &rpcUser.AddReceiveAddressReq{
		UserId: ctx.Value(constant.UserId).(int32),
		ReceiveAddress: &rpcUser.ReceiveAddress{
			Name:          req.Name,
			PhoneNumber:   req.PhoneNumber,
			DefaultStatus: req.DefaultStatus,
			Province:      req.Province,
			City:          req.City,
			Region:        req.Region,
			DetailAddress: req.DetailAddress,
		},
	})
	if err != nil {
		hlog.CtxErrorf(ctx, "rpc调用添加收货地址错误，req：%v，err：%v", req, err)
		return nil, errors.New("添加收货地址错误，请稍后再试")
	}
	resp = &user.AddReceiveAddressResponse{
		StatusCode: addReceiveAddressResp.StatusCode,
		StatusMsg:  addReceiveAddressResp.StatusMsg,
	}
	return resp, nil
}
