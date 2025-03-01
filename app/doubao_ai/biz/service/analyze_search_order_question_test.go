package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/doubao_ai/biz/dal"
	doubao_ai "douyin_mall/doubao_ai/kitex_gen/doubao_ai"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/google/uuid"
	"testing"
)

func TestAnalyzeSearchOrderQuestion_Run(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	s := NewAnalyzeSearchOrderQuestionService(ctx)
	// init req and assert value

	testUuid := uuid.New().String()
	req := &doubao_ai.SearchOrderQuestionReq{
		Uuid:     testUuid,
		UserId:   1,
		Question: "我想查询去年买的笔记本和电脑",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
	expectedResp := &doubao_ai.SearchOrderQuestionResp{
		StatusCode:  0,
		StatusMsg:   constant.GetMsg(0),
		StartTime:   "2024-01-01 00:00:00",
		EndTime:     "2024-12-31 23:59:59",
		SearchTerms: []string{"笔记本", "电脑"}}
	assert.DeepEqual(t,
		expectedResp.String(),
		resp.String())

	req = &doubao_ai.SearchOrderQuestionReq{
		Uuid:     testUuid,
		UserId:   1,
		Question: "我说错了，不是笔记本和电脑，而是笔记本电脑",
	}
	resp, err = s.Run(req)
	expectedResp = &doubao_ai.SearchOrderQuestionResp{
		StatusCode:  0,
		StatusMsg:   constant.GetMsg(0),
		StartTime:   "2024-01-01 00:00:00",
		EndTime:     "2024-12-31 23:59:59",
		SearchTerms: []string{"笔记本电脑"}}
	assert.DeepEqual(t,
		expectedResp.String(),
		resp.String())

	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

}
