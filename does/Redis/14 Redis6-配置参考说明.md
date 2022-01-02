### Redis-参数说明

1、Redis默认不是以守护进程的方式运行的，可以通过修改为守护进程运行方式，将no改为yes

```
daemonize yes
```

2、当Redis以守护进程方式运行时，Redis默认会把pid写入默认路径(/var/run/)，我们可以通过指定的

```
pidfile /var/run/redis6379.pid
```

3、指定Redis监听端口，默认端口为6379

```
port 6379
```

4、绑定的主机地址

```
bin 127.0.0.1
```

5、当客户端闲置多长时间后关闭连接，如果指定为0，表示关闭该功能

```
timeout 300
```

6、指定日志记录级别，Redis总共支持4个级别：debug、verbose、notice、warning（默认verbose）

```
loglevel varbose
```

7、日志记录方式，默认为标准输出，如果配置Redis为守护进程，而这里又配置为日志记录方式为标准输出，则日志将发送给/dev/null

```
logfile stdout
```

8、设置数据库的数量，默认数据库为0，可以使用弄select<dbid>命令多个条件配置

``` 
databases 16
```

9、指定多长时间内，有多少次更行操作，就将数据同步数据文件，可以多个条件配合

```
save <seconds> <changes>
# 默认三个条件,默认也是关闭的（这里是6.2版本）
# save 3600 1
# save 300 100
# save 60 10000
```

说明：3600秒(1小时)内有1个更改，300秒(5分钟)内100个更改，60秒内有10000个更改。

10、指定存储至本地数据库是否压缩数据，默认yes，Redis采用LZF压缩算法，如果为了节省CPU时间，可以选择关闭，但是会导致数据文件变大。

```
rdbcompression yes
```

11、指定本地数据库文件名，默认dump.rdb

```
dbfilename dump.rdb
```

12、指定本地数据库存放目录

```
dir ./
```

13、设置当本机为slave服务器时，设置master服务的IP地址及端口，在Redis启动时，它会自动从Master进行同步

```
slaveof <master> <masterport>
```

14、当master服务设置了密码保护时，slave服务连接master的密码

```
masterauth <master-password>
```

15、设置Redis密码，如果设置密码，客户端在连接Redis时需要通过AUTH<password> 命令提供密码，默认关闭。

```
requirepass foobared
```

16、设置同一时间最大客户连接数据，默认无限大，Redis可以同时打开的客户端连接数为Redis进程可以打开的最大文件描述符数，如果设置为maxclients 0，表示不作限制，打给你客户端连接数达到最大极限时，Redis会关闭新的连接并向客户端返回max number of clients reached错误

```
maxclients 128
```

17、指定Redis最大内存限制，Redis在启动时会把数据加载到内存中，达到最大内存后，Redis会先尝试清除已经过期或即将过期的Key，当此方法处理后，仍然达到最大内存设置，将无法再进行写入操作，但仍然可以进行读操作。Redis新的VM机制，会把Key存放在内存，Value会存放在swap区。

```
maxmemory <bytes>
```

18、指定是否在每次更新操作后进行日志记录，Redis默认情况下是异步的把数据写入磁盘，如果不开启，可能在断电时候导致一段时间内数据丢失，因为Redis本身同步数据文件是按上面save条件来同步的，所有的数据会在一段时间内只存在在内存中，默认no

19、指定更新文件名，默认为appendonly.aof

```
appendfilename appendonly.aof
```

20、指定更新日志条件，共三个可选值

```
no：表示等操作系统进行数据缓存同步磁盘
always：表示每次更新操作系统后，手动调用fsync将数据写到磁盘(慢，安全)
everysec：表示每秒同步一次
appendfsync everysec
```

21、指定是否启用虚拟机制，默认no，简单介绍一下，VM机制将数据分页存放，由于Redis将访问较少的页冷数据swap到磁盘上，访问多的页面由于磁盘自动换出到内存

```
vm-enabled no
```

22、虚拟机内存文件路径，默认值为/tmp/redis.swap，不可以多个redis实例共享

```
vm-swap-file /tmp/redis.swap
```

23、将所有大于vm-max-memory设置多小，所有索引数据都是内存存储的，也就是，当vm-max-memory设置为0的时候，其实是有value都存在于磁盘，默认值为0.

```
vm-max-memory 0
```

24、Redis swap文件分成了很多的page，一个对象可以保存在多个page上面，但是一个page不能被多个对象共享，vm-page-size是根据存储的数据大小来设定的，作者建议如果存储很多小对象，page大小最好设置为32或者64bytes，如果存储很大对象，则可以使用更大page，如果不确定就使用默认就好。

```
vm-page-size 32
```

25、设置swap文件中的page数量，由于页表是存在内存中，磁盘上要8个page将消耗1byte的内存

```
vm-page 134217728
```

26、设置访问swap文件的线程数量，最好不要超过机器的核数，如果设置为0，那么所有对swap文件的操作都是串行的，可能会造成比较长时间的延迟，默认值为4

```
vm-max-threads 4
```

27、设置在向客户端应答时，是否把较小的包含合并为一个包发送，默认开启

```
glueoutputbuf yes
```

28、指定在超过一定的数量或者最大的元素超过某一个临界值时，采用一种特殊的哈希算法

```
hash-max-zipmap-entries 64
hash-max-zipmap-value 512
```

29、指定是否激活重置哈希，默认开启

```
activerehashing yes
```

30、指定包含其他的配置文件，可以同一主机多个Redis实例之间使用同一份配置文件，而同时各个实例所有自己的特定配置文件

```
include /path/to/local.conf
```



