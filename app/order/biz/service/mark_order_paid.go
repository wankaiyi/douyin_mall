package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/order/biz/dal/mysql"
	"douyin_mall/order/biz/model"
	order "douyin_mall/order/kitex_gen/order"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type MarkOrderPaidService struct {
	ctx context.Context
} // NewMarkOrderPaidService new MarkOrderPaidService
func NewMarkOrderPaidService(ctx context.Context) *MarkOrderPaidService {
	return &MarkOrderPaidService{ctx: ctx}
}

// Run create note info
func (s *MarkOrderPaidService) Run(req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	ctx := s.ctx
	orderId := req.OrderId
	rowAffected, err := model.MarkOrderPaid(ctx, mysql.DB, orderId)
	if err != nil {
		klog.CtxErrorf(ctx, "数据库修改订单状态为支付成功结果失败，orderId:%s,err:%v", orderId, err)
		return nil, errors.New("数据库修改订单状态失败")
	}
	if rowAffected == 0 {
		klog.CtxErrorf(ctx, "修改订单状态为支付成功结果失败，原因是订单状态已被修改或不存在，orderId:%s", orderId)
		return &order.MarkOrderPaidResp{StatusCode: 3000, StatusMsg: constant.GetMsg(3000)}, nil
	}
	klog.CtxInfof(ctx, "修改订单状态为支付成功结果成功，orderId:%s", orderId)
	return &order.MarkOrderPaidResp{StatusCode: 0, StatusMsg: constant.GetMsg(0)}, nil
}
