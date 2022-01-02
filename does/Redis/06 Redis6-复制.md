### 模拟复制(一主多从)

第一步：我们模拟使用三台机子就可以了，先将redis.conf复制出来三份，如下：

![85098D3F-8297-49AE-9CAA-B928746C7363.png](http://ww1.sinaimg.cn/large/006FuVcvgy1gum4vlqo5pj60jo062js602.jpg)

第二步：配置文件

（1）Vim redis6379.conf

```redis
################################# GENERAL #####################################
daemonize yes
pidfile /var/run/redis6379.pid

# 日志文件
# Specify the log file name. Also the empty string can be used to force
# Redis to log on the standard output. Note that if you use standard
# output for logging but daemonize, logs will be sent to /dev/null
logfile "6379.log"

################################ SNAPSHOTTING  ################################
# The filename where to dump the DB
dbfilename dump6379.rdb
```

（2）Vim redis6380.conf

```redis
################################## NETWORK #####################################
# Accept connections on the specified port, default is 6379 (IANA #815344).
# If port 0 is specified Redis will not listen on a TCP socket.
# 修改网络端口
port 6380
################################# GENERAL #####################################
# 开启守护进程
# When Redis is supervised by upstart or systemd, this parameter has no impact.
daemonize yes
pidfile /var/run/redis6380.pid
logfile "6380.log"
################################ SNAPSHOTTING  ################################
# The filename where to dump the DB
dbfilename dump6380.rdb
```

（3）Vim redis6381.conf

```
################################## NETWORK #####################################
# Accept connections on the specified port, default is 6379 (IANA #815344).
# If port 0 is specified Redis will not listen on a TCP socket.
# 修改网络端口
port 6381
################################# GENERAL #####################################
# 开启守护进程
# When Redis is supervised by upstart or systemd, this parameter has no impact.
daemonize yes
pidfile /var/run/redis6380.pid
logfile "6381.log"
################################ SNAPSHOTTING  ################################
# The filename where to dump the DB
dbfilename dump6381.rdb
```

第三步：1主2从

（1）首先启动三台机子，并且查看每台都是master：

```redis
redis-server /usr/local/etc/redis6379.conf
redis-cli -p 6379
```

![78E01BB9-3C24-4955-9503-06653C0D5F89.png](http://ww1.sinaimg.cn/large/006FuVcvgy1guma4l7fr6j611y0ki0vc02.jpg)

（2）将6379设置为master的，6380和6381作为从机。只需要在另外两台机子执行该命令即可：

```redis
slaveof 127.0.0.1 6379
```

![CB1D066E-7978-4E88-824F-E7DF5A04B24F.png](http://ww1.sinaimg.cn/large/006FuVcvgy1gumakav2tsj618y0mk77v02.jpg)

此时，你在6379上做任何动作，其他从机都能get到。**注意：只有主机才能进行读写，从机只能读不能写。**由于复制只能主机到从机，对于从机的任何修改主机都是无法感知的。如果你不使用默认从机只读模式，修改从机会造成主从数据不一致，

（3）假设（三种主从复制重点，后面还有哨兵模式）

* 第一种：将6379这台主机shutdown了，会怎么样？这时从机待命，等主机回来。
* 第二种：6380从机挂掉了，再次回来，会是什么角色？回来是master角色。如何重写连接回来呢？仍然是使用：`slaveof 127.0.0.1 6379`即可。
* 第三种：主机宕机时，那么就让其他从机转换为主机，使用命令：`slaveof no one`即可。

### (一主多从)总结：

一主一从结构是最简单的，我们来看看**一主多从**结构，对于读占比较大的场景，可以把读命令发送到从机来分担主机的压力。如果对于写并发量较高的场景，多个从机会导致主机写命令的多次发送从而度过消耗网络带宽，同时也加重了主机的负载影响服务稳定性。这时可以使用树状主从结构，也就是从节点不但可以复制主节点 数据，同时可以作为其他从节点的主节点继续向下层复制。这样的目的是有效降低主节点负载和需要传送给从节点的数据量。

**下一篇：哨兵模式**



**更多文章已被Github收录：https://github.com/niutongg/JavaLeague**

![4B544609-FA1C-4FE1-B99A-041A94838FAF.png](http://ww1.sinaimg.cn/large/006FuVcvgy1gulvyk74wcj61kc1mgdox02.jpg)