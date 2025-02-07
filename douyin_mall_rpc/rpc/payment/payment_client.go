package payment

import (
	"context"
	payment "douyin_mall/rpc/kitex_gen/payment"

	"douyin_mall/rpc/kitex_gen/payment/paymentservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() paymentservice.Client
	Service() string
	Charge(ctx context.Context, Req *payment.ChargeReq, callOptions ...callopt.Option) (r *payment.ChargeResp, err error)
	CancelCharge(ctx context.Context, Req *payment.CancelChargeReq, callOptions ...callopt.Option) (r *payment.CancelChargeResp, err error)
	PaymentOrderRecord(ctx context.Context, Req *payment.PaymentOrderRecordReq, callOptions ...callopt.Option) (r *payment.PaymentOrderRecordResp, err error)
	PaymentTransactionRecord(ctx context.Context, Req *payment.PaymentTransactionRecordReq, callOptions ...callopt.Option) (r *payment.PaymentTransactionRecordResp, err error)
	IdempotentControl(ctx context.Context, Req *payment.IdempotentControlReq, callOptions ...callopt.Option) (r *payment.IdempotentControlResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := paymentservice.NewClient(dstService, opts...)
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
	kitexClient paymentservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() paymentservice.Client {
	return c.kitexClient
}

func (c *clientImpl) Charge(ctx context.Context, Req *payment.ChargeReq, callOptions ...callopt.Option) (r *payment.ChargeResp, err error) {
	return c.kitexClient.Charge(ctx, Req, callOptions...)
}

func (c *clientImpl) CancelCharge(ctx context.Context, Req *payment.CancelChargeReq, callOptions ...callopt.Option) (r *payment.CancelChargeResp, err error) {
	return c.kitexClient.CancelCharge(ctx, Req, callOptions...)
}

func (c *clientImpl) PaymentOrderRecord(ctx context.Context, Req *payment.PaymentOrderRecordReq, callOptions ...callopt.Option) (r *payment.PaymentOrderRecordResp, err error) {
	return c.kitexClient.PaymentOrderRecord(ctx, Req, callOptions...)
}

func (c *clientImpl) PaymentTransactionRecord(ctx context.Context, Req *payment.PaymentTransactionRecordReq, callOptions ...callopt.Option) (r *payment.PaymentTransactionRecordResp, err error) {
	return c.kitexClient.PaymentTransactionRecord(ctx, Req, callOptions...)
}

func (c *clientImpl) IdempotentControl(ctx context.Context, Req *payment.IdempotentControlReq, callOptions ...callopt.Option) (r *payment.IdempotentControlResp, err error) {
	return c.kitexClient.IdempotentControl(ctx, Req, callOptions...)
}
