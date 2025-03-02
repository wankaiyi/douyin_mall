package service

import (
	"context"
	doubao_ai "douyin_mall/doubao_ai/kitex_gen/doubao_ai"
	"testing"
)

func TestAddChatMessage_Run(t *testing.T) {
	ctx := context.Background()
	s := NewAddChatMessageService(ctx)
	// init req and assert value

	req := &doubao_ai.AddChatMessageReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
