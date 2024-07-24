package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"log"
	"time"
)

func noNameSpaceOperation() {
	// create  serverConfig
	serverConfigs := []constant.ServerConfig{
		*constant.NewServerConfig(
			"127.0.0.1",
			8848,
			constant.WithContextPath("/nacos"),
		),
	}
	// create clientConfig
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId(""),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
		constant.WithUsername("nacos"),
		constant.WithPassword("nacos"),
	)
	// create config client
	client, err := clients.NewConfigClient(vo.NacosClientParam{
		ClientConfig:  &clientConfig,
		ServerConfigs: serverConfigs,
	})
	if err != nil {
		log.Printf("create config client error %s\n", err)
		panic(err)
	}
	// publish config
	_, err = client.PublishConfig(vo.ConfigParam{
		DataId:  "test-dataId-1",
		Group:   "test-group1",
		Content: "Hello, I am test data1",
	})
	_, err = client.PublishConfig(vo.ConfigParam{
		DataId:  "test-dataId-2",
		Group:   "test-group2",
		Type:    "yaml", // 提交的数据类型
		Content: "Hello, I am test data2 change",
	})
	if err != nil {
		log.Printf("publish config error => %s\n", err)
		return
	}
	time.Sleep(2 * time.Second)

	// get config
	testGroup1Content, err := client.GetConfig(vo.ConfigParam{
		DataId: "test-dataId-1",
		Group:  "test-group1",
	})
	if err != nil {
		log.Printf("get test-group1 config error => %s\n", err)
		return
	}
	fmt.Printf("testGroup1Content is： %s \n", testGroup1Content)
	testGroup2Content, err := client.GetConfig(vo.ConfigParam{
		DataId: "test-dataId-2",
		Group:  "test-group2",
	})
	if err != nil {
		log.Printf("get test-group2 config error => %s\n", err)
		return
	}
	fmt.Printf("testGroup2Content is： %s \n", testGroup2Content)
	// listen config change
	err = client.ListenConfig(vo.ConfigParam{
		DataId: "test-dataId-1",
		Group:  "test-group1",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Printf("[namespace: %s], [group: %s], [dataId: %s] changed, content is 【%s】\n", namespace, group, dataId, data)
		},
	})

	err = client.ListenConfig(vo.ConfigParam{
		DataId: "test-dataId-2",
		Group:  "test-group2",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Printf("[namespace: %s], [group: %s], [dataId: %s] changed, content is 【%s】\n", namespace, group, dataId, data)
		},
	})
	if err != nil {
		return
	}
	time.Sleep(30 * time.Second)
	// delete config
	_, err = client.DeleteConfig(vo.ConfigParam{
		DataId: "test-dataId-2",
		Group:  "test-group2",
	})
	if err != nil {
		return
	}
	fmt.Println("delete test-group2 success")
	// cancel listen config change
	err = client.CancelListenConfig(vo.ConfigParam{
		DataId: "test-dataId-1",
		Group:  "test-group1",
	})
	if err != nil {
		return
	}
	fmt.Println("cancel listen test-group1 success")
	time.Sleep(30 * time.Second)
	// search config
	searchedConfig, err := client.SearchConfig(vo.SearchConfigParam{
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
}

func withNameSpaceOperation() {
	// create  serverConfig
	serverConfigs := []constant.ServerConfig{
		*constant.NewServerConfig(
			"127.0.0.1",
			8848,
			constant.WithContextPath("/nacos"),
		),
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
		fmt.Printf("create naming client error")
		return
	}
	// publish config
	_, err = namingClient.PublishConfig(vo.ConfigParam{
		DataId:  "naming-cms-dataId-1",
		Group:   "DEFAULT_GROUP",
		Content: "hello，naming client cms",
	})
	if err != nil {
		fmt.Printf("publish naming client error => %s \n", err)
		return
	}
	fmt.Println("publish naming client success")
	// get config
	gotConfig, err := namingClient.GetConfig(vo.ConfigParam{
		DataId: "naming-cms-dataId-1",
		Group:  "DEFAULT_GROUP",
	})
	fmt.Printf("got naming client config => %s \n", gotConfig)

}

func main() {
	//noNameSpaceOperation()
	withNameSpaceOperation()
}
