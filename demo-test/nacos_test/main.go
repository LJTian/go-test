package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"time"
)

var client naming_client.INamingClient

func RegisterInstance(addr string, port uint64,
	serviceName string, clusterName string, groupName string, data map[string]string) {
	_, err := client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          addr,
		Port:        port,
		ServiceName: serviceName,
		Weight:      10,
		ClusterName: clusterName,
		GroupName:   groupName,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    data,
	})
	if err != nil {
		panic(err)
	}
}

func DeregisterInstance(addr string, port uint64,
	serviceName string, clusterName string, groupName string) {
	_, err := client.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          addr,
		Port:        port,
		ServiceName: serviceName,
		Cluster:     clusterName,
		GroupName:   groupName,
		Ephemeral:   true, //it must be true
	})
	if err != nil {
		panic(err)
	}
}

func main() {

	// 使用配置中心
	sc := []constant.ServerConfig{
		{
			IpAddr: "10.211.55.3",
			Port:   8848,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         "65e42af7-c53e-43fe-bb91-b2c0478ded22", //namespace id
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// 创建动态配置客户端的另一种方式 (推荐)
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}
	config, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "dbconfig.yaml",
		Group:  "dev",
	})
	if err != nil {
		return
	}

	fmt.Println(config)

	// 检测配置文件改变
	configClient.ListenConfig(vo.ConfigParam{
		DataId: "dbconfig.yaml",
		Group:  "dev",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("配置文件发生更改")
			fmt.Printf("namespace is %s\n, group is %s\n, dataId is %s\n, data is %s\n",
				namespace, group, dataId, data)
		},
	})

	time.Sleep(3000 * time.Second)

	// 下面是获取服务的相关操作
	//// a more graceful way to create naming client
	//var err error
	//client, err = clients.NewNamingClient(
	//	vo.NacosClientParam{
	//		ClientConfig:  &cc,
	//		ServerConfigs: sc,
	//	},
	//)
	//if err != nil {
	//	panic(err)
	//}
	//

	//
	//// 注册服务
	//RegisterInstance("10.0.0.12", 8848,
	//	"test1", "user_svr", "dev", map[string]string{
	//		"tian": "lijun1",
	//	})
	//
	//RegisterInstance("10.0.0.14", 8848,
	//	"test1", "user_svr", "dev", map[string]string{
	//		"tian": "lijun2",
	//	})
	//
	//RegisterInstance("10.0.0.15", 8848,
	//	"test1", "user_svr", "dev", map[string]string{
	//		"tian": "lijun3",
	//	})
	//
	////// 遍历服务
	////instances, err := client.GetService(vo.GetServiceParam{
	////	Clusters:    []string{"user_svr"},
	////	ServiceName: "test1",
	////	GroupName:   "dev",
	////})
	////if err != nil {
	////	fmt.Println(err)
	////}
	////
	////fmt.Println(instances)
	//
	//// 注销服务
	//DeregisterInstance("10.0.0.15", 8848,
	//	"test1", "user_svr", "dev")
	//fmt.Println("注销服务")
	//
	//// 遍历服务(注销之后，服务为NULL,不知道为什么)
	//instances, err := client.GetService(vo.GetServiceParam{
	//	Clusters:    []string{"user_svr"},
	//	ServiceName: "test1",
	//	GroupName:   "dev",
	//})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(instances)
}
