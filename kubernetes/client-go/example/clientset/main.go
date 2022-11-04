package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 1、加载配置文件，生成 config 对象
	config, err := clientcmd.BuildConfigFromFlags("", "../../config")
	if err != nil {
		panic(err.Error())
	}
	// 2、实例化 ClientSet
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// 3、
	pods, err := clientSet.CoreV1().
		Pods("default").
		List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, item := range pods.Items {
		fmt.Printf("namespace:%v name:%v\n", item.Namespace, item.Name)
	}
}
