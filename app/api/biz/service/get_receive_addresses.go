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

type GetReceiveAddressesService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetReceiveAddressesService(Context context.Context, RequestContext *app.RequestContext) *GetReceiveAddressesService {
	return &GetReceiveAddressesService{RequestContext: RequestContext, Context: Context}
}

func (h *GetReceiveAddressesService) Run(req *user.Empty) (resp *user.GetReceiveAddressesResponse, err error) {
	ctx := h.Context
	getAddressesResp, err := rpc.UserClient.GetReceiveAddress(ctx, &rpcUser.GetReceiveAddressReq{
		UserId: ctx.Value(constant.UserId).(int32),
	})
	if err != nil {
		hlog.CtxErrorf(ctx, "rpc调用查询收货地址失败，req：%v，err：%v", req, err)
		return nil, errors.New("查询收货地址失败")
	}
	var receiveAddresses []*user.ReceiveAddress
	for _, address := range getAddressesResp.ReceiveAddress {
		receiveAddresses = append(receiveAddresses, &user.ReceiveAddress{
			Id:            address.Id,
			Name:          address.Name,
			PhoneNumber:   address.PhoneNumber,
			DefaultStatus: address.DefaultStatus,
			Province:      address.Province,
			City:          address.City,
			Region:        address.Region,
			DetailAddress: address.DetailAddress,
		})
	}
	return &user.GetReceiveAddressesResponse{
		Addresses:  receiveAddresses,
		StatusCode: getAddressesResp.StatusCode,
		StatusMsg:  getAddressesResp.StatusMsg,
	}, nil
}
