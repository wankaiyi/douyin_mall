package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/user/biz/dal/mysql"
	"douyin_mall/user/biz/model"
	user "douyin_mall/user/kitex_gen/user"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	if req.Password != req.ConfirmPassword {
		resp = &user.RegisterResp{StatusCode: 1000, StatusMsg: constant.GetMsg(1000)}
		return
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	newUser := &model.User{
		Username:    req.Username,
		Sex:         req.Sex,
		Email:       req.Email,
		Password:    string(hashedPassword),
		Description: req.Description,
		Avatar:      req.Avatar,
	}
	if err = model.CreateUser(mysql.DB, s.ctx, newUser); err != nil {
		klog.CtxInfof(s.ctx, "创建用户失败, err: %v", err)
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			// 用户已存在
			resp = &user.RegisterResp{StatusCode: 1002, StatusMsg: constant.GetMsg(1002)}
			return resp, nil
		} else {
			return nil, err
		}
	}

	resp = &user.RegisterResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
		UserId:     newUser.ID}
	return resp, nil
}
