package cart

import (
	"context"
	cart "douyin_mall/rpc/kitex_gen/cart"

	"douyin_mall/rpc/kitex_gen/cart/cartservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() cartservice.Client
	Service() string
	AddItem(ctx context.Context, Req *cart.AddItemReq, callOptions ...callopt.Option) (r *cart.AddItemResp, err error)
	GetCart(ctx context.Context, Req *cart.GetCartReq, callOptions ...callopt.Option) (r *cart.GetCartResp, err error)
	EmptyCart(ctx context.Context, Req *cart.EmptyCartReq, callOptions ...callopt.Option) (r *cart.EmptyCartResp, err error)
	SearchCarts(ctx context.Context, Req *cart.SearchCartsReq, callOptions ...callopt.Option) (r *cart.SearchCartsResp, err error)
	InsertCart(ctx context.Context, Req *cart.InsertCartReq, callOptions ...callopt.Option) (r *cart.InsertCartResp, err error)
	SelectCart(ctx context.Context, Req *cart.SelectCartReq, callOptions ...callopt.Option) (r *cart.SelectCartResp, err error)
	DeleteCart(ctx context.Context, Req *cart.DeleteCartReq, callOptions ...callopt.Option) (r *cart.DeleteCartResp, err error)
	UpdateCart(ctx context.Context, Req *cart.UpdateCartReq, callOptions ...callopt.Option) (r *cart.UpdateCartResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := cartservice.NewClient(dstService, opts...)
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
	kitexClient cartservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() cartservice.Client {
	return c.kitexClient
}

func (c *clientImpl) AddItem(ctx context.Context, Req *cart.AddItemReq, callOptions ...callopt.Option) (r *cart.AddItemResp, err error) {
	return c.kitexClient.AddItem(ctx, Req, callOptions...)
}

func (c *clientImpl) GetCart(ctx context.Context, Req *cart.GetCartReq, callOptions ...callopt.Option) (r *cart.GetCartResp, err error) {
	return c.kitexClient.GetCart(ctx, Req, callOptions...)
}

func (c *clientImpl) EmptyCart(ctx context.Context, Req *cart.EmptyCartReq, callOptions ...callopt.Option) (r *cart.EmptyCartResp, err error) {
	return c.kitexClient.EmptyCart(ctx, Req, callOptions...)
}

func (c *clientImpl) SearchCarts(ctx context.Context, Req *cart.SearchCartsReq, callOptions ...callopt.Option) (r *cart.SearchCartsResp, err error) {
	return c.kitexClient.SearchCarts(ctx, Req, callOptions...)
}

func (c *clientImpl) InsertCart(ctx context.Context, Req *cart.InsertCartReq, callOptions ...callopt.Option) (r *cart.InsertCartResp, err error) {
	return c.kitexClient.InsertCart(ctx, Req, callOptions...)
}

func (c *clientImpl) SelectCart(ctx context.Context, Req *cart.SelectCartReq, callOptions ...callopt.Option) (r *cart.SelectCartResp, err error) {
	return c.kitexClient.SelectCart(ctx, Req, callOptions...)
}

func (c *clientImpl) DeleteCart(ctx context.Context, Req *cart.DeleteCartReq, callOptions ...callopt.Option) (r *cart.DeleteCartResp, err error) {
	return c.kitexClient.DeleteCart(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateCart(ctx context.Context, Req *cart.UpdateCartReq, callOptions ...callopt.Option) (r *cart.UpdateCartResp, err error) {
	return c.kitexClient.UpdateCart(ctx, Req, callOptions...)
}
