package middleware

import (
	"bytes"
	"compress/gzip"
	"context"
	logUtils "douyin_mall/api/mtl/log"
	"douyin_mall/common/constant"
	"douyin_mall/common/utils/env"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/google/uuid"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

type DouyinTextFormatter struct {
}

func (f *DouyinTextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(entry.Message + "\n"), nil
}

type SystemLog struct {
	IP       string `json:"ip"`
	URI      string `json:"uri"`
	Method   string `json:"method"`
	CostTime int64  `json:"cost_time"`
	Status   int    `json:"status"`
	Resp     string `json:"resp"`
	ErrMsg   string `json:"err_msg,omitempty"`
}

func TraceLogMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		startTime := time.Now()
		ip := c.ClientIP()
		uri := string(c.Request.Path())
		method := string(c.Request.Method())
		traceId := uuid.New().String()
		ctx = context.WithValue(ctx, constant.TraceId, traceId)

		c.Next(ctx)

		statusCode := c.Response.StatusCode()
		resp := c.Response.Body()
		if isGzipEncoded(c) {
			if decoded, err := decompressGzip(resp); err == nil {
				resp = decoded
			} else {
				resp = []byte("Failed to decompress")
			}
		}

		costTime := time.Since(startTime).Milliseconds()
		log := logrus.New()
		log.SetReportCaller(true)
		log.SetFormatter(&DouyinTextFormatter{})
		logger := hertzlogrus.NewLogger(hertzlogrus.WithLogger(log))
		if currentEnv, err := env.GetString("env"); err == nil && currentEnv == "dev" {
			logger.SetOutput(os.Stdout)
		} else {
			logger.SetOutput(logUtils.AsyncWriter)
		}
		marshal, _ := sonic.Marshal(SystemLog{
			IP:       ip,
			URI:      uri,
			Method:   method,
			CostTime: costTime,
			Status:   statusCode,
			Resp:     string(resp),
		})
		logger.Info(string(marshal))
	}
}

func isGzipEncoded(c *app.RequestContext) bool {
	return bytes.Contains(c.Response.Header.Peek("Content-Encoding"), []byte("gzip"))
}

// 解压缩 Gzip 数据
func decompressGzip(data []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	var out bytes.Buffer
	_, err = io.Copy(&out, reader)
	return out.Bytes(), err
}
