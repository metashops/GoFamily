## 一、Redis安装

（1）先把GCC安装

```
yum install gcc
```

（2）解压，`-C /opt/module/`安装到model目录下，没有model目录创建就好了

```
tar -zxvf redis-6.2.1.tar.gz -C /opt/module/
```

（3）进入解压后的redis-6.2.1目录，执行该命令：`make`（）就是把当前的redis文件编译

（4）再执行：`make install`，如果你在普通用户上就使用：`sudo make install`

（5）默认安装在`/uer/local/bin`目录下，详细文件如下

* redis-server：Redis 服务器启动命令
* redis-cli：客户端操作
* redis-benchmark：性能测试工具，可以自己本机运行，看看性能
* redis-check-aof：修复有问题的AOF文件
* redis-check-rdb
* redis-sentinel ：Redis集群

（6）启动

* 前台启动直接运行：redis-server（很少使用）
* 后台启动（常用）
  * 第一步首先：备份redis.conf，该文件在你自定义安装的目录redis-6.2.1这个目录下
  * 第二步：这里复制到`sudo cp redis.conf /etc/redis.conf`
  * 第三步：`vi redis.conf`将 daemonized no 该为 yes
  * 第四步：用刚才修改的文件启动redis，进入默认安装的位置执行/usr/local/bin下启动命令：`redis-server /etc/redis.conf`
  * 第五步：查看进程`ps -ef | grep redis`

（7）关闭

单实例关闭：redis-cli shutdown

终端里头直接：shutdown

多实例关闭并指定端口：redis-cli -p 6379 shutdown

（8）设置配置文件

* 网络相关配置该这行注释掉：#bind=127.0.0.1(默认只能本机访问，不能远程)
* 将protected-mode no
* 设置密码将# requirepass foobared去掉，foobared是密码可以设置123456
* 再次启动redis然后进入客户端后需要验证：auth "你设置的密码"
* 配置日志：touch /user/local/bin/redis-log.log，并将权限更改为`sudo chmod 666 redis-log.log`

（9）rdm工具连接



## 二、Redis ——5种基本数据结构

> Redis常见5种数据结构：String**、**Hash**、**List**、**Set**、**SortedSet

### Redis 键(key)

* 查看所有的键：Keys * 
* 判断某个key是否存在：exists [key的名字]
* 给key设置过期时间：expire key [秒钟]
* 查看key还有多少时间过期：`ttl [key] `，-1表示永久不过期，-2表示已经过期
* 查看key是什么类型：type key
* 删除key命令：del [需要删除的key]
* 清空数据库：flushdb



### String

> 使用场景：比如需要统计的场景，热点文章的访问量以及转发量、分布式锁setnx、Web集群session共享

1、介绍

字符串类型是Redis最基础的数据结构。键和值都是字符串类型的，而且其他几种的数据结构都是在字符串类型基础上构建的，所以学好字符串类型，几乎后面的很简单啦。String可以理解与Memcached一模一样的类型，一个key对应一个value。String类型是二进制安全的，可以包含简单字符串、复杂的字符串、JSON、XML、数字，但是字符串value最多可以是512M。

2、常用命令

（1）set/get/append/strlen

（2）Incr/decr/Incrb/decrby：一定要是数字才能进行加减

* key的值自增+1，命令格式：incr k1(需要自增的key)
* key的值自减-1，命令格式：decr k1(需要自减的key)
* key的值自增+n，命令格式：incrby k1 2(需要自增的key,2就是每次增加2)
* key的值自减-n，命令格式：decrby k1 2(需要自增的key,2就是每次减2)

（3）getrange/setrange

* 获取指定区间的值，类似between...and的关系：getrange key 0 2（0-2说明我要获取0-2之间的值两端可以取到）
* 设置指定区间范围的值，格式是setrange key值 具体值(如：setrange k3 0 888表示从0位开始设置为888，888占3位)

（4）setex(set with expire)/setnx(set if not exist)

* 键秒值，意思加入某k1=23，现在设置为20后失效：setex k1 20 23（只要超过20秒那么键和值全失效）

（5）mset/mget/msetnx

* mset：批量操作

* mget：批量获取

  ```redis
  127.0.0.1:6379> mset k1 v1 k2 v2 k3 v3
  OK
  127.0.0.1:6379> keys *
  1) "k1"
  2) "k2"
  3) "k3"
  127.0.0.1:6379> set k1 v1 k2 v2 k3 v3
  (error) ERR syntax error
  # 获取多个值
  127.0.0.1:6379> mget k1 k2 k3
  1) "v1"
  2) "v2"
  3) "v3"
  ```

#### 总结：String操作及使用场景

字符串常用操作：

