package main

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	/**
	1. k8s 配置文件
	2. 保证能通过这个配置文件连接到集群
	*/

	// 1、加载配置文件，生成 config 对象
	config, err := clientcmd.BuildConfigFromFlags("", "../../config")
	if err != nil {
		panic(err.Error())
	}

	// 2、配置 API 路径
	config.APIPath = "api" // pods, /api/v1/pods

	// 3、配置分组版本
	config.GroupVersion = &corev1.SchemeGroupVersion // Group = "",Version: "v1"

	// 4、配置数据的编解码工具
	config.NegotiatedSerializer = &scheme.Codecs

	// 5、实例化 RESTClient 对象
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err.Error())
	}

	// 6、定义接收返回值的变量(接收什么类型的数据)
	result := &corev1.PodList{}

	// 7、跟APIServer交互
	err = restClient.Get(). // Get 请求
				Namespace("default").                                          // 指定名称空间
				Resource("pods").                                              // 指定需要查询的资源，资源名称
				VersionedParams(&metav1.ListOptions{}, scheme.ParameterCodec). // 参数及参数系列化工具
				Do(context.TODO()).                                            // 触发请求
				Into(result)                                                   // 写入返回结果
	if err != nil {
		panic(err.Error())
	}
	for _, item := range result.Items {
		fmt.Printf("namespace:%v name:%v\n", item.Namespace, item.Name)
	}
}
