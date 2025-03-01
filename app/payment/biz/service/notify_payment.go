package service

import (
	"context"
	"douyin_mall/common/constant"
	redsync "douyin_mall/payment/biz/dal/red_sync"
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
	// Finish your business logic.
	notifyData := req.NotifyData
	//alipayTradeNo := notifyData["alipayTradeNo"]
	tradeStatus := notifyData["tradeStatus"]
	//alipayAmount := notifyData["alipayAmount"]
	orderId := notifyData["orderId"]
	userId, _ := strconv.ParseInt(notifyData["userId"], 10, 64)
	//得到互斥锁
	rsync := redsync.GetRedsync()
	mutexName := "order_" + orderId
	mutex := rsync.NewMutex(mutexName)

	//加锁
	if err := mutex.Lock(); err != nil {
		klog.CtxErrorf(s.ctx, "获取互斥锁失败，lock failed,orderId:%s,,err:%s", orderId, err.Error())
		return nil, err
	}
	defer mutex.Unlock()

	getOrderResp, err := rpc.OrderClient.GetOrder(s.ctx, &order.GetOrderReq{
		OrderId: orderId,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if getOrderResp.Order.Status == constant.OrderStatusUnpaid {
		// 检查支付状态，如果支付成功，则更新订单状态
		if tradeStatus != "TRADE_SUCCESS" {
			klog.CtxErrorf(s.ctx, "orderId:%s,pay filed,tradeStatus:%s", orderId, tradeStatus)
			resp = &payment.NotifyPaymentResp{
				StatusCode: 5006,
				StatusMsg:  constant.GetMsg(5006),
			}
			return resp, nil
		}
		//更新订单状态
		markOrderPaidResp, err := rpc.OrderClient.MarkOrderPaid(s.ctx, &order.MarkOrderPaidReq{
			OrderId: orderId,
			UserId:  int32(userId),
		})
		if err != nil {
			klog.CtxErrorf(s.ctx, "orderId:%s,更新订单状态失败,err:%s", orderId, err.Error())
			return nil, errors.WithStack(err)
		}
		if markOrderPaidResp.StatusCode != 0 {
			klog.CtxErrorf(s.ctx, "orderId:%s,更新订单状态失败,err:%s", orderId, markOrderPaidResp.StatusMsg)
			return &payment.NotifyPaymentResp{
				StatusCode: markOrderPaidResp.StatusCode,
				StatusMsg:  constant.GetMsg(int(markOrderPaidResp.StatusCode)),
			}, nil
		}
		//发送支付成功消息
		producer.SendPaymentSuccessOrderIdMsg(s.ctx, orderId)
	}

	klog.CtxInfof(s.ctx, "订单已支付或已取消，无需更新订单状态")
	resp = &payment.NotifyPaymentResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
	}
	return resp, nil

}
