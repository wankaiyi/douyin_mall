package doubao_ai

import (
	"context"
	doubao_ai "douyin_mall/rpc/kitex_gen/doubao_ai"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func AnalyzeSearchOrderQuestion(ctx context.Context, req *doubao_ai.SearchOrderQuestionReq, callOptions ...callopt.Option) (resp *doubao_ai.SearchOrderQuestionResp, err error) {
	resp, err = defaultClient.AnalyzeSearchOrderQuestion(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AnalyzeSearchOrderQuestion call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func AnalyzePlaceOrderQuestion(ctx context.Context, req *doubao_ai.PlaceOrderQuestionReq, callOptions ...callopt.Option) (resp *doubao_ai.PlaceOrderQuestionResp, err error) {
	resp, err = defaultClient.AnalyzePlaceOrderQuestion(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AnalyzePlaceOrderQuestion call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func AddChatMessage(ctx context.Context, req *doubao_ai.AddChatMessageReq, callOptions ...callopt.Option) (resp *doubao_ai.AddChatMessageResp, err error) {
	resp, err = defaultClient.AddChatMessage(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AddChatMessage call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
