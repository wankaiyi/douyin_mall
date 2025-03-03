package conf

import (
	"context"
	nacosUtils "douyin_mall/common/infra/nacos"
	"fmt"
	"github.com/alibaba/sentinel-golang/ext/datasource"
	"github.com/kitex-contrib/config-nacos/nacos"
	"github.com/kr/pretty"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"os"
	"sync"

	sentinelNacosDataSource "github.com/alibaba/sentinel-golang/pkg/datasource/nacos"
	"github.com/cloudwego/kitex/pkg/klog"
	"gopkg.in/yaml.v2"
)

var (
	conf *Config
	once sync.Once
)

type Config struct {
	Env      string
	Kitex    Kitex    `yaml:"kitex"`
	MySQL    MySQL    `yaml:"mysql"`
	Redis    Redis    `yaml:"redis"`
	Registry Registry `yaml:"registry"`
	Alert    Alert    `yaml:"alert"`
	Kafka    Kafka    `yaml:"kafka"`
	Jaeger   Jaeger   `yaml:"jaeger"`
}

type MySQL struct {
	DSN string `yaml:"dsn"`
}

type Redis struct {
	Address  string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type Kitex struct {
	Service       string `yaml:"service"`
	Address       string `yaml:"address"`
	MetricsPort   string `yaml:"metrics_port"`
	LogLevel      string `yaml:"log_level"`
	LogFileName   string `yaml:"log_file_name"`
	LogMaxSize    int    `yaml:"log_max_size"`
	LogMaxBackups int    `yaml:"log_max_backups"`
	LogMaxAge     int    `yaml:"log_max_age"`
}

type Registry struct {
	RegistryAddress string `yaml:"registry_address"`
	Username        string `yaml:"username"`
	Password        string `yaml:"password"`
}

type Alert struct {
	FeishuWebhook string `yaml:"feishu_webhook"`
}

type Kafka struct {
	ClsKafka ClsKafka `yaml:"cls_kafka"`
}

type ClsKafka struct {
	Usser    string `yaml:"user"`
	Password string `yaml:"password"`
	TopicId  string `yaml:"topic_id"`
}

type Jaeger struct {
	Endpoint string `yaml:"endpoint"`
}

// GetConf gets configuration instance
func GetConf() *Config {
	once.Do(initConf)
	return conf
}

func initConf() {
	client, err := nacos.NewClient(nacos.Options{
		Address:     os.Getenv("NACOS_ADDR"),
		NamespaceID: "e45ccc29-3e7d-4275-917b-febc49052d58",
		Group:       "DEFAULT_GROUP",
		Username:    "nacos",
		Password:    os.Getenv("NACOS_PASSWORD"),
		Port:        8848,
	})
	if err != nil {
		panic(err)
	}
	param := vo.ConfigParam{
		DataId: "cart_conf.yaml",
		Group:  "DEFAULT_GROUP",
		Type:   "yaml",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Printf("Config changed - namespace: %s, group: %s, data-id: %s\n", namespace, group, dataId)

			// 解析 YAML 配置
			var config interface{}
			err := yaml.Unmarshal([]byte(data), &config)
			if err != nil {
				fmt.Printf("Error parsing YAML: %v\n", err)
				return
			}

			// 输出解析结果
			fmt.Printf("Parsed YAML: %v\n", config)
		},
	}

	client.RegisterConfigCallback(param, func(data string, parser nacos.ConfigParser) {
		// 处理配置数据的逻辑
		if conf == nil {
			conf = new(Config)
		}
		err := yaml.Unmarshal([]byte(data), &conf)
		if err != nil {
			klog.Error("Error parsing YAML: %v\n", err)
			return
		}
		_, err = pretty.Printf("%+v\n", conf)
		if err != nil {
			klog.Error("pretty print error - %v", err)
		}
	}, 5000)
	conf.Env = GetEnv()

	clientConfig, serverConfigs := nacosUtils.GetNacosConfig()

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		klog.Fatalf("初始化nacos配置客户端失败: %v", err)
	}

	//sentinel规则nacos动态配置
	sentinelRuleInit(configClient)

}
func sentinelRuleInit(configClient config_client.IConfigClient) {
	//sentine流控规则nacos动态配置
	flowRulesHandler := datasource.NewFlowRulesHandler(datasource.FlowRuleJsonArrayParser)
	//熔断规则nacos动态配置
	circuitBreakerRulesHandler := datasource.NewCircuitBreakerRulesHandler(datasource.CircuitBreakerRuleJsonArrayParser)

	flowSource, err := sentinelNacosDataSource.NewNacosDataSource(configClient, "DEFAULT_GROUP", "sentinel_cart_flow_rules.json", flowRulesHandler)
	if err != nil {
		klog.CtxErrorf(context.Background(), "获取sentinel流控规则失败: %v", err)
		return
	}
	err = flowSource.Initialize()
	if err != nil {
		klog.CtxErrorf(context.Background(), "初始化sentinel规则失败: %v", err)
		return
	}
	circuitBreakerSource, err := sentinelNacosDataSource.NewNacosDataSource(configClient, "DEFAULT_GROUP", "sentinel_cart_circuit_breaker_rules.json", circuitBreakerRulesHandler)
	if err != nil {
		klog.CtxErrorf(context.Background(), "获取sentinel熔断规则失败: %v", err)
		return
	}
	err = circuitBreakerSource.Initialize()
	if err != nil {
		klog.CtxErrorf(context.Background(), "初始化sentinel熔断规则失败: %v", err)
		return
	}

	klog.Info("初始化sentinel规则成功")
}

func GetEnv() string {
	e := os.Getenv("env")
	if len(e) == 0 {
		return "test"
	}
	return e
}

func LogLevel() klog.Level {
	level := GetConf().Kitex.LogLevel
	switch level {
	case "trace":
		return klog.LevelTrace
	case "debug":
		return klog.LevelDebug
	case "info":
		return klog.LevelInfo
	case "notice":
		return klog.LevelNotice
	case "warn":
		return klog.LevelWarn
	case "error":
		return klog.LevelError
	case "fatal":
		return klog.LevelFatal
	default:
		return klog.LevelInfo
	}
}
