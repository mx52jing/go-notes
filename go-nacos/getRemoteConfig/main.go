package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"gopkg.in/yaml.v3"
	"time"
)

type AllConfig struct {
	Mysql  MysqlConfig  `yaml:"mysql"`
	Redis  RedisConfig  `yaml:"redis"`
	Wechat WechatConfig `yaml:"wechat"`
}

type MysqlConfig struct {
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type WechatConfig struct {
	AppID     string `yaml:"appID"`
	AppSecret string `yaml:"appSecret"`
	BaseURL   string `yaml:"baseURL"`
}

func convertYamlToStruct(in []byte, out interface{}) error {
	return yaml.Unmarshal(in, out)
}

func fetchRemoteConfig() {
	// create server config
	serverConfigs := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848, constant.WithContextPath("/nacos")),
	}
	// create naming clientConfig
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId("d90cfc52-948b-41e7-82f7-bd22dbb36221"),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
		constant.WithUsername("nacos"),
		constant.WithPassword("nacos"),
	)
	// create config client for dynamic configuration
	namingClient, err := clients.NewConfigClient(vo.NacosClientParam{
		ServerConfigs: serverConfigs,
		ClientConfig:  &clientConfig,
	})
	if err != nil {
		panic(err)
	}
	configParam := vo.ConfigParam{
		DataId: "test-config",
		Group:  "DEFAULT_GROUP",
		Type:   "yaml",
	}
	remoteConfig, err := namingClient.GetConfig(configParam)
	var allConfig AllConfig
	err = convertYamlToStruct([]byte(remoteConfig), &allConfig)
	//fmt.Printf("allConfig is %v \n", allConfig.Wechat)
	if err != nil {
		panic(err)
	}
	err = namingClient.ListenConfig(vo.ConfigParam{
		DataId: "test-config",
		Group:  "DEFAULT_GROUP",
		Type:   "yaml",
		OnChange: func(namespace, group, dataId, data string) {
			//fmt.Printf("[namespace: %s], [group: %s], [dataId: %s] changed, content is 【%s】, contentType is [%T]\n", namespace, group, dataId, data, data)
			err = convertYamlToStruct([]byte(data), &allConfig)
			if err != nil {
				panic(err)
			}
			fmt.Printf("修改后的allConfig为%v \n", allConfig)
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("startup successful")
	time.Sleep(5 * time.Second)
	// search config
	searchedConfig, err := namingClient.SearchConfig(vo.SearchConfigParam{
		Search:   "blur", // accurate/blur 准确/模糊搜索
		DataId:   "",
		Group:    "",
		PageNo:   1,
		PageSize: 10,
	})
	if err != nil {
		fmt.Printf("search config error %s \n", err)
		return
	}
	fmt.Printf("searchd config is %v \n", searchedConfig)
	time.Sleep(60 * time.Second)
}

func main() {
	fetchRemoteConfig()
}
