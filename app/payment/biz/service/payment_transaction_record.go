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

type PaymentTransactionRecordService struct {
	ctx context.Context
} // NewPaymentTransactionRecordService new PaymentTransactionRecordService
func NewPaymentTransactionRecordService(ctx context.Context) *PaymentTransactionRecordService {
	return &PaymentTransactionRecordService{ctx: ctx}
}

// Run create note info
func (s *PaymentTransactionRecordService) Run(req *payment.PaymentTransactionRecordReq) (resp *payment.PaymentTransactionRecordResp, err error) {
	// Finish your business logic.
	var paymentTransaction model.PaymentTransaction
	paymentTransaction.OrderID = req.OrderId
	paymentTransaction.AlipayTradeNo = req.AlipayTradeNo
	paymentTransaction.TradeStatus = req.TradeStatus
	paymentTransaction.Callback = req.Callback
	paymentTransaction.RequestParams = req.RequestParams
	err = model.AddPaymentTransaction(mysql.DB, &paymentTransaction)
	if err != nil {
		klog.CtxErrorf(s.ctx, "add payment transaction record failed, err: %v, req: %+v", err, req)
		return nil, errors.WithStack(err)
	}
	resp = &payment.PaymentTransactionRecordResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
	}

	return
}
