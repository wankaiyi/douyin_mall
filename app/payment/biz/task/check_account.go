package task

import (
	"archive/zip"
	"context"
	commonConstant "douyin_mall/common/constant"
	"douyin_mall/payment/biz/dal/alipay"
	"douyin_mall/payment/biz/dal/mysql"
	"douyin_mall/payment/biz/model"
	"douyin_mall/payment/infra/rpc"
	"douyin_mall/rpc/kitex_gen/order"
	"encoding/csv"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/klog"

	"strconv"

	"github.com/xxl-job/xxl-job-executor-go"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"os"
	"sync"
	"time"
)

func CheckAccountTask(ctx context.Context, param *xxl.RunReq) (msg string) {
	klog.CtxInfof(ctx, "对账任务开始 CheckAccountTask start")
	// 下载对账文件
	downloadQueryZip(ctx)
	// 解析对账文件
	resolutionCsv(ctx)

	klog.CtxInfof(ctx, "对账任务结束 CheckAccountTask end")
	return "task done"

}

func downloadQueryZip(ctx context.Context) {
	// 目标URL
	billDownloadUrl, err := alipay.QueryBill(ctx)
	if err != nil {
		klog.CtxErrorf(ctx, "Get billDownloadUrl error: %v", err)
		return
	}

	//下载对账zip文件
	// 创建Hertz客户端
	cli, err := client.NewClient()
	if err != nil {
		panic(err)
	}

	// 发送GET请求
	statusCode, body, err := cli.Get(ctx, nil, billDownloadUrl)
	if err != nil {
		klog.CtxErrorf(ctx, "client.Get error: %v", err)
	}
	if len(body) < 100 {
		klog.CtxErrorf(ctx, "Request time out!!!")
		return
	}
	klog.CtxInfof(ctx, "statusCode: %v, body: %v", statusCode, body)
	klog.CtxInfof(ctx, "body len: %v", len(body))
	// 确保关闭响应体

	// 检查HTTP状态码
	if statusCode != consts.StatusOK {
		klog.CtxErrorf(ctx, "Get status code: %v", statusCode)
	}

	now := time.Now().Local()
	yesterday := now.Add(-time.Hour * 24).Format("2006-01-02")
	zipPath := "./resource/" + yesterday + ".zip"

	// 写入文件
	err = os.WriteFile(zipPath, body, 0666)
	if err != nil {
		klog.CtxErrorf(ctx, "os.WriteFile error: %v", err)
	}
	fmt.Println("download success")
}

// resolutionZip 解析对账文件
func resolutionZip(ctx context.Context) (*os.File, io.ReadCloser, error) {
	now := time.Now().Local()
	yesterday := now.Add(-time.Hour * 24).Format("2006-01-02")
	zipPath := "./resource/" + yesterday + ".zip"

	zipFile, err := os.Open(zipPath)
	if err != nil {
		klog.CtxErrorf(ctx, "os.Open error: %v", err)
	}

	zipFileInfo, err := zipFile.Stat()
	if err != nil {
		klog.CtxErrorf(ctx, "zipFile.Stat error: %v", err)
	}

	zipReader, err := zip.NewReader(zipFile, zipFileInfo.Size())

	if err != nil {
		klog.CtxErrorf(ctx, "zip.NewReader error: %v", err)
	}

	var csvFile io.ReadCloser
	userId := "20887210564732070156"
	yesterdayDate := now.Add(-time.Hour * 24).Format("20060102")

	for _, file := range zipReader.File {
		fileName, err := decodeGbk(file.Name)
		if err != nil {
			klog.CtxErrorf(ctx, "decodeGbk error: %v", err)
		}
		if fileName == userId+"_"+yesterdayDate+"_业务明细.csv" {

			csvFile, err = file.Open()

			if err != nil {
				klog.CtxErrorf(ctx, "csvFile.Open error: %v", err)
			}
			break
		}
	}
	return zipFile, csvFile, err
}

// resolutionCsv 获取csv文件并解析
func resolutionCsv(ctx context.Context) {
	zipFile, csvFile, err := resolutionZip(ctx)
	defer zipFile.Close()
	defer csvFile.Close()
	if err != nil {
		klog.CtxErrorf(ctx, "resolutionZip error: %v", err)
	}
	csvReader := csv.NewReader(csvFile)
	startIndexCh := make(chan int)
	endIndexCh := make(chan int)

	for i := 0; ; i++ {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			klog.CtxErrorf(ctx, "csvReader.Read error: %v", err)
			continue
		}

		go findStartIndex(ctx, i, record, startIndexCh)
		go findEndIndex(ctx, i, record, endIndexCh)

	}

	startIndex := <-startIndexCh
	endIndex := <-endIndexCh
	if startIndex+2 == endIndex {
		date := time.Now().Local().Format("2006-01-02")
		klog.CtxInfof(ctx, "日期:%s 没有数据", date)

		klog.CtxInfof(ctx, "All data processed.")

		return
	}

	klog.CtxInfof(ctx, "开始读取并处理数据")

	//重置读取指针
	resetZipFile, resetCsv, err := resolutionZip(ctx)

	defer resetZipFile.Close()
	defer resetCsv.Close()
	if err != nil {
		klog.CtxErrorf(ctx, "resolutionZip error: %v", err)
	}
	resetReader := csv.NewReader(resetCsv)
	var wg sync.WaitGroup

	concurrencyLimit := 10 // 限制并发数量

	sem := make(chan int, concurrencyLimit) // 创建带缓冲的 channel

	for i := 0; ; i++ {
		record, err := resetReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			klog.CtxErrorf(ctx, "csvReader.Read error: %v", err)
			continue
		}
		wg.Add(1)
		sem <- 1 // 获取一个信号量
		i := i
		go func() {
			defer func() { <-sem }() // 释放信号量
			// 处理数据
			handlerData(ctx, i, record, startIndex, endIndex, &wg)
		}()
	}

	wg.Wait()
	fmt.Println("数据处理完成 All data processed.")
}

