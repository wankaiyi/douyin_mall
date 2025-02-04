package service

import (
	"context"
	auth "douyin_mall/auth/kitex_gen/auth"
	"testing"
)

func TestCheckIfUserBanned_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCheckIfUserBannedService(ctx)
	// init req and assert value

	req := &auth.CheckIfUserBannedReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
