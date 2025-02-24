package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/rpc/kitex_gen/auth"
	"douyin_mall/user/biz/dal/mysql"
	"douyin_mall/user/biz/model"
	"douyin_mall/user/infra/rpc"
	"douyin_mall/user/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	loginUser, err := model.GetUserByUsername(mysql.DB, s.ctx, req.Username)
	if err != nil {
		// 数据库的错误
		klog.CtxErrorf(s.ctx, "用户登录失败，req: %v, err: %v", req, err)
		return nil, errors.WithStack(err)
	}

	bannedResp, err := rpc.AuthClient.CheckIfUserBanned(s.ctx, &auth.CheckIfUserBannedReq{
		UserId: loginUser.ID,
	})
	if err != nil {
		klog.CtxErrorf(s.ctx, "rpc鉴权服务查询用户是否封禁失败，userId: %d, err: %v", loginUser.ID, err)
		return nil, errors.WithStack(err)
	}
	if bannedResp.IsBanned {
		resp = &user.LoginResp{
			StatusCode: 1010,
			StatusMsg:  constant.GetMsg(1010),
		}
		return resp, nil
	}

	comparePwdErr := bcrypt.CompareHashAndPassword([]byte(loginUser.Password), []byte(req.Password))
	if loginUser == nil || comparePwdErr != nil {
		// 用户名或密码错误
		resp = &user.LoginResp{
			StatusCode: 1003,
			StatusMsg:  constant.GetMsg(1003),
		}
		return
	}
	client := rpc.AuthClient
	deliveryTokenResp, err := client.DeliverTokenByRPC(s.ctx, &auth.DeliverTokenReq{
		UserId: loginUser.ID,
		Role:   string(loginUser.Role),
	})
	if err != nil {
		klog.CtxErrorf(s.ctx, "调用用户授权服务发放令牌失败，UserId: %v, err: %v", loginUser.ID, err)
		return nil, errors.WithStack(err)
	}
	resp = &user.LoginResp{
		StatusCode:   0,
		StatusMsg:    constant.GetMsg(0),
		AccessToken:  deliveryTokenResp.AccessToken,
		RefreshToken: deliveryTokenResp.RefreshToken,
	}
	return
}
