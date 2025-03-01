package alipay

import (
	"context"
	"douyin_mall/payment/conf"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/pkg/errors"
	"io"
	"os"
	"time"
)

var (
	appid      string
	notifyUrl  string
	privateKey string

	// AlipayPublicContent 支付宝公钥证书
	AlipayPublicContent []byte

	// AlipayRootContent 支付宝根证书
	AlipayRootContent []byte

	// AppPublicContent 应用公钥证书
	AppPublicContent []byte

	Client *alipay.Client

	//订单标题
	subject = "抖音商城-支付"
	//产品码 沙箱环境仅支持value = FAST_INSTANT_TRADE_PAY
	product_code = "FAST_INSTANT_TRADE_PAY"
)

func Init() {
	cwd, _ := os.Getwd()
	fmt.Println("当前工作目录:", cwd)
	AlipayPublicContentPath, _ := os.Open("resource/alipay_cert/alipayPublicCert.crt")
	AlipayRootContentPath, _ := os.Open("resource/alipay_cert/alipayRootCert.crt")
	AppPublicContentPath, _ := os.Open("resource/alipay_cert/appPublicCert.crt")
	defer AlipayPublicContentPath.Close()
	defer AlipayRootContentPath.Close()
	defer AppPublicContentPath.Close()

	appid = conf.GetConf().AliPay.AppId
	notifyUrl = conf.GetConf().AliPay.NotifyUrl
	privateKey = conf.GetConf().AliPay.PrivateKey

	AlipayPublicContent, _ = io.ReadAll(AlipayPublicContentPath)
	AlipayRootContent, _ = io.ReadAll(AlipayRootContentPath)
	AppPublicContent, _ = io.ReadAll(AppPublicContentPath)

	// 初始化支付宝客户端
	// appid：应用ID
	// privateKey：应用私钥，支持PKCS1和PKCS8
	// isProd：是否是正式环境，沙箱环境请选择新版沙箱应用。
	var err error
	Client, err = alipay.NewClient(appid, privateKey, false)
	if err != nil {
		klog.CtxErrorf(context.Background(), "初始化支付宝客户端失败, 错误信息: %s", err)
		return
	}

	// 打开Debug开关，输出日志，默认关闭
	Client.DebugSwitch = gopay.DebugOn

	// 设置支付宝请求 公共参数
	//    注意：具体设置哪些参数，根据不同的方法而不同，此处列举出所有设置参数
	Client.SetLocation(alipay.LocationShanghai). // 设置时区，不设置或出错均为默认服务器时间
							SetCharset(alipay.UTF8).  // 设置字符编码，不设置默认 utf-8
							SetSignType(alipay.RSA2). // 设置签名类型，不设置默认 RSA2
							SetNotifyUrl(notifyUrl).  // 设置异步通知URL
							SetCharset(alipay.UTF8)

	// 自动同步验签（只支持证书模式）
	// 传入 alipay_public_cert.crt 内容
	klog.CtxInfof(context.Background(), "支付宝自动同步验签,%+v", AlipayPublicContent)
	Client.AutoVerifySign(AlipayPublicContent)

	// 公钥证书模式，需要传入证书，以下两种方式二选一
	// 证书路径
	//err = client.SetCertSnByPath()
	// 证书内容
	err = Client.SetCertSnByContent(AppPublicContent, AlipayRootContent, AlipayPublicContent)

	if err != nil {
		klog.CtxErrorf(context.Background(), "设置支付宝证书失败, 错误信息: %s", err)
	}

}

func Pay(ctx context.Context, orderId int64, totalAmount float32, userId int32) (result string, err error) {

	// 构建支付请求参数
	bodyMap := make(gopay.BodyMap)
	bodyMap.Set("out_trade_no", orderId)
	bodyMap.Set("total_amount", totalAmount)
	bodyMap.Set("subject", subject)
	bodyMap.Set("product_code", product_code)
	bodyMap.Set("query_options", userId)

	//定时关闭订单
	expireTime := time.Now().Local().Add(time.Minute * 10).Format("2006-01-02 15:04:05")
	bodyMap.Set("time_expire", expireTime)
	paymentUrl, err := Client.TradePagePay(ctx, bodyMap)
	if err != nil {
		return "", err
	}
	// 跳转到支付页面
	return paymentUrl, nil
}

//取消支付支付宝返回格式
//{
//	"alipay_trade_close_response": {
//	"code": "10000",
//	"msg": "Success",
//	"out_trade_no": "1634156123238",
//	"trade_no": "2025013122001473210504640665"
//	},
//"sign": ""
//}

func CancelPay(ctx context.Context, orderId int64) (result bool, err error) {

	bodyMap := make(gopay.BodyMap)
	bodyMap.Set("out_trade_no", orderId)
	aliRsp, err := Client.TradeClose(ctx, bodyMap)
	if err != nil {
		return false, err
	}

	if aliRsp.Response.Code != "10000" || aliRsp.Response.Msg != "Success" {
		return false, errors.New(aliRsp.Response.Msg)
	}
	return true, nil

}

// QueryBill 昨日支付宝对账账单下载
func QueryBill(ctx context.Context) (billDownloadUrl string, err error) {

	now := time.Now().Local()
	yesterday := now.Add(-time.Hour * 24).Format("2006-01-02")

	bodyMap := make(gopay.BodyMap)
	bodyMap.Set("bill_type", "trade")
	bodyMap.Set("bill_date", yesterday)
	rsp, err := Client.DataBillDownloadUrlQuery(ctx, bodyMap)
	if err != nil {
		return "", err
	}
	if rsp.Response.Code != "10000" {
		return "", errors.New(rsp.Response.Msg)
	}
	return rsp.Response.BillDownloadUrl, nil

}

// QueryOrder 查询支付宝订单支付状态
func QueryOrder(ctx context.Context, orderId int64) (tradeStatus string, err error) {
	bodyMap := make(gopay.BodyMap)
	bodyMap.Set("out_trade_no", orderId)
	aliRsp, err := Client.TradeQuery(ctx, bodyMap)
	if err != nil {
		return "", errors.WithStack(err)
	}
	if aliRsp.Response.Code != "10000" {
		klog.CtxErrorf(ctx, "支付宝查询订单失败, 订单号: %d, 错误信息: %s", orderId, aliRsp.Response.Msg)
	}
	return aliRsp.Response.Msg, nil

}
