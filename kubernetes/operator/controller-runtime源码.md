自定义控制器底层是如何实现的？

controller 控制器封装 workqueue 、SetFields、Started等。

其中很重要的两个 Watch 和 Start 方法。 Watch 类似于把我们 Event 监听函数注册去接收事件，再把这些数据传入到 workqueue 里头； Start 就是去启动控制循环，不断从 workqueue 拿数据去消费的。

> 怎么地方调用 Watch 和 Start 方法？在 controller 使用 manager 接口来进行管理控制器的。

### 了解 GV & GVK & GVR 缩写

- GV: Api Group & Version
  - API Group 是相关 API 功能的集合
  - 每个 Group 拥有一或多个 Versions
- GVK: Group Version Kind
  - 每个 GV 都包含 N 个 api 类型，称之为 `Kinds`，不同 `Version` 同一个 `Kinds` 可能不同
- GVR: Group Version Resource
  - `Resource` 是 `Kind` 的对象标识，一般来 `Kind` 和 `Resource` 是 `1:1` 的，但是有时候存在 `1:n` 的关系，不过对于 Operator 来说都是 `1:1` 的关系

实例：源码在位置 examples/crd/main.go

```go
func main() {
	ctrl.SetLogger(zap.New())
	// 根据config创建manager，GetConfigOrDie()默认读取配置是 ～/.kube/config
  // 这样就可以连接 apiserver，然后可以获取client等，最后可以访问资源对象了
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	// in a real controller, we'd create a new scheme for this
  // 将api注册到scheme，scheme给我们提供了GVK到types映射能力（多个CRD就需要多次调用AddToScheme）
	err = api.AddToScheme(mgr.GetScheme())
	if err != nil {
		setupLog.Error(err, "unable to add scheme")
		os.Exit(1)
	}

	err = ctrl.NewControllerManagedBy(mgr).
		For(&api.ChaosPod{}).
		Owns(&corev1.Pod{}).
		Complete(&reconciler{
			Client: mgr.GetClient(),
			scheme: mgr.GetScheme(),
		})
	if err != nil {
		setupLog.Error(err, "unable to create controller")
		os.Exit(1)
	}

	err = ctrl.NewWebhookManagedBy(mgr).
		For(&api.ChaosPod{}).
		Complete()
	if err != nil {
		setupLog.Error(err, "unable to create webhook")
		os.Exit(1)
	}

	setupLog.Info("starting manager")
  // 启动manager（启动注册到manager中的controller）
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
```

make run ENDLE_WEBHOOKS=false