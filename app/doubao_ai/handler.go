package main

import (
	"context"
	"douyin_mall/doubao_ai/biz/service"
	doubao_ai "douyin_mall/doubao_ai/kitex_gen/doubao_ai"
)

// DoubaoAiServiceImpl implements the last service interface defined in the IDL.
type DoubaoAiServiceImpl struct{}

// AnalyzeSearchOrderQuestion implements the DoubaoAiServiceImpl interface.
func (s *DoubaoAiServiceImpl) AnalyzeSearchOrderQuestion(ctx context.Context, req *doubao_ai.SearchOrderQuestionReq) (resp *doubao_ai.SearchOrderQuestionResp, err error) {
	resp, err = service.NewAnalyzeSearchOrderQuestionService(ctx).Run(req)

	return resp, err
}
