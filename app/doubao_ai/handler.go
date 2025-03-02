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

// AnalyzePlaceOrderQuestion implements the DoubaoAiServiceImpl interface.
func (s *DoubaoAiServiceImpl) AnalyzePlaceOrderQuestion(ctx context.Context, req *doubao_ai.PlaceOrderQuestionReq) (resp *doubao_ai.PlaceOrderQuestionResp, err error) {
	resp, err = service.NewAnalyzePlaceOrderQuestionService(ctx).Run(req)

	return resp, err
}

// AddChatMessage implements the DoubaoAiServiceImpl interface.
func (s *DoubaoAiServiceImpl) AddChatMessage(ctx context.Context, req *doubao_ai.AddChatMessageReq) (resp *doubao_ai.AddChatMessageResp, err error) {
	resp, err = service.NewAddChatMessageService(ctx).Run(req)

	return resp, err
}
