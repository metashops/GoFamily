### Redis三大删除策略

问题：如果一个键过期了，它会立马从内存被删除吗？

分析问题？

该问题是考察你对Redis删除策略的理解，最终面试想听的是Redis内存淘汰机制问题。

回答：

不是的，当Redis的key过期会有三种情况删除，立即删除、惰性删除、定期删除

立即删除：这种方式对内存友好，但是对CPU消耗很大，就是当CPU很忙的时候，然而key过期时间又到了，删除是需要占用CPU时间的，所有这样对给CPU额外的压力。（拿时间换空间）

惰性删除：数据过期了，但不做处理，要等到下一次被访问到发现过期了，就删除。这种方式有可能会导致的有些过期数据永远不会被删除，因为有些数据只访问一次就结束了。（拿空间换时间）

定期删除：key每隔一段时间执行一次删除过期的操作。

上面三种方式删除过期数据都是有很大缺陷的，那么就引出内存淘汰策略，也是面试最终想听的答案。

内存淘汰策略有8种策略：

* noeviction：不对任务key进行删除
* Volatile-lru
* Allkeys-lru
* Volatile-lfu
* Allkeys-lfu
* volatile-random
* Allkeys-random
* Volatile-ttl

有三种算法：先进先出算法、LRU算法及LFU算法

引出：默认使用哪种，你是使用哪种？

大概是两种维度：一个是对所有设置了过期时间的key中删除，另一个是所有key种删除，然后使用LRU、LFU、random及ttl方式进行删除



### 五大数据结构底层源码

> 一切皆字典——RedisObject
>
> 动态字符串sds.c
>
> 集合inset.c
>
> 压缩列表zippiest.c
>
> 快速链表quickest.c
>
> 字典dict.c
>
> streams的底层实现结构listpack.c和rax.c



### mysql 与 缓存数据一致性

问题：你如何解决数据一致问题？双写一致性，你先动缓存还是数据库？

阿里巴巴 MySQL binlog 增量订阅&消费组件的 Canal (管道)

Canal 原理：订阅mysql变化，