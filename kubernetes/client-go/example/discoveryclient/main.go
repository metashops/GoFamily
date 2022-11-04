package main

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 1、加载配置文件,生成 config文件
	config, err := clientcmd.BuildConfigFromFlags("", "../../config")
	if err != nil {
		panic(err.Error())
	}

	// 2、实例化DiscoveryClient
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// 3、发送请求获取GVR数据
	_, apiResource, err := discoveryClient.ServerGroupsAndResources()
	if err != nil {
		panic(err.Error())
	}
	for _, list := range apiResource {
		gv, err := schema.ParseGroupVersion(list.GroupVersion)
		if err != nil {
			panic(err.Error())
		}
		for _, resource := range list.APIResources {
			fmt.Printf("name:%v,group:%v,version:%v\n", resource.Name, gv.Group, gv.Version)
		}
	}
}
