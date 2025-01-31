package utils

import (
	"bytes"
	"github.com/cloudwego/hertz/pkg/protocol"
	"io"
	"net/http"
)

// ConvertToHTTPRequest 将 Hertz 的 *protocol.Request 转换为标准库的 *http.Request
func ConvertToHTTPRequest(req *protocol.Request) (*http.Request, error) {
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
