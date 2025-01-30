package service

import (
	"bytes"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/xlog"
	"io"
	"net/http"
	"os"
)

var (
	appid      string
	notifyUrl  string
	privateKey string
	//订单标题
	subject string
	//产品码 沙箱环境仅支持value = FAST_INSTANT_TRADE_PAY
	product_code string
	// 支付宝公钥证书
	AlipayPublicContent []byte

	// 支付宝根证书
	AlipayRootContent []byte

	// 应用公钥证书
	AppPublicContent []byte
)

func init() {
	AlipayPublicContentPath, _ := os.Open(os.Getenv("ALIPAY_PUBLIC_CONTENT_PATH"))
	AlipayRootContentPath, _ := os.Open(os.Getenv("ALIPAY_ROOT_CONTENT_PATH"))
	AppPublicContentPath, _ := os.Open(os.Getenv("APP_PUBLIC_CONTENT_PATH"))
	defer AlipayPublicContentPath.Close()
	defer AlipayRootContentPath.Close()
	defer AppPublicContentPath.Close()

	appid = os.Getenv("APPID")
	notifyUrl = os.Getenv("NOTIFY_URL")
	privateKey = os.Getenv("PRIVATE_KEY")
	subject = "douyin_mall_order"
	product_code = "FAST_INSTANT_TRADE_PAY"

	AlipayPublicContent, _ = io.ReadAll(AlipayPublicContentPath)
	AlipayRootContent, _ = io.ReadAll(AlipayRootContentPath)
	AppPublicContent, _ = io.ReadAll(AppPublicContentPath)

}

func PayInit() (client *alipay.Client, err error) {

	// 初始化支付宝客户端
	// appid：应用ID
	// privateKey：应用私钥，支持PKCS1和PKCS8
	// isProd：是否是正式环境，沙箱环境请选择新版沙箱应用。
	client, err = alipay.NewClient(appid, privateKey, false)
	if err != nil {
		xlog.Error(err)
		return
	}

	// 打开Debug开关，输出日志，默认关闭
	client.DebugSwitch = gopay.DebugOn

	// 设置支付宝请求 公共参数
	//    注意：具体设置哪些参数，根据不同的方法而不同，此处列举出所有设置参数
	client.SetLocation(alipay.LocationShanghai). // 设置时区，不设置或出错均为默认服务器时间
							SetCharset(alipay.UTF8).  // 设置字符编码，不设置默认 utf-8
							SetSignType(alipay.RSA2). // 设置签名类型，不设置默认 RSA2
							SetNotifyUrl(notifyUrl)   // 设置异步通知URL
	//SetReturnUrl("https://www.fmm.ink"). // 设置返回URL
	// 设置biz_content加密KEY，设置此参数默认开启加密（目前不可用，设置后会报错）
	//client.SetAESKey("1234567890123456")

	// 自动同步验签（只支持证书模式）
	// 传入 alipayPublicCert.crt 内容
	client.AutoVerifySign(AlipayPublicContent)

	// 公钥证书模式，需要传入证书，以下两种方式二选一
	// 证书路径
	//err = client.SetCertSnByPath()
	// 证书内容
	err = client.SetCertSnByContent(AppPublicContent, AlipayRootContent, AlipayPublicContent)

	if err != nil {
		xlog.Error(err)
	}
	return
}

func Pay(ctx context.Context, orderId int64, totalAmount float32) (result string, err error) {
	// 构建支付请求对象
	client, err := PayInit()
	if err != nil {
		return "", err
	}
	// 构建支付请求参数
	bodyMap := make(gopay.BodyMap)
	bodyMap.Set("out_trade_no", orderId)
	bodyMap.Set("total_amount", totalAmount)
	bodyMap.Set("subject", subject)
	bodyMap.Set("product_code", product_code)
	paymentUrl, err := client.TradePagePay(ctx, bodyMap)
	if err != nil {
		return "", err
	}
	// 跳转到支付页面
	return paymentUrl, nil
}

// 支付宝支付通知异步回调
func PayNotify(ctx *context.Context, notifyBody string, c *app.RequestContext) {
	// 解析异步通知的参数
	// req：*http.Request
	request, err := convertToHTTPRequest(&c.Request)
	if err != nil {
		xlog.Error(err)
		return
	}
	notifyReq, err := alipay.ParseNotifyToBodyMap(request) // c.Request 是 gin 框架的写法
	if err != nil {
		xlog.Error(err)
		return
	}
	//或
	// value：url.Values
	//
	//notifyReq, err = alipay.ParseNotifyByURLValues()
	//if err != nil {
	//	xlog.Error(err)
	//	return
	//}

	//// 支付宝异步通知验签（公钥模式）
	//ok, err = alipay.VerifySign(aliPayPublicKey, notifyReq)

	// 支付宝异步通知验签（公钥证书模式）
	_, err = alipay.VerifySignWithCert(AlipayPublicContent, notifyReq)
	if err != nil {
		xlog.Error(err)
		return
	}

	// 如果需要，可将 BodyMap 内数据，Unmarshal 到指定结构体指针 ptr
	//err = notifyReq.Unmarshal(ptr)

	// ====异步通知，返回支付宝平台的信息====
	// 文档：https://opendocs.alipay.com/open/203/105286
	// 程序执行完后必须打印输出“success”（不包含引号）。如果商户反馈给支付宝的字符不是success这7个字符，支付宝服务器会不断重发通知，直到超过24小时22分钟。一般情况下，25小时以内完成8次通知（通知的间隔频率一般是：4m,10m,10m,1h,2h,6h,15h）

	// 此写法是 gin 框架返回支付宝的写法
	//c.String(http.StatusOK, "%s", "success")

	// 此写法是 echo 框架返回支付宝的写法
	c.String(http.StatusOK, "success")
}

// convertToHTTPRequest 将 Hertz 的 *protocol.Request 转换为标准库的 *http.Request
func convertToHTTPRequest(req *protocol.Request) (*http.Request, error) {
	// 读取请求体
	body := io.NopCloser(bytes.NewReader(req.Body()))

	// 创建 *http.Request
	httpReq, err := http.NewRequest(
		string(req.Method()),
		req.URI().String(),
		body,
	)
	if err != nil {
		return nil, err
	}

	// 复制 Header
	req.Header.VisitAll(func(key, value []byte) {
		httpReq.Header.Set(string(key), string(value))
	})

	return httpReq, nil
}
