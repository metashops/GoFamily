### 缓存穿透

1、缓存穿透？

就是大量请求去查询一条记录，先到Redis查询，然后再到数据库都查不到该条记录。这样每次请求都会打到数据库上，导致数据库压力增加，这样我们称为缓存穿透。假如1秒有5000请求，其中有4000请求都是缓存和数据查询不到的，这样会导致数据库扛不住，甚至会导致数据库崩溃。比如遭到黑客恶意攻击，会出现缓存穿透。

2、解决方案

（1）空对象缓存：如果一个查返回的数据为空，不管该数据具体是否存在，都要把这个空值结果进行缓存，设置空结果的过期时间会很短。

（2）设置可访问的名单：使用Bitmaps类型定义一个访问名单，名单的ID作为bitmaps的偏移量，每次访问和bitmap里面的ID进行比较，如果访问不在bitmaps里面，进行拦截。

（3）采用Google的guava中的布隆过滤器：Bloom filter，是1970年是布隆提出的，它实际上是一个很长的二进制向量(位图)和一系列随机映射函数(哈希函数)

（4）进行实时监控：当发现Redis命中率急速下降，则进行排查。



### 缓存击穿

1、什么缓存击穿？什么时候会出现缓存击穿？

缓存中某个热点key突然失效了，然后大量请求直接打到数据库上，这样就会造成某一时刻数据库请求压力过大。

2、解决方案

（1）预先设置热门数据：在redis高峰访问之前，把一些热门数据提前存入到redis里面，加大这些热门数据的key的时长。

（2）实时调整：现场监控那些数据热门，实时调整key的过期时长

* 对于访问频繁的热点key，干脆不设置过期时间
* 互斥独占锁防止击穿（加独占锁）

* 总结：互斥更新、随机退避、差异失效时间

### 缓存雪崩

1、什么事缓存雪崩？

Redis 主机挂了，Redis全盘崩溃，比如缓存中有大量数据同时过期。

2、解决方案

* 构建多级缓存架构：nginx缓存+redis缓存+其他缓存等

* Redis缓存集群实现高可用：配主从+哨兵这样

* Ehcache 本地缓存 + Hystrix 或者使用阿里的sentinel限流 和 服务降级

* 需要开启Redis持久化机制AOF/RDB，启动服务尽快恢复缓存集群。

