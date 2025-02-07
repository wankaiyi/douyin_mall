package task

import (
	"archive/zip"
	"context"
	"douyin_mall/payment/biz/dal/mysql"
	"douyin_mall/payment/biz/model"
	"douyin_mall/payment/biz/service"
	"encoding/csv"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
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
	hlog.CtxInfof(ctx, "对账任务开始 CheckAccountTask start")
	// 下载对账文件
	downloadQueryZip(ctx)
	// 解析对账文件
	resolutionCsv(ctx)

	hlog.Infof("对账任务结束 CheckAccountTask end")
	return "task done"

}

func downloadQueryZip(ctx context.Context) {
	// 目标URL
	billDownloadUrl, err := service.QueryPay(ctx)
	if err != nil {
		hlog.CtxErrorf(ctx, "Get billDownloadUrl error: %v", err)
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
		hlog.CtxErrorf(ctx, "client.Get error: %v", err)
	}
	if len(body) < 100 {
		hlog.CtxErrorf(ctx, "Request time out!!!")
		return
	}
	hlog.CtxInfof(ctx, "statusCode: %v, body: %v", statusCode, body)
	hlog.CtxInfof(ctx, "body len: %v", len(body))
	// 确保关闭响应体

	// 检查HTTP状态码
	if statusCode != consts.StatusOK {
		hlog.CtxErrorf(ctx, "Get status code: %v", statusCode)
	}

	now := time.Now().Local()
	yesterday := now.Add(-time.Hour * 24).Format("2006-01-02")
	zipPath := "./resource/accountFile/" + yesterday + ".zip"

	// 写入文件
	err = os.WriteFile(zipPath, body, 0666)
	if err != nil {
		hlog.CtxErrorf(ctx, "os.WriteFile error: %v", err)
	}
	fmt.Println("download success")
}

// resolutionZip 解析对账文件
func resolutionZip(ctx context.Context) (*os.File, io.ReadCloser, error) {
	now := time.Now().Local()
	yesterday := now.Add(-time.Hour * 24).Format("2006-01-02")
	zipPath := "./resource/accountFile/" + yesterday + ".zip"

	zipFile, err := os.Open(zipPath)
	if err != nil {
		hlog.CtxErrorf(ctx, "os.Open error: %v", err)
	}

	zipFileInfo, err := zipFile.Stat()
	if err != nil {
		hlog.CtxErrorf(ctx, "zipFile.Stat error: %v", err)
	}

	zipReader, err := zip.NewReader(zipFile, zipFileInfo.Size())

	if err != nil {
		hlog.CtxErrorf(ctx, "zip.NewReader error: %v", err)
	}

	var csvFile io.ReadCloser
	userId := "20887210564732070156"
	yesterdayDate := now.Add(-time.Hour * 24).Format("20060102")

	for _, file := range zipReader.File {
		fileName, err := decodeGbk(file.Name)
		if err != nil {
			hlog.CtxErrorf(ctx, "decodeGbk error: %v", err)
		}
		if fileName == userId+"_"+yesterdayDate+"_业务明细.csv" {

			csvFile, err = file.Open()

			if err != nil {
				hlog.CtxErrorf(ctx, "csvFile.Open error: %v", err)
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
		hlog.CtxErrorf(ctx, "resolutionZip error: %v", err)
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
			hlog.CtxErrorf(ctx, "csvReader.Read error: %v", err)
			continue
		}

		go findStartIndex(ctx, i, record, startIndexCh)
		go findEndIndex(ctx, i, record, endIndexCh)

	}

	startIndex := <-startIndexCh
	endIndex := <-endIndexCh
	if startIndex+2 == endIndex {
		date := time.Now().Local().Format("2006-01-02")
		hlog.CtxInfof(ctx, "日期:%s 没有数据", date)
		//检查数据库值当日是否有数据

		fmt.Println("All data processed.")

		return
	}

	hlog.CtxInfof(ctx, "开始读取并处理数据")

	//重置读取指针
	resetZipFile, resetCsv, err := resolutionZip(ctx)

	defer resetZipFile.Close()
	defer resetCsv.Close()
	if err != nil {
		hlog.CtxErrorf(ctx, "resolutionZip error: %v", err)
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
			hlog.CtxErrorf(ctx, "csvReader.Read error: %v", err)
			continue
		}
		wg.Add(1)
		sem <- 1 // 获取一个信号量
		go func() {
			defer func() { <-sem }() // 释放信号量
			// 处理数据
			handlerData(ctx, i, record, startIndex, endIndex, sem, &wg)
		}()
	}

	wg.Wait()
	fmt.Println("数据处理完成 All data processed.")
}

func findStartIndex(ctx context.Context, index int, content []string, startIndexCh chan int) {

	data, err := decodeGbk(content[0])
	if err != nil {
		hlog.CtxErrorf(ctx, "decodeGbk error: %v", err)
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
		hlog.CtxErrorf(ctx, "decodeGbk error: %v", err)
		return
	}
	if data == "#-----------------------------------------业务明细列表结束------------------------------------" {
		endIndexCh <- index
	}

	return

}

func handlerData(ctx context.Context, index int, content []string, startIndex int, endIndex int, sem chan int, wg *sync.WaitGroup) {
	if index > startIndex+1 && index < endIndex {
		//for循环处理数据
		//0为支付宝交易号，1为订单号，13为商家实收的金额
		alipayTradeNo := content[0]
		orderId := content[1]
		amount := content[13]
		fAmount, err := strconv.ParseFloat(amount, 64)
		if err != nil {
			hlog.CtxErrorf(ctx, "strconv.ParseFloat error: %v", err)
		}
		//检查数据库中是否存在数据
		paymentOrder, err := model.GetPaymentOrdersByOrderID(mysql.DB, ctx, orderId)
		if err != nil {
			hlog.CtxErrorf(ctx, "model.GetPaymentOrdersByOrderID error: %v", err)
			return
		}

		if paymentOrder == nil {
			hlog.CtxErrorf(ctx, "订单号:%s 支付宝交易号:%s 支付金额:%f 数据库中不存在!!!", orderId, alipayTradeNo, fAmount)
			return
		} else if paymentOrder.Amount != fAmount {
			hlog.CtxWarnf(ctx, "订单号:%s 支付宝交易号:%s 支付金额:%f 数据库金额:%f 不一致!!!", orderId, alipayTradeNo, fAmount, paymentOrder.Amount)
			return
		}
		hlog.CtxInfof(ctx, "订单号:%s 支付宝交易号:%s 支付金额:%f 数据库金额:%f  对账状态正常", orderId, alipayTradeNo, fAmount, paymentOrder.Amount)
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