>Set key value      //存入字符串键值对
>
>Mset key1 value1 key1 value1...  //批量存储字符串键值对操作
>
>Setnx key value  //存入一个不存在的字符串键值对
>
>GET key  //获取一个字符串键值对
>
>MGET key1 key2... //批量获取字符串键值对
>
>DEL key1 //删除一个键
>
>EXPIRE key secondes //设置一个键过期时间（秒）如：expire key 10

原子加减

> INCR key //将key中存储的数字值+1
>
> DECR key //将key中存储的数字值-1
>
> INCRBY key increment //将key所有存储的值加上increment
>
> DECRBY key decrement //将key所有存储的值减去decrement

#### 

### List

> 发布与订阅或者消息队列、慢查询

使用List命令(rpush和Lpop)实现：队列（右进左出）

```
127.0.0.1:6379> rpush list02 1 2 3 4 5
(integer) 5
127.0.0.1:6379> lpop list02
"1"
127.0.0.1:6379> lpop list02
"2"
127.0.0.1:6379> lpop list02
"3"
127.0.0.1:6379> lpop list02
"4"
127.0.0.1:6379> lpop list02
"5"
```

![6AB38AC5-0A18-49FA-B43F-5FDA342092B2.png](http://ww1.sinaimg.cn/large/006FuVcvgy1gupgfqrj5xj60xw07agmd02.jpg)

使用List命令(rpush和rpop)实现：栈（右进右出）

```
127.0.0.1:6379> rpush list01 1 2 3 4 5 
(integer) 5
127.0.0.1:6379> rpop list01
"5"
127.0.0.1:6379> rpop list01
"4"
127.0.0.1:6379> rpop list01
"3"
127.0.0.1:6379> rpop list01
"2"
127.0.0.1:6379> rpop list01
"1"
```

![B7ADFB85-0FF2-4300-A3E8-F18801007E26.png](http://ww1.sinaimg.cn/large/006FuVcvgy1gupgjdwkfxj60si07qaap02.jpg)

（3）lindex,按照索引下标获取元素(下标0开始)

```
127.0.0.1:6379> rpush list02 1 2 3 4 5
(integer) 5
127.0.0.1:6379> lindex list02 0
"1"
127.0.0.1:6379> lindex list02 2
"3"
127.0.0.1:6379>
```

（4）llen获取长度：llen list

（5）rpoplpush

```redis
127.0.0.1:6379> lrange list 0 -1
1) "4"
2) "3"
3) "2"
4) "1"
5) "0"
6) "-1"
127.0.0.1:6379> lrange list2 0 -1
1) "10"
2) "9"
3) "8"
127.0.0.1:6379> rpoplpush list list2
"-1"
127.0.0.1:6379> lrange list2 0 -1
1) "-1"
2) "10"
3) "9"
4) "8"
```

（6）insert key before/after 值1 值2 ；在某位置插入值



### Set

> 要存储的数据不能重复以及需要获取多个数据源的交集和并集的这样场景。（比如：微信微博的收藏、点赞、标签）

Redis的set相当于Java语言的HashSet，它内部结构的键值对是无序的唯一的，有去重功能。

（1）sadd/smembers/sismember

```
127.0.0.1:6379> sadd set0 1 2 3 4 5
127.0.0.1:6379> smembers set0
1) "1"
2) "2"
3) "3"
4) "4"
5) "5"
127.0.0.1:6379>
// 有就返回1，没有就返回0
127.0.0.1:6379> sismember set0 x
(integer) 0
127.0.0.1:6379> sismember set0 3
(integer) 1
127.0.0.1:6379>
```

（2）scard,获取集合里面元素个数

```
127.0.0.1:6379> scard set0
(integer) 5
```

（3）srem key value 删除集合元素

```
127.0.0.1:6379> srem set0 4
(integer) 1
127.0.0.1:6379> sismember set0 4
(integer) 0
```

（4）srandmember key n（假如n=3，那么随机出3个数）

```
127.0.0.1:6379> sadd set1 1 2 3 4 5 6 7 8 9 10 11
(integer) 11
127.0.0.1:6379> srandmember set1 5
1) "6"
2) "7"
3) "1"
4) "8"
5) "5"
127.0.0.1:6379> srandmember set1 2
1) "1"
2) "9"
127.0.0.1:6379>
```

（5）spop key 随机出栈

```
127.0.0.1:6379> spop set1
"11"
```

（6）smoke key1 key2 在key1里某个值，作用是将key1里面的某个值赋给key2

（7）数学集合

* 差集:sdiff
* 交集:sinter
* 并集:sunion



### Hash(字典)

> 场景：适合对象数据的存储、电商购物车

有点特殊，但是常用：KV模式不变的，但V是一个键值对。Hash的字典相当于Java语言的HashMap，它是无序的字典。底层也是：数组+链表二维结构。

（1）hset/hget/hmset/hmget/hgetall/hdel

* Hset/hget：设置和获取

  ```
  127.0.0.1:6379> hset user id 11
  (integer) 1
  127.0.0.1:6379> hget user id
  "11"
  ```

* Hmset/hmget：设置多个和获取多个

  ```
  127.0.0.1:6379> hmset customer id 11 name 12 age 23
  OK
  127.0.0.1:6379> hmget customer id name  age
  1) "11"
  2) "12"
  3) "23"
  127.0.0.1:6379>
  ```

* Hgetall 是更方便

  ```
  127.0.0.1:6379> hgetall customer
  1) "id"
  2) "11"
  3) "name"
  4) "12"
  5) "age"
  6) "23"
  ```

* hdel删除

* hlen获取长度

（2）exists key 在key里面的某个值的key

（3）hkeys / hvals

```
127.0.0.1:6379> hkeys customer
1) "id"
2) "name"
3) "age"
127.0.0.1:6379> hvals customer
1) "11"
2) "12"
3) "23"
```

（4）hincrby / hincrbyfloat



（5）hsetnx（有了就插不了，没有就可以）



### Sorted set(ZSet)

> 场景：比如需要对某个数据全重进行排序。如：直播系统实时排序、粉丝列表等（zset存储粉丝列表，value值是粉丝用户ID，score是关注时间，然后可以对按照关注时间进行排序）

类似Java的SortedSet和HashMap的结合体，一方面是set保证内部value唯一性，另一方面可以给每个value赋予一个score，代表这个value的排序全重。底层：**跳跃列表的数据结构**

（1）zadd/zrang/(添加/获取)

zrange zset01 0 -1 withscores：是全部获取

zrange zset01 0 -1：只获取v1 ~ v5

```
127.0.0.1:6379> zadd zset01 10  v1 20 v2 30 v3 40 v4 50 v5 60 v6
(integer) 6
127.0.0.1:6379> zrange zset01 0 -1
1) "v1"
2) "v2"
3) "v3"
4) "v4"
5) "v5"
6) "v6"
127.0.0.1:6379> zrange zset01 0 -1 withscores
 1) "v1"
 2) "10"
 3) "v2"
 4) "20"
 5) "v3"
 6) "30"
 7) "v4"
 8) "40"
 9) "v5"
10) "50"
11) "v6"
12) "60"
```

（2）zrangebyscore key 开始 score 结束 score

获取60～90之间的

```
127.0.0.1:6379> zrangebyscore zset01 60 90
1) "v1"
2) "v2"
3) "v3"
4) "v4"
```

加`(`不包含的

```
127.0.0.1:6379> zrangebyscore zset01 (60 (90
1) "v2"
2) "v3"
```

limit获取部分的

```
127.0.0.1:6379> zrangebyscore zset01 60 90 limit 2 2
1) "v3"
2) "v4"
```

（3）zrem key 某score下对应的value值，作用是删除元素

```
127.0.0.1:6379> zrem zset01 v5
(integer) 1
```

（4）zcard / zcount key score区间/zrank key values值，作用是获取下标值/zscore key对应值

```
127.0.0.1:6379> zrange zset01 0 -1 withscores
 1) "v1"
 2) "60"
 3) "v2"
 4) "70"
 5) "v3"
 6) "80"
 7) "v4"
 8) "90"
 9) "v5"
10) "100"
# zcard统计全部
127.0.0.1:6379> zcard zset01
(integer) 4
# 统计60～80有多少人
127.0.0.1:6379> zcount zset01 60 80
(integer) 3
# zrank获取下标
127.0.0.1:6379> zrank zset01 v4
(integer) 3
127.0.0.1:6379> zscore zset01 v3
"80"
```

（5）zrevrank key values值，作用是逆序获得下标值

```
127.0.0.1:6379> zrevrank zset01 v4
(integer) 0
127.0.0.1:6379> zrevrank zset01 v3
(integer) 1
```

（6）zrevrange

```
127.0.0.1:6379> zrevrange zset01 0 -1
1) "v4"
2) "v3"
3) "v2"
4) "v1"
127.0.0.1:6379>
```

（7）zrevrangebyscore key 结果score开始score取反

```
127.0.0.1:6379> zrevrangebyscore zset01 90 60
1) "v4"
2) "v3"
3) "v2"
4) "v1"
127.0.0.1:6379>
```



**更多文章已被Github收录：https://github.com/niutongg/JavaLeague**

![4B544609-FA1C-4FE1-B99A-041A94838FAF.png](http://ww1.sinaimg.cn/large/006FuVcvgy1gulvyk74wcj61kc1mgdox02.jpg)

