package service

import (
	"context"
	"douyin_mall/api/biz/utils"
	"douyin_mall/api/infra/rpc"
	payment "douyin_mall/rpc/kitex_gen/payment"
	"errors"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-pay/gopay/alipay"
	"io"
	"net/http"
	"os"
	"strconv"
)

var orderRecordCh = make(chan int)
var transactionRecordCh = make(chan int)

type NotifyService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewNotifyService(Context context.Context, RequestContext *app.RequestContext) *NotifyService {
	return &NotifyService{RequestContext: RequestContext, Context: Context}
}

func (h *NotifyService) Run(RequestContext *app.RequestContext) (err error) {
	//defer func() {
	//hlog.CtxInfof(h.Context, "req = %+v", req)
	//hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	AlipayPublicContentPath, _ := os.Open(os.Getenv("ALIPAY_PUBLIC_CONTENT_PATH"))
	AlipayPublicContent, _ := io.ReadAll(AlipayPublicContentPath)

	// 解析异步通知的参数
	// req：*http.Request
	request, err := utils.ConvertToHTTPRequest(&RequestContext.Request)
	if err != nil {
		hlog.CtxErrorf(h.Context, "ConvertToHTTPRequest error: %v", err)
		return
	}
	notifyReq, err := alipay.ParseNotifyToBodyMap(request) // c.Request 是 gin 框架的写法
	if err != nil {
		hlog.CtxErrorf(h.Context, "ParseNotifyToBodyMap error: %v", err)
		return
	}

	alipayTradeNo := notifyReq.GetString("trade_no")
	orderId := notifyReq.GetString("out_trade_no")
	idempotentControlReq := &payment.IdempotentControlReq{
		OrderId: orderId,
	} //幂等控制
	_, err = rpc.PaymentClient.IdempotentControl(h.Context, idempotentControlReq)
	if err != nil {
		hlog.CtxErrorf(h.Context, "订单号重复！err: %v", err)
		return errors.New("订单号重复")
	}

	body := notifyReq.GetString("body")
	tradeStatus := notifyReq.GetString("trade_status")
	amount, _ := strconv.ParseFloat(notifyReq.GetString("total_amount"), 32)
	//todo userID应该通过订单查找
	userId := h.Context.Value("user_id").(int32)

	if tradeStatus != "TRADE_SUCCESS" {
		hlog.CtxErrorf(h.Context, "trade_status is not TRADE_SUCCESS")
		//记录订单信息和记录交易流水信息
		go RecordOrderInfo(h.Context, orderId, userId, float32(amount), 2)
		go RecordTransactionInfo(h.Context, orderId, alipayTradeNo, tradeStatus, notifyReq, body)

	} else {
		hlog.CtxInfof(h.Context, "trade_status is TRADE_SUCCESS")
		//记录订单信息和记录交易流水信息
		go RecordOrderInfo(h.Context, orderId, userId, float32(amount), 1)
		go RecordTransactionInfo(h.Context, orderId, alipayTradeNo, tradeStatus, notifyReq, body)
		//todo 支付成功逻辑

	}

	// 支付宝异步通知验签（公钥证书模式）
	_, err = alipay.VerifySignWithCert(AlipayPublicContent, notifyReq)
	if err != nil {
		hlog.CtxErrorf(h.Context, err.Error())
	}

	// ====异步通知，返回支付宝平台的信息====
	// 程序执行完后必须打印输出“success”（不包含引号）。如果商户反馈给支付宝的字符不是success这7个字符，支付宝服务器会不断重发通知，直到超过24小时22分钟。一般情况下，25小时以内完成8次通知（通知的间隔频率一般是：4m,10m,10m,1h,2h,6h,15h）

	// 此写法是 echo 框架返回支付宝的写法
	RequestContext.String(http.StatusOK, "success")
	<-orderRecordCh
	<-transactionRecordCh
	return
}
func RecordOrderInfo(ctx context.Context, orderId string, userId int32, amount float32, status int32) {
	var paymentOrderRecordReq *payment.PaymentOrderRecordReq
	paymentOrderRecordReq.OrderId = orderId
	paymentOrderRecordReq.Status = status
	paymentOrderRecordReq.UserId = userId
	paymentOrderRecordReq.Amount = amount
	_, err := rpc.PaymentClient.PaymentOrderRecord(ctx, paymentOrderRecordReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "记录订单信息失败！err: %v，req: %+v", err, paymentOrderRecordReq)
	}
	orderRecordCh <- 1

}
func RecordTransactionInfo(ctx context.Context, orderId string, alipayTradeNo string, tradeStatus string, requestParams interface{}, callback string) {
	var paymentTransactionRecordReq *payment.PaymentTransactionRecordReq
	paymentTransactionRecordReq.OrderId = orderId
	paymentTransactionRecordReq.AlipayTradeNo = alipayTradeNo
	paymentTransactionRecordReq.TradeStatus = tradeStatus
	paymentTransactionRecordReq.Callback = callback
	reqParams, err := sonic.Marshal(requestParams)
	if err != nil {
		hlog.CtxErrorf(ctx, "序列化requestParams失败！err: %v,reqParams:%+v", err, requestParams)
		transactionRecordCh <- 1
		return
	}
	paymentTransactionRecordReq.RequestParams = string(reqParams)
	_, err = rpc.PaymentClient.PaymentTransactionRecord(ctx, paymentTransactionRecordReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "记录交易流水信息失败！err: %v,req: %+v", err, paymentTransactionRecordReq)

	}
	transactionRecordCh <- 1

}
