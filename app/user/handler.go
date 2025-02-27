package main

import (
	"context"
	"douyin_mall/user/biz/service"
	user "douyin_mall/user/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	resp, err = service.NewRegisterService(ctx).Run(req)

	return resp, err
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	resp, err = service.NewLoginService(ctx).Run(req)

	return resp, err
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *user.GetUserReq) (resp *user.GetUserResp, err error) {
	resp, err = service.NewGetUserService(ctx).Run(req)

	return resp, err
}

// UpdateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *user.UpdateUserReq) (resp *user.UpdateUserResp, err error) {
	resp, err = service.NewUpdateUserService(ctx).Run(req)

	return resp, err
}

// DeleteUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteUser(ctx context.Context, req *user.DeleteUserReq) (resp *user.DeleteUserResp, err error) {
	resp, err = service.NewDeleteUserService(ctx).Run(req)

	return resp, err
}

// GetUserRoleById implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserRoleById(ctx context.Context, req *user.GetUserRoleByIdReq) (resp *user.GetUserRoleByIdResp, err error) {
	resp, err = service.NewGetUserRoleByIdService(ctx).Run(req)

	return resp, err
}

// AddReceiveAddress implements the UserServiceImpl interface.
func (s *UserServiceImpl) AddReceiveAddress(ctx context.Context, req *user.AddReceiveAddressReq) (resp *user.AddReceiveAddressResp, err error) {
	resp, err = service.NewAddReceiveAddressService(ctx).Run(req)

	return resp, err
}

// GetReceiveAddress implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetReceiveAddress(ctx context.Context, req *user.GetReceiveAddressReq) (resp *user.GetReceiveAddressResp, err error) {
	resp, err = service.NewGetReceiveAddressService(ctx).Run(req)

	return resp, err
}
