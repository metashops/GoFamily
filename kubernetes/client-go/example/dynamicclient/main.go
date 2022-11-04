package main

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 1、加载配置文件
	config, err := clientcmd.BuildConfigFromFlags("", "../../config")
	if err != nil {
		panic(err.Error())
	}

	// 2、实例化动态客户端对象
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// 3、配置需要调用的GVR
	gvr := schema.GroupVersionResource{
		Group:    "", // 无名资源组，不需要写
		Version:  "v1",
		Resource: "pods",
	}

	// 4、发送请求，且得到返回结果(动态客户端获取到的是非结构化结构体，需要用非结构化变量去存储)
	unStructData, err := dynamicClient.Resource(gvr).Namespace("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	// 5、unStructData 转换为结构化数据
	podList := &corev1.PodList{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unStructData.UnstructuredContent(), podList)
	if err != nil {
		panic(err.Error())
	}

	// 6、打印
	for _, item := range podList.Items {
		fmt.Printf("namespace:%v, name:%v\n", item.Namespace, item.Name)
	}
}
