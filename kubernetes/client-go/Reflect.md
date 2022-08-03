前面我们说了 Informer 通过对 APIServer 的资源对象执行 List 和 Watch 操作，把获取到的数据存储在本地的缓存中，其中实现这个的核心功能就是 Reflector，我们可以称其为反射器，从名字我们可以看出来它的主要功能就是反射，就是将 Etcd 里面的数据反射到本地存储（DeltaFIFO）中。Reflector 首先通过 List 操作获取所有的资源对象数据，保存到本地存储，然后通过 Watch 操作监控资源的变化，触发相应的事件处理，比如前面示例中的 Add 事件、Update 事件、Delete 事件。Reflector 结构体的定义位于 `staging/src/k8s.io/client-go/tools/cache/reflector.go` 下面：

```go
// k8s.io/client-go/tools/cache/reflector.go

// Reflector(反射器) 监听指定的资源，将所有的变化都反射到给定的存储中去
type Reflector struct {
	// name 标识这个反射器的名称，默认为 文件:行数（比如reflector.go:125）
  // 默认名字通过 k8s.io/apimachinery/pkg/util/naming/from_stack.go 下面的 GetNameFromCallsite 函数生成
	name string

  // 期望放到 Store 中的类型名称，如果提供，则是 expectedGVK 的字符串形式
  // 否则就是 expectedType 的字符串，它仅仅用于显示，不用于解析或者比较。
	expectedTypeName string
	// An example object of the type we expect to place in the store.
	// Only the type needs to be right, except that when that is
	// `unstructured.Unstructured` the object's `"apiVersion"` and
	// `"kind"` must also be right.
  // 我们放到 Store 中的对象类型
	expectedType reflect.Type
	// 如果是非结构化的，我们期望放在 Sotre 中的对象的 GVK
	expectedGVK *schema.GroupVersionKind
	// 与 watch 源同步的目标 Store
	store Store
	// 用来执行 lists 和 watches 操作的 listerWatcher 接口（最重要的）
	listerWatcher ListerWatcher

	// backoff manages backoff of ListWatch
	backoffManager wait.BackoffManager

	resyncPeriod time.Duration
	// ShouldResync 会周期性的被调用，当返回 true 的时候，就会调用 Store 的 Resync 操作
	ShouldResync func() bool
	// clock allows tests to manipulate time
	clock clock.Clock
	// paginatedResult defines whether pagination should be forced for list calls.
	// It is set based on the result of the initial list call.
	paginatedResult bool
	// Kubernetes 资源在 APIServer 中都是有版本的，对象的任何修改(添加、删除、更新)都会造成资源版本更新，lastSyncResourceVersion 就是指的这个版本
	lastSyncResourceVersion string
	// 如果之前的 list 或 watch 带有 lastSyncResourceVersion 的请求中是一个 HTTP 410（Gone）的失败请求，则 isLastSyncResourceVersionGone 为 true
	isLastSyncResourceVersionGone bool
	// lastSyncResourceVersionMutex 用于保证对 lastSyncResourceVersion 的读/写访问。
	lastSyncResourceVersionMutex sync.RWMutex
	// WatchListPageSize is the requested chunk size of initial and resync watch lists.
	// If unset, for consistent reads (RV="") or reads that opt-into arbitrarily old data
	// (RV="0") it will default to pager.PageSize, for the rest (RV != "" && RV != "0")
	// it will turn off pagination to allow serving them from watch cache.
	// NOTE: It should be used carefully as paginated lists are always served directly from
	// etcd, which is significantly less efficient and may lead to serious performance and
	// scalability problems.
	WatchListPageSize int64
}

// NewReflector 创建一个新的反射器对象，将使给定的 Store 保持与服务器中指定的资源对象的内容同步。
// 反射器只把具有 expectedType 类型的对象放到 Store 中，除非 expectedType 是 nil。
// 如果 resyncPeriod 是非0，那么反射器会周期性地检查 ShouldResync 函数来决定是否调用 Store 的 Resync 操作
// `ShouldResync==nil` 意味着总是要执行 Resync 操作。
// 这使得你可以使用反射器周期性地处理所有的全量和增量的对象。
func NewReflector(lw ListerWatcher, expectedType interface{}, store Store, resyncPeriod time.Duration) *Reflector {
  // 默认的反射器名称为 file:line
	return NewNamedReflector(naming.GetNameFromCallsite(internalPackages...), lw, expectedType, store, resyncPeriod)
}

// NewNamedReflector 与 NewReflector 一样，只是指定了一个 name 用于日志记录
func NewNamedReflector(name string, lw ListerWatcher, expectedType interface{}, store Store, resyncPeriod time.Duration) *Reflector {
	realClock := &clock.RealClock{}
	r := &Reflector{
		name:          name,
		listerWatcher: lw,
		store:         store,
		backoffManager: wait.NewExponentialBackoffManager(800*time.Millisecond, 30*time.Second, 2*time.Minute, 2.0, 1.0, realClock),
		resyncPeriod:   resyncPeriod,
		clock:          realClock,
	}
	r.setExpectedType(expectedType)
	return r
}
// Run repeatedly uses the reflector's ListAndWatch to fetch all the
// objects and subsequent deltas.
// Run will exit when stopCh is closed.
// 把reflector实例化，拿到对象之后，去启动reflector
func (r *Reflector) Run(stopCh <-chan struct{}) {
	klog.V(3).Infof("Starting reflector %s (%s) from %s", r.expectedTypeName, r.resyncPeriod, r.name)
	wait.BackoffUntil(func() {
		if err := r.ListAndWatch(stopCh); err != nil {
			r.watchErrorHandler(r, err)
		}
	}, r.backoffManager, true, stopCh)
	klog.V(3).Infof("Stopping reflector %s (%s) from %s", r.expectedTypeName, r.resyncPeriod, r.name)
}
```

总结：

1. Reflector 利用ClientSet客户端列举（List全量）
2. 将全量对象用 replace() 接口同步到本地缓存 Store(DaltaFIFO)，且更新资源版本号
3. 开启一个协程定时执行 resync操作，同步是把全量对象以同步事件的方式通知出去
4. 然后通过ClientSet客户端监控(Watch)资源，监控当前资源版本号以后的对象
5. 一旦有发生变化，那么根据变化的类型(ADD,UPDATE,DELETE)调用DeltaFIFO的相应接口，产生一个相应的对象Delta，同时更新当前资源版本号