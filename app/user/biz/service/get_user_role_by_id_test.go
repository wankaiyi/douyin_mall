package service

import (
	"context"
	user "douyin_mall/user/kitex_gen/user"
	"testing"
)

func TestGetUserRoleById_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetUserRoleByIdService(ctx)
	// init req and assert value

	req := &user.GetUserRoleByIdReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
