package service

import (
	"context"
	"douyin_mall/api/biz/utils"
	"douyin_mall/api/conf"
	"douyin_mall/api/infra/rpc"
	rpcpayment "douyin_mall/rpc/kitex_gen/payment"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-pay/gopay/alipay"
	"io"
	"net/http"
	"os"
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
	AlipayPublicContentPath, _ := os.Open(conf.GetConf().AliPay.AlipayPublicContentPath)
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
	// 支付宝异步通知验签（公钥证书模式）
	_, err = alipay.VerifySignWithCert(AlipayPublicContent, notifyReq)
	if err != nil {
		hlog.CtxErrorf(h.Context, err.Error())
	}

	var notifyData = make(map[string]string)
	//notifyData["alipayTradeNo"] = notifyReq.GetString("trade_no")
	notifyData["tradeStatus"] = notifyReq.GetString("trade_status")
	//notifyData["alipayAmount"] = notifyReq.GetString("total_amount")
	notifyData["orderId"] = notifyReq.GetString("out_trade_no")

	notifyPaymentResp, err := rpc.PaymentClient.NotifyPayment(h.Context, &rpcpayment.NotifyPaymentReq{
		NotifyData: notifyData,
	})
	if err != nil {
		hlog.CtxErrorf(h.Context, err.Error())
	}
	hlog.CtxInfof(h.Context, "notifyPaymentResp = %v", notifyPaymentResp)

	// ====异步通知，返回支付宝平台的信息====
	// 程序执行完后必须打印输出“success”（不包含引号）。如果商户反馈给支付宝的字符不是success这7个字符，支付宝服务器会不断重发通知，直到超过24小时22分钟。一般情况下，25小时以内完成8次通知（通知的间隔频率一般是：4m,10m,10m,1h,2h,6h,15h）
	RequestContext.String(http.StatusOK, "success")

	return
}
