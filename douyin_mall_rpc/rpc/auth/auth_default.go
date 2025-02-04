package auth

import (
	"context"
	auth "douyin_mall/rpc/kitex_gen/auth"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq, callOptions ...callopt.Option) (resp *auth.DeliveryResp, err error) {
	resp, err = defaultClient.DeliverTokenByRPC(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeliverTokenByRPC call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func VerifyTokenByRPC(ctx context.Context, req *auth.VerifyTokenReq, callOptions ...callopt.Option) (resp *auth.VerifyResp, err error) {
	resp, err = defaultClient.VerifyTokenByRPC(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "VerifyTokenByRPC call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func RefreshTokenByRPC(ctx context.Context, req *auth.RefreshTokenReq, callOptions ...callopt.Option) (resp *auth.RefreshTokenResp, err error) {
	resp, err = defaultClient.RefreshTokenByRPC(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "RefreshTokenByRPC call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func RevokeTokenByRPC(ctx context.Context, req *auth.RevokeTokenReq, callOptions ...callopt.Option) (resp *auth.RevokeResp, err error) {
	resp, err = defaultClient.RevokeTokenByRPC(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "RevokeTokenByRPC call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func AddPermission(ctx context.Context, req *auth.AddPermissionReq, callOptions ...callopt.Option) (resp *auth.Empty, err error) {
	resp, err = defaultClient.AddPermission(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AddPermission call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CheckIfUserBanned(ctx context.Context, req *auth.CheckIfUserBannedReq, callOptions ...callopt.Option) (resp *auth.CheckIfUserBannedResp, err error) {
	resp, err = defaultClient.CheckIfUserBanned(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CheckIfUserBanned call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
