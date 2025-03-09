package conf

import (
	"context"
	"douyin_mall/api/biz/middleware"
	nacosUtils "douyin_mall/common/infra/nacos"
	"github.com/alibaba/sentinel-golang/ext/datasource"
	sentinelNacosDataSource "github.com/alibaba/sentinel-golang/pkg/datasource/nacos"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/kr/pretty"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/yaml.v2"
	"os"
	"sync"
)

var (
	conf *Config
	once sync.Once
)

type Config struct {
	Env string

	Hertz     Hertz     `yaml:"hertz"`
	Jaeger    Jaeger    `yaml:"jaeger"`
	Kafka     Kafka     `yaml:"kafka"`
	Alert     Alert     `yaml:"alert"`
	AliPay    AliPay    `yaml:"alipay"`
	Redis     Redis     `yaml:"redis"`
	LuaScript LuaScript `yaml:"lua_script"`
}

type Hertz struct {
	Service         string   `yaml:"service"`
	Address         string   `yaml:"address"`
	MetricsPort     string   `yaml:"metrics_port"`
	EnablePprof     bool     `yaml:"enable_pprof"`
	EnableGzip      bool     `yaml:"enable_gzip"`
	EnableAccessLog bool     `yaml:"enable_access_log"`
	LogLevel        string   `yaml:"log_level"`
	LogFileName     string   `yaml:"log_file_name"`
	LogMaxSize      int      `yaml:"log_max_size"`
	LogMaxBackups   int      `yaml:"log_max_backups"`
	LogMaxAge       int      `yaml:"log_max_age"`
	RegistryAddr    []string `yaml:"registry_addr"`
}

type Jaeger struct {
	Endpoint string `yaml:"endpoint"`
}

type Kafka struct {
	ClsKafka ClsKafka `yaml:"cls_kafka"`
}

type ClsKafka struct {
	Usser    string `yaml:"user"`
	Password string `yaml:"password"`
	TopicId  string `yaml:"topic_id"`
}

type Alert struct {
	FeishuWebhook string `yaml:"feishu_webhook"`
}
type AliPay struct {
	AlipayPublicContentPath string `yaml:"ali_pay_public_content_path"`
}
type LuaScript struct {
	LimitScript string `yaml:"limit_script"`
}

type Redis struct {
	Address  string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// GetConf gets configuration instance
func GetConf() *Config {
	once.Do(initConf)
	return conf
}

func initConf() {
	clientConfig, serverConfigs := nacosUtils.GetNacosConfig()

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		hlog.Fatalf("初始化nacos配置客户端失败: %v", err)
	}

	// 定义配置项
	configs := []struct {
		DataId        string
		Group         string
		Type          vo.ConfigType
		UnmarshalFunc func([]byte) error
	}{
		{
			DataId: "api_conf.yaml",
			Group:  "DEFAULT_GROUP",
			Type:   vo.YAML,
			UnmarshalFunc: func(data []byte) error {
				if conf == nil {
					conf = new(Config)
				}
				if err := yaml.Unmarshal(data, conf); err != nil {
					return err
				}
				conf.Env = GetEnv()
				return nil
			},
		},
		{
			DataId: "uri_whitelist_config",
			Group:  "DEFAULT_GROUP",
			Type:   vo.JSON,
			UnmarshalFunc: func(data []byte) error {
				middleware.Whitelist = make(map[string]struct{})
				return sonic.Unmarshal(data, &middleware.Whitelist)
			},
		},
	}
	//sentinel规则nacos动态配置
	sentinelRuleInit(configClient)

	// 监听与初始化配置
	for _, cfg := range configs {
		listenAndLoadConfig(configClient, cfg)
	}
}

func sentinelRuleInit(configClient config_client.IConfigClient) {
	//sentine流控规则nacos动态配置
	flowRulesHandler := datasource.NewFlowRulesHandler(datasource.FlowRuleJsonArrayParser)

	flowSource, err := sentinelNacosDataSource.NewNacosDataSource(configClient, "DEFAULT_GROUP", "sentinel_api_flow_rules.json", flowRulesHandler)
	if err != nil {
		klog.CtxErrorf(context.Background(), "获取配置失败规则失败: %v", err)
		return
	}
	err = flowSource.Initialize()
	if err != nil {
		klog.CtxErrorf(context.Background(), "初始化sentinel规则失败: %v", err)
		return
	}
	hotSpotParamRulesHandler := datasource.NewHotSpotParamRulesHandler(datasource.HotSpotParamRuleJsonArrayParser)
	hotspotSource, err := sentinelNacosDataSource.NewNacosDataSource(configClient, "DEFAULT_GROUP", "sentinel_api_hot_spot_rules.json", hotSpotParamRulesHandler)

	if err != nil {
		klog.CtxErrorf(context.Background(), "初始化sentinel规则失败: %v", err)
	}
	err = hotspotSource.Initialize()
	if err != nil {
		klog.CtxErrorf(context.Background(), "初始化sentinel规则失败: %v", err)
	}
	klog.Info("初始化sentinel规则成功")
}
func listenAndLoadConfig(client config_client.IConfigClient, cfg struct {
	DataId        string
	Group         string
	Type          vo.ConfigType
	UnmarshalFunc func([]byte) error
}) {
	// 监听配置
	err := client.ListenConfig(vo.ConfigParam{
		DataId: cfg.DataId,
		Group:  cfg.Group,
		Type:   cfg.Type,
		OnChange: func(namespace, group, dataId, data string) {
			if err := cfg.UnmarshalFunc([]byte(data)); err != nil {
				hlog.Errorf("解析配置 %s 失败: %v", cfg.DataId, err)
				return
			}
			prettyPrint(cfg.DataId)
		},
	})
	if err != nil {
		hlog.Errorf("监听配置 %s 失败: %v", cfg.DataId, err)
	}

	// 初始化配置
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: cfg.DataId,
		Group:  cfg.Group,
		Type:   cfg.Type,
	})
	if err != nil {
		hlog.Fatalf("获取配置 %s 失败: %v", cfg.DataId, err)
	}

	if err := cfg.UnmarshalFunc([]byte(content)); err != nil {
		hlog.Fatalf("解析配置 %s 失败: %v", cfg.DataId, err)
	}

	prettyPrint(cfg.DataId)
}

func prettyPrint(name string) {
	var data interface{}
	if name == "api_conf.yaml" {
		data = conf
	} else {
		data = middleware.Whitelist
	}
	if _, err := pretty.Printf("%+v\n", data); err != nil {
		hlog.Errorf("pretty print error: %v", err)
	}
}

func GetEnv() string {
	e := os.Getenv("env")
	if len(e) == 0 {
		return "test"
	}
	return e
}

func LogLevel() hlog.Level {
	level := GetConf().Hertz.LogLevel
	switch level {
	case "trace":
		return hlog.LevelTrace
	case "debug":
		return hlog.LevelDebug
	case "info":
		return hlog.LevelInfo
	case "notice":
		return hlog.LevelNotice
	case "warn":
		return hlog.LevelWarn
	case "error":
		return hlog.LevelError
	case "fatal":
		return hlog.LevelFatal
	default:
		return hlog.LevelInfo
	}
}
