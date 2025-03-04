package task

import (
	"context"
	"douyin_mall/order/biz/dal/mysql"
	"douyin_mall/order/biz/model"
	"douyin_mall/order/infra/kafka/producer"
	"douyin_mall/order/infra/rpc"
	"douyin_mall/rpc/kitex_gen/payment"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
	"github.com/xxl-job/xxl-job-executor-go"
	"gorm.io/gorm"
	"sync"
	"time"
)

var (
	maxConcurrency = 10
)

// CancelOrderTask 定时取消超时的订单，作为mq延时取消订单的兜底
func CancelOrderTask(ctx context.Context, param *xxl.RunReq) (msg string) {
	// 查询已超时的订单
	now := time.Now().In(time.UTC)
	orderIdList, err := model.GetOverdueOrder(ctx, mysql.DB, now.Add(-10*time.Minute))
	if err != nil {
		klog.Errorf("定时任务查询超时订单失败: %v", err)
		return "查询超时订单失败" + err.Error()
	}
	klog.Infof("定时任务查询到超时订单：%v", orderIdList)

	if orderIdList == nil || len(orderIdList) == 0 {
		return "success"
	}

	// 取消超时订单的支付，只会取消未支付的订单
	concurrentCancelCharges(ctx, &sync.WaitGroup{}, make(chan struct{}, maxConcurrency), orderIdList)

	var affectedRows int64
	orderIdCount := len(orderIdList)
	batchSize := 10
	for i := 0; i < orderIdCount; i += batchSize {
		end := i + batchSize
		if end > len(orderIdList) {
			end = len(orderIdList)
		}
		batchOrderIds := orderIdList[i:end]
		err = mysql.DB.Transaction(func(tx *gorm.DB) error {
			affectedRows, err = model.CancelOrderList(ctx, tx, batchOrderIds)
			if err != nil {
				return errors.New("数据库修改订单状态失败")
			}
			canceledOrderIdList, err := model.SelectCanceledSuccessOrders(ctx, mysql.DB, batchOrderIds)
			if err != nil {
				return errors.New("数据库查询已取消订单失败")
			}

			err = producer.SendCancelOrderSuccessMessages(ctx, canceledOrderIdList)
			if err != nil {
				return errors.New("发送取消订单成功消息失败")
			}
			return nil
		})
		if err != nil {
			klog.Errorf("定时任务取消超时订单的支付失败: %v", err)
			return "取消超时订单失败"
		}
	}
	klog.Infof("定时任务取消超时订单成功，耗时%.1f秒，本次执行的超时订单：%v，成功取消%d个订单", time.Since(now).Seconds(), orderIdList, affectedRows)
	return "success"
}

func concurrentCancelCharges(ctx context.Context, wg *sync.WaitGroup, guard chan struct{}, orderIdList []string) {
	for _, orderId := range orderIdList {
		wg.Add(1)
		guard <- struct{}{}

		go func(orderId string) {
			defer wg.Done()
			defer func() { <-guard }()
			if err := cancelCharge(ctx, orderId); err != nil {
				klog.Errorf("定时任务取消超时订单的支付失败，订单ID：%s, err： %v", orderId, err)
			}
		}(orderId)
	}
}

func cancelCharge(ctx context.Context, orderId string) error {
	cancelChargeResp, err := rpc.PaymentClient.CancelCharge(ctx, &payment.CancelChargeReq{
		OrderId: orderId,
	})
	if err != nil {
		return err
	}
	klog.Info("取消支付结果，resp:%v", cancelChargeResp)
	if cancelChargeResp.StatusCode != 0 {
		return errors.New("取消支付失败")
	}
	return nil
}
