package service

import (
	"context"
	doubao_ai "douyin_mall/doubao_ai/kitex_gen/doubao_ai"
	"testing"
)

func TestAnalyzeSearchOrderQuestion_Run(t *testing.T) {
	ctx := context.Background()
	s := NewAnalyzeSearchOrderQuestionService(ctx)
	// init req and assert value

	req := &doubao_ai.SearchOrderQuestionReq{
		Uuid:     "test_uuid",
		UserId:   1,
		Question: "我想查询去年买的笔记本和电脑",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	req = &doubao_ai.SearchOrderQuestionReq{
		Uuid:     "test_uuid",
		UserId:   1,
		Question: "我说错了，不是笔记本和电脑，而是笔记本电脑",
	}
	resp, err = s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

}
