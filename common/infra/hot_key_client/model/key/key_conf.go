package key

type KeyConf struct {
	// 微服务名称，用于区分不同微服务
	ServiceName string `json:"service_name"`
	//要判断为热key的key的阈值
	Threshold int64 `json:"threshold"`
	//判断为热key的key的判断时间,单位秒
	Interval int64 `json:"interval"`
	//是hotkey的话在本地的缓存时间，单位毫秒秒
	Duration int64 `json:"duration"`
}
