package service

import (
	"context"
	"douyin_mall/payment/biz/util"
	"errors"
	"github.com/go-pay/gopay"
	"time"
)

var (
	//订单标题
	subject = "抖音商城-支付"
	//产品码 沙箱环境仅支持value = FAST_INSTANT_TRADE_PAY
	product_code = "FAST_INSTANT_TRADE_PAY"
)

func Pay(ctx context.Context, orderId int64, totalAmount float32) (result string, err error) {
	// 构建支付请求对象
	client, err := util.PayInit()
	if err != nil {
		return "", err
	}
	// 构建支付请求参数
	bodyMap := make(gopay.BodyMap)
	bodyMap.Set("out_trade_no", orderId)
	bodyMap.Set("total_amount", totalAmount)
	bodyMap.Set("subject", subject)
	bodyMap.Set("product_code", product_code)
	//定时关闭订单
	expireTime := time.Now().Local().Add(time.Minute * 5).Format("2006-01-02 15:04:05")
	bodyMap.Set("time_expire", expireTime)
	paymentUrl, err := client.TradePagePay(ctx, bodyMap)
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

	client, err := util.PayInit()
	if err != nil {
		return false, err
	}
	bodyMap := make(gopay.BodyMap)
	bodyMap.Set("out_trade_no", orderId)
	aliRsp, err := client.TradeClose(ctx, bodyMap)
	if err != nil {
		return false, err
	}

	if aliRsp.Response.Code != "10000" || aliRsp.Response.Msg != "Success" {
		return false, errors.New(aliRsp.Response.Msg)
	}
	return true, nil

}

// QueryPay 支付宝对账账单下载
func QueryPay(ctx context.Context, orderId int64) (result string, err error) {
	client, err := util.PayInit()
	if err != nil {
		return "", err

	}
	now := time.Now().Local().Format("2022-01-01")

	bodyMap := make(gopay.BodyMap)
	bodyMap.Set("bill_type", "trade")
	bodyMap.Set("bill_date", now)
	rsp, err := client.DataBillDownloadUrlQuery(ctx, bodyMap)
	if err != nil {
		return "", err
	}
	if rsp.Response.Code != "10000" {
		return "", errors.New(rsp.Response.Msg)
	}
	return rsp.Response.BillDownloadUrl, nil

}
