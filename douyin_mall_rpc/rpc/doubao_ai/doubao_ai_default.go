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
