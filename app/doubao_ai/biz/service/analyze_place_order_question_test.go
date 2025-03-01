package service

import (
	"context"
	"douyin_mall/doubao_ai/biz/dal"
	doubao_ai "douyin_mall/doubao_ai/kitex_gen/doubao_ai"
	"testing"
)

func TestAnalyzePlaceOrderQuestion_Run(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	s := NewAnalyzePlaceOrderQuestionService(ctx)
	// init req and assert value

	req := &doubao_ai.PlaceOrderQuestionReq{
		UserId:   13,
		Uuid:     "123456",
		Question: "你好，请问有什么可以帮到您的？",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
