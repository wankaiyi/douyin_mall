package constant

import (
	"sync"
)

const DefaultErrorId = 500

var (
	once         sync.Once
	commonMsgMap map[int]string
)

func init() {
	once.Do(func() {
		commonMsgMap = map[int]string{
			0:   "成功",
			500: "服务器异常",

			// 用户服务
			1000: "二次确认密码不一致",
			1001: "用户名或密码错误",
			1002: "用户已存在，请登录",
			1003: "邮箱或密码错误",
			1004: "无权限操作",
			1005: "已经在其他地方登录",
			1006: "登录过期，请重新登录",
			1007: "性别不存在",
			1008: "用户名已存在",

			// 豆包AI服务
			2000: "输入不能为空",

			//支付服务
			5000: "支付异常",
			5001: "支付参数错误",
			5002: "支付失败",
			5004: "支付系统错误，取消失败",
		}
	})
}

func GetMsg(errorID int) string {
	if msg, exists := commonMsgMap[errorID]; exists {
		return msg
	}
	return commonMsgMap[DefaultErrorId]
}
