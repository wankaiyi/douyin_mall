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

			// 用户服务 & 鉴权服务
			1000: "二次确认密码不一致",
			1001: "用户名或密码错误",
			1002: "用户已存在，请登录",
			1003: "用户名或密码错误",
			1004: "无权限操作",
			1005: "已经在其他地方登录",
			1006: "登录过期，请重新登录",
			1007: "性别不存在",
			1008: "用户名已存在",
			1009: "资源不存在", // 没权限
			1010: "您的账号已被冻结",
			1011: "默认收货地址已存在",

			// 豆包AI服务
			2000: "输入不能为空",

			//支付服务
			5000: "支付异常",
			5001: "支付参数错误",
			5002: "支付失败",
			5004: "支付系统错误，取消失败",
			5005: "支付系统错误，订单记录失败",
			5006: "支付宝异步通知失败",

			//商品服务
			6000: "商品新增失败",
			6001: "商品更新失败",
			6002: "商品删除失败",
			6003: "商品查询失败",
			6004: "商品不存在",
			6005: "商品已下架",
			6006: "商品库存不足",
			6007: "商品已售完",
			6008: "商品已下单",
			6009: "商品已售空",
			6010: "商品已锁定",
			6011: "商品已过期",
			6012: "商品库存不足",
			6013: "商品搜索失败",
			6014: "商品库存锁定失败",
			//分类
			7014: "商品分类新增失败",
			7015: "商品分类删除失败",
			7016: "商品分类更新失败",
			7017: "商品分类查询失败",
			//品牌
			8018: "商品品牌新增失败",
			8019: "商品品牌删除失败",
			8020: "商品品牌更新失败",
			8021: "商品品牌查询失败",
		}
	})
}

func GetMsg(errorID int) string {
	if msg, exists := commonMsgMap[errorID]; exists {
		return msg
	}
	return commonMsgMap[DefaultErrorId]
}
