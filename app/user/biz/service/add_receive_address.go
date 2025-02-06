package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/user/biz/dal/mysql"
	"douyin_mall/user/biz/model"
	user "douyin_mall/user/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type AddReceiveAddressService struct {
	ctx context.Context
} // NewAddReceiveAddressService new AddReceiveAddressService
func NewAddReceiveAddressService(ctx context.Context) *AddReceiveAddressService {
	return &AddReceiveAddressService{ctx: ctx}
}

// Run create note info
func (s *AddReceiveAddressService) Run(req *user.AddReceiveAddressReq) (resp *user.AddReceiveAddressResp, err error) {
	ctx := s.ctx
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		if req.DefaultStatus == model.AddressDefaultStatusYes {
			existingAddr, err := model.ExistDefaultAddress(ctx, tx, req.UserId)
			if err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					klog.CtxErrorf(ctx, "查询默认地址是否存在失败，req：%v，err：%v", req, err)
					return errors.WithStack(err)
				}
			} else {
				existingAddr.DefaultStatus = model.AddressDefaultStatusNo
				err = model.UpdateAddress(ctx, mysql.DB, existingAddr)
				if err != nil {
					klog.CtxErrorf(ctx, "更新默认地址失败，req：%v，err：%v", req, err)
					return errors.WithStack(err)
				}

			}
		}
		address := model.Address{
			UserId:        req.UserId,
			Name:          req.Name,
			PhoneNumber:   req.PhoneNumber,
			DefaultStatus: int8(req.DefaultStatus),
			Province:      req.Province,
			City:          req.City,
			Region:        req.Region,
			DetailAddress: req.DetailAddress,
		}
		err = model.CreateAddress(ctx, mysql.DB, &address)
		if err != nil {
			klog.CtxErrorf(ctx, "添加收货地址失败，req：%v，err：%v", req, err)
			return errors.WithStack(err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	resp = &user.AddReceiveAddressResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
	}
	return resp, nil
}
