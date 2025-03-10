package middleware

import (
	"context"
	"douyin_mall/api/infra/rpc"
	"douyin_mall/common/constant"
	"douyin_mall/rpc/kitex_gen/auth"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

var (
	Whitelist = map[string]struct{}{}
)

func AuthorizationMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		path := string(c.Request.URI().Path())
		if _, exist := Whitelist[path]; !exist {
			authClient := rpc.AuthClient
			verifyResp, err := authClient.VerifyTokenByRPC(ctx, &auth.VerifyTokenReq{
				RefreshToken: c.Request.Header.Get("refresh_token"),
				AccessToken:  c.Request.Header.Get("access_token"),
				Path:         path,
				Method:       string(c.Request.Method()),
			})
			if err != nil {
				hlog.CtxErrorf(ctx, "rpc权限校验失败，err: %v", err)
				c.JSON(consts.StatusOK, utils.H{
					"status_code": 500,
					"status_msg":  constant.GetMsg(500)})
				c.Abort()
			} else {
				if verifyResp.StatusCode != 0 {
					c.JSON(consts.StatusOK, verifyResp)
					c.Abort()
				}
				ctx = context.WithValue(ctx, constant.UserId, verifyResp.UserId)
			}
		}
		c.Next(ctx)

	}
}
