package service

import (
	"context"
	"douyin_mall/auth/biz/dal/redis"
	"douyin_mall/auth/casbin"
	auth "douyin_mall/auth/kitex_gen/auth"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
)

type AddPermissionService struct {
	ctx context.Context
} // NewAddPermissionService new AddPermissionService
func NewAddPermissionService(ctx context.Context) *AddPermissionService {
	return &AddPermissionService{ctx: ctx}
}

// Run create note info
func (s *AddPermissionService) Run(req *auth.AddPermissionReq) (resp *auth.Empty, err error) {
	err = casbin.AddPolicy(req.Role, req.Path, req.Method)
	if err != nil {
		klog.Errorf("添加 Casbin 权限策略失败，role: %s, path: %s, method: %s, err: %v", req.Role, req.Path, req.Method, err)
	}
	redis.RedisClient.Publish(s.ctx, "casbin_policy_updates", fmt.Sprintf("新增权限策略，role: %s, path: %s, method: %s", req.Role, req.Path, req.Method))
	return nil, err
}
