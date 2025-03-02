package feishu

import (
	"context"
	"douyin_mall/common/utils/http"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/klog"
	"strings"
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
	_, err := http.Post(ctx, url, map[string]string{
		"Content-Type": "application/json",
	}, body, true)
	if err != nil {
		if strings.HasPrefix(text, "服务") {
			klog.Error("飞书告警异常", "url", url, "请求体", string(body), "错误", err)
		} else {
			hlog.Error("飞书告警异常", "url", url, "请求体", string(body), "错误", err)
		}
		return
	}
}
