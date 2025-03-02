package doubao_ai

import (
	"context"
	doubao_ai "douyin_mall/rpc/kitex_gen/doubao_ai"

	"douyin_mall/rpc/kitex_gen/doubao_ai/doubaoaiservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() doubaoaiservice.Client
	Service() string
	AnalyzeSearchOrderQuestion(ctx context.Context, Req *doubao_ai.SearchOrderQuestionReq, callOptions ...callopt.Option) (r *doubao_ai.SearchOrderQuestionResp, err error)
	AnalyzePlaceOrderQuestion(ctx context.Context, Req *doubao_ai.PlaceOrderQuestionReq, callOptions ...callopt.Option) (r *doubao_ai.PlaceOrderQuestionResp, err error)
	AddChatMessage(ctx context.Context, Req *doubao_ai.AddChatMessageReq, callOptions ...callopt.Option) (r *doubao_ai.AddChatMessageResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := doubaoaiservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient doubaoaiservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() doubaoaiservice.Client {
	return c.kitexClient
}

func (c *clientImpl) AnalyzeSearchOrderQuestion(ctx context.Context, Req *doubao_ai.SearchOrderQuestionReq, callOptions ...callopt.Option) (r *doubao_ai.SearchOrderQuestionResp, err error) {
	return c.kitexClient.AnalyzeSearchOrderQuestion(ctx, Req, callOptions...)
}

func (c *clientImpl) AnalyzePlaceOrderQuestion(ctx context.Context, Req *doubao_ai.PlaceOrderQuestionReq, callOptions ...callopt.Option) (r *doubao_ai.PlaceOrderQuestionResp, err error) {
	return c.kitexClient.AnalyzePlaceOrderQuestion(ctx, Req, callOptions...)
}

func (c *clientImpl) AddChatMessage(ctx context.Context, Req *doubao_ai.AddChatMessageReq, callOptions ...callopt.Option) (r *doubao_ai.AddChatMessageResp, err error) {
	return c.kitexClient.AddChatMessage(ctx, Req, callOptions...)
}
