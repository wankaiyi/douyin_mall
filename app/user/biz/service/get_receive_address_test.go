package service

import (
	"context"
	"douyin_mall/user/biz/dal"
	user "douyin_mall/user/kitex_gen/user"
	"testing"
)

func TestGetReceiveAddress_Run(t *testing.T) {
	ctx := context.Background()
	dal.Init()
	s := NewGetReceiveAddressService(ctx)
	// init req and assert value

	req := &user.GetReceiveAddressReq{UserId: 13}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
