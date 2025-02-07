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

type IdempotentControlService struct {
	ctx context.Context
} // NewIdempotentControlService new IdempotentControlService
func NewIdempotentControlService(ctx context.Context) *IdempotentControlService {
	return &IdempotentControlService{ctx: ctx}
}

// Run create note info
func (s *IdempotentControlService) Run(req *payment.IdempotentControlReq) (resp *payment.IdempotentControlResp, err error) {
	// Finish your business logic.
	var idempotentControl model.IdempotentControl
	idempotentControl.OrderID = req.OrderId
	err = model.AddIdempotentControl(mysql.DB, s.ctx, idempotentControl.OrderID)
	if err != nil {
		klog.CtxErrorf(s.ctx, "add idempotent control failed, err: %v ,req: %+v", err, req)
		return nil, errors.WithStack(err)
	}
	resp = &payment.IdempotentControlResp{
		StatusCode: 0,
		StatusMsg:  constant.GetMsg(0),
	}
	return
}
