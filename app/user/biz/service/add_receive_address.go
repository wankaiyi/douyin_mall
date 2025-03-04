package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/user/biz/dal/mysql"
	"douyin_mall/user/biz/dal/redis"
	"douyin_mall/user/biz/model"
	user "douyin_mall/user/kitex_gen/user"
	redisUtils "douyin_mall/user/utils/redis"
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
	addr := req.ReceiveAddress
	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		userId := req.UserId
		if addr.DefaultStatus == model.AddressDefaultStatusYes {
			existingAddr, err := model.ExistDefaultAddress(ctx, tx, userId)
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
			UserId:        userId,
			Name:          addr.Name,
			PhoneNumber:   addr.PhoneNumber,
			DefaultStatus: int8(addr.DefaultStatus),
			Province:      addr.Province,
			City:          addr.City,
			Region:        addr.Region,
			DetailAddress: addr.DetailAddress,
		}
		err = model.CreateAddress(ctx, mysql.DB, &address)
		if err != nil {
			klog.CtxErrorf(ctx, "添加收货地址失败，req：%v，err：%v", req, err)
			return errors.WithStack(err)
		}
		// 缓存删除失败则回滚事务，防止数据不一致
		err = redis.RedisClient.Del(ctx, redisUtils.GetUserAddressKey(userId)).Err()
		if err != nil {
			klog.CtxErrorf(ctx, "删除用户地址缓存失败，req：%v，err：%v", req, err)
			return errors.WithStack(err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	klog.CtxInfof(ctx, "添加收货地址成功，req：%v", req)
	resp = &user.AddReceiveAddressResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
	}
	return resp, nil
}