func findStartIndex(ctx context.Context, index int, content []string, startIndexCh chan int) {

	data, err := decodeGbk(content[0])
	if err != nil {
		klog.CtxErrorf(ctx, "decodeGbk error: %v", err)
		return
	}
	if data == "#-----------------------------------------业务明细列表----------------------------------------" {
		startIndexCh <- index
	}

	return

}
func findEndIndex(ctx context.Context, index int, content []string, endIndexCh chan int) {

	data, err := decodeGbk(content[0])
	if err != nil {
		klog.CtxErrorf(ctx, "decodeGbk error: %v", err)
		return
	}
	if data == "#-----------------------------------------业务明细列表结束------------------------------------" {
		endIndexCh <- index
	}

	return

}

func handlerData(ctx context.Context, index int, content []string, startIndex int, endIndex int, wg *sync.WaitGroup) {
	if index > startIndex+1 && index < endIndex {
		//for循环处理数据
		//0为支付宝交易号，1为订单号，13为商家实收的金额
		alipayTradeNo := content[0]
		orderId := content[1]
		AlipayAmount := content[13]
		AlipayAmountFloat, err := strconv.ParseFloat(AlipayAmount, 64)
		if err != nil {
			klog.CtxErrorf(ctx, "strconv.ParseFloat error: %v", err)
		}
		//检查数据库中是否存在数据
		orderResp, err := rpc.OrderClient.GetOrder(ctx, &order.GetOrderReq{
			OrderId: orderId,
		})
		if err != nil {
			klog.CtxErrorf(ctx, "获取订单数据失败,err： %v，orderId：%s", err, orderId)
			return
		}

		if orderResp == nil {
			klog.CtxErrorf(ctx, "订单号:%s 支付宝交易号:%s 支付金额:%f 数据库中不存在!!!", orderId, alipayTradeNo, AlipayAmountFloat)
			return
			//订单金额与支付金额不一致
		} else if orderResp.Order.Cost != AlipayAmountFloat {
			klog.CtxWarnf(ctx, "订单号:%s 支付宝交易号:%s 支付金额:%f 数据库金额:%f 不一致!!!", orderId, alipayTradeNo, AlipayAmountFloat, orderResp.Order.Cost)

			err := model.CreateCheckRecord(mysql.DB, ctx, &model.CheckRecord{
				ReconDate:     time.Now().Local(),
				OrderId:       orderId,
				AlipayTradeNo: alipayTradeNo,
				AlipayAmount:  AlipayAmountFloat,
				LocalAmount:   orderResp.Order.Cost,
				Status:        commonConstant.InconsistentCheckRecordStatus,
			})
			if err != nil {
				klog.CtxErrorf(ctx, "创建对账记录失败,err: %v ,orderId: %s", err, orderId)
			}
			return
			//订单金额与支付金额一致
		}
		err = model.CreateCheckRecord(mysql.DB, ctx, &model.CheckRecord{
			ReconDate:     time.Now().Local(),
			OrderId:       orderId,
			AlipayTradeNo: alipayTradeNo,
			AlipayAmount:  AlipayAmountFloat,
			LocalAmount:   orderResp.Order.Cost,
			Status:        commonConstant.ConsistentCheckRecordStatus,
		})
		if err != nil {
			klog.CtxErrorf(ctx, "创建对账记录失败,err: %v ,orderId: %s", err, orderId)
			return
		}

		klog.CtxInfof(ctx, "订单号:%s 支付宝交易号:%s 支付金额:%f 数据库金额:%f  对账状态正常", orderId, alipayTradeNo, AlipayAmountFloat, orderResp.Order.Cost)
	}
	defer wg.Done()

}

// decodeGbk 将GBK编码的字节转换为UTF-8
func decodeGbk(input string) (string, error) {
	decoder := simplifiedchinese.GBK.NewDecoder()
	return transformString(decoder, input)
}

// transformString 使用指定的编码转换器来转换字符串
func transformString(decoder transform.Transformer, input string) (string, error) {

	transformed, _, err := transform.Bytes(decoder, []byte(input))
	if err != nil {
		return "", err
	}
	return string(transformed), nil
}
