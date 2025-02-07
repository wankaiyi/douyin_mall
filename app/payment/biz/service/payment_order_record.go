package service

import (
	"context"
	"douyin_mall/common/constant"
	"douyin_mall/payment/biz/dal/mysql"
	"douyin_mall/payment/biz/model"
	payment "douyin_mall/payment/kitex_gen/payment"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
)

type PaymentOrderRecordService struct {
	ctx context.Context
} // NewPaymentOrderRecordService new PaymentOrderRecordService
func NewPaymentOrderRecordService(ctx context.Context) *PaymentOrderRecordService {
	return &PaymentOrderRecordService{ctx: ctx}
}

// Run create note info
func (s *PaymentOrderRecordService) Run(req *payment.PaymentOrderRecordReq) (resp *payment.PaymentOrderRecordResp, err error) {
	// Finish your business logic.
	var paymentOrders model.PaymentOrder
	paymentOrders.OrderID = req.OrderId
	paymentOrders.UserID = req.UserId
	paymentOrders.Amount = float64(req.Amount)
	paymentOrders.Status = req.Status
	err = model.AddPaymentOrders(mysql.DB, s.ctx, &paymentOrders)
	if err != nil {
		klog.CtxErrorf(s.ctx, "create payment order record failed, err: %v, req: %+v", err, req)
		return nil, errors.WithStack(err)
	}
	resp = &payment.PaymentOrderRecordResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
	}
	return
}
