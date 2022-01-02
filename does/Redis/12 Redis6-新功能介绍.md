### Redis6.0 的新功能

#### 1、ACL(访问控制列表)

Redis 的ACL是Access Control List（访问控制列表）的缩写。使用ACL功能可以对用户进行更全面的权限控制。

使用命令

（1）acl list 用于展现用户权限列表

```redis
127.0.0.1:6379> acl list
1) "user default on nopass * &* +@all"
```

（2）acl cat 用于查看数据结构的命令，比如：`acl string`会列出所有list操作的命令，当你忘记的时候很方便。

```redis
127.0.0.1:6379> acl cat list
 1) "rpush"
 2) "rpop"
 3) "sort"
 4) "lrange"
 ....省略
```



#### 2、IO多线程

Redis6.0 支持多线程，IO多线程其实指客户端交互部分的网络，IO交互处理模块多线程，而非执行命令多线程。注：(Redis6.0 执行命令仍然是单线程的)

Redis6.0 是支持多线程的，但只是用来处理网络数据的读写和协议解析，执行命令仍然是单线程的。

多线程的设计是要去控制Key、lua、事务、LPUSH/LPOP等并发问题。（默认是关闭的），需要配置文件配置。

```redis
######### THREADED I/O #############
# 开启多线程
io-threads-do-reads yes
# 4 线程数量
io-threads 4
```



#### 3、工具支持

之前老版本Redis搭建集群需要单独安装ruby环境的,Redis 5.0将redis-trip.rb的功能集成到redis-cli。而Redis 6.0 官方redis-benchmark工具支持cluster模式，通过多线程的方式对多个分片进行压测。

命令查看：下面--cluster Enable cluster mode

```redis
dis/6.2.5/bin> redis-benchmark --help
-h <hostname>      Server hostname (default 127.0.0.1)
 -p <port>          Server port (default 6379)
 -s <socket>        Server socket (overrides host and port)
 -a <password>      Password for Redis Auth
 --user <username>  Used to send ACL style 'AUTH username pass'. Needs -a.
 -c <clients>       Number of parallel connections (default 50)
 -n <requests>      Total number of requests (default 100000)
 -d <size>          Data size of SET/GET value in bytes (default 3)
 --dbnum <db>       SELECT the specified db number (default 0)
 --threads <num>    Enable multi-thread mode.
 --cluster          Enable cluster mode.
 --enable-tracking  Send CLIENT TRACKING on before starting benchmark.
 -k <boolean>       1=keep alive 0=reconnect (default 1)
 -r <keyspacelen>   Use random keys for SET/GET/INCR, random values for SADD,
                    random members and scores for ZADD.
```



#### 4、其他功能

RESP3新的redis通信协议、Client side caching客户端缓存、Proxy集群代理模式、Modules API。