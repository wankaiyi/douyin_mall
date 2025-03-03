package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/payment/biz/dal/alipay"
	"douyin_mall/payment/infra/kafka/producer"
	"douyin_mall/payment/infra/rpc"
	payment "douyin_mall/payment/kitex_gen/payment"
	"douyin_mall/rpc/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
	"strconv"
)

type NotifyPaymentService struct {
	ctx context.Context
} // NewNotifyPaymentService new NotifyPaymentService
func NewNotifyPaymentService(ctx context.Context) *NotifyPaymentService {
	return &NotifyPaymentService{ctx: ctx}
}

// Run create note info
func (s *NotifyPaymentService) Run(req *payment.NotifyPaymentReq) (resp *payment.NotifyPaymentResp, err error) {
	klog.Infof("收到支付宝支付成功回调，req：%+v", req)
	ctx := s.ctx
	notifyData := req.NotifyData
	//alipayTradeNo := notifyData["alipayTradeNo"]
	tradeStatus := notifyData["tradeStatus"]
	//alipayAmount := notifyData["alipayAmount"]
	orderId := notifyData["orderId"]
	userId, _ := strconv.ParseInt(notifyData["userId"], 10, 64)

	getOrderReq := &order.GetOrderReq{
		OrderId: orderId,
	}
	getOrderResp, err := rpc.OrderClient.GetOrder(ctx, getOrderReq)
	if err != nil {
		klog.Errorf("支付回调rpc查询订单失败，req: %v,err: %v", getOrderReq, err)
		return nil, errors.WithStack(err)
	}
	if getOrderResp.Order.Status == constant.OrderStatusUnpaid {
		// 检查支付状态，如果支付成功，则更新订单状态
		if tradeStatus != "TRADE_SUCCESS" {
			klog.CtxErrorf(ctx, "orderId:%s,pay filed,tradeStatus:%s", orderId, tradeStatus)
			resp = &payment.NotifyPaymentResp{
				StatusCode: 5006,
				StatusMsg:  constant.GetMsg(5006),
			}
			return resp, nil
		}
		//更新订单状态
		markOrderPaidResp, err := rpc.OrderClient.MarkOrderPaid(ctx, &order.MarkOrderPaidReq{
			OrderId: orderId,
			UserId:  int32(userId),
		})
		if err != nil {
			klog.CtxErrorf(ctx, "orderId:%s,更新订单状态失败,err:%s", orderId, err.Error())
			//退款
			refund(ctx, orderId, getOrderResp)
			return nil, errors.WithStack(err)
		}
		if markOrderPaidResp.StatusCode != 0 {
			klog.CtxInfof(ctx, "修改订单状态失败,orderId: %s,resp: %v", orderId, markOrderPaidResp)
			return &payment.NotifyPaymentResp{
				StatusCode: markOrderPaidResp.StatusCode,
				StatusMsg:  constant.GetMsg(int(markOrderPaidResp.StatusCode)),
			}, nil
		}
		//发送支付成功消息
		producer.SendPaymentSuccessOrderIdMsg(ctx, orderId)
	}

	resp = &payment.NotifyPaymentResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
	}
	return resp, nil

}

func refund(ctx context.Context, orderId string, getOrderResp *order.GetOrderResp) {
	//退款
	result, refundErr := alipay.Refund(ctx, orderId, getOrderResp.Order.Cost)
	if refundErr != nil {
		klog.CtxErrorf(ctx, "orderId:%s,退款失败,err:%s", orderId, refundErr.Error())
	}
	if !result {
		klog.CtxErrorf(ctx, "orderId:%s,退款失败", orderId)
	}
	klog.CtxInfof(ctx, "orderId:%s,退款成功", orderId)
}
