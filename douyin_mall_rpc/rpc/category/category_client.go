package category

import (
	"context"
	product "douyin_mall/rpc/kitex_gen/product"

	"douyin_mall/rpc/kitex_gen/product/categoryservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() categoryservice.Client
	Service() string
	SelectCategory(ctx context.Context, Req *product.CategorySelectReq, callOptions ...callopt.Option) (r *product.CategorySelectResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := categoryservice.NewClient(dstService, opts...)
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
	kitexClient categoryservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() categoryservice.Client {
	return c.kitexClient
}

func (c *clientImpl) SelectCategory(ctx context.Context, Req *product.CategorySelectReq, callOptions ...callopt.Option) (r *product.CategorySelectResp, err error) {
	return c.kitexClient.SelectCategory(ctx, Req, callOptions...)
}
