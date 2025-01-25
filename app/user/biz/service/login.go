package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/rpc/kitex_gen/auth"
	"douyin_mall/user/biz/dal/mysql"
	"douyin_mall/user/biz/infra/rpc"
	"douyin_mall/user/biz/model"
	"douyin_mall/user/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
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
	//if 1 == 1 {
	//	panic("test panic")
	//}
	loginUser, err := model.GetUserByEmail(mysql.DB, s.ctx, req.Email)
	if err != nil {
		// 数据库的错误
		klog.Error("用户登录失败，req: %v, err: %v", req, err)
		return nil, err
	}
	comparePwdErr := bcrypt.CompareHashAndPassword([]byte(loginUser.Password), []byte(req.Password))
	if loginUser == nil || comparePwdErr != nil {
		// 邮箱或密码错误
		resp = &user.LoginResp{
			StatusCode: 1003,
			StatusMsg:  constant.GetMsg(1003),
		}
		return
	}
	client := rpc.AuthClient
	deliveryTokenResp, err := client.DeliverTokenByRPC(s.ctx, &auth.DeliverTokenReq{UserId: loginUser.ID})
	if err != nil {
		klog.Error("调用用户授权服务发放令牌失败，UserId: %v, err: %v", loginUser.ID, err)
		return nil, err
	}
	resp = &user.LoginResp{
		StatusCode:   0,
		StatusMsg:    constant.GetMsg(0),
		AccessToken:  deliveryTokenResp.AccessToken,
		RefreshToken: deliveryTokenResp.RefreshToken,
	}
	return
}
