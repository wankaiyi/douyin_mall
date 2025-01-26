package feishu

import (
	"context"
	"douyin_mall/common/utils/http"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func SendFeishuAlert(ctx context.Context, url string, text string) {
	body := []byte(fmt.Sprintf(`{
		"msg_type": "interactive",
		"card": {
			"header": {
				"template": "blue"
			},
			"elements": [
				{
					"tag": "markdown",
					"content": "%s"
				}
			]
		}
	}`, text))
	result, err := http.Post(ctx, url, map[string]string{
		"Content-Type": "application/json",
	}, body)
	if err != nil {
		klog.Error("飞书告警异常", "url", url, "请求体", string(body), "错误", err)
	}
	klog.Info("飞书告警异常", "url", url, "请求体", string(body), "响应", result.StatusCode)
}
