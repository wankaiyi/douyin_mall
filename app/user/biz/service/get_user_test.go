package service

import (
	"context"
	user "douyin_mall/user/kitex_gen/user"
	"testing"
)

func TestGetUser_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetUserService(ctx)
	// init req and assert value

	req := &user.GetUserReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
