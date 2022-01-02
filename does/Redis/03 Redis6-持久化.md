## 持久化

Redis支持RDB(Redis DataBase)和AOF(Append Only File)两种持久化机制，持久化功能有效地避免因进程退出造成的数据丢失问题，当下次重启时利用之前持久化的文件即可实现数据恢复。（Redis的数据是存放在内存中的，假如**突然宕机**，数据就会丢失，所以我们需要一种机制来确保Redis即使在宕机时数据不会丢失，此时**Redis持久化机制**登场）。

建议看[官方文档](https://redis.io/topics/persistence)说明。

## RDB

官方介绍：在指定的时间间隔内将内存的数据集快照写入磁盘，也就是行话讲的**Snapshot快照**，它恢复时是将快照文件直接读到内存里。

### RDB原理？

在保存RDB文件时，Redis会单独创建(fork)的一个子进程来进行持久化，此时父进程不需要做其他IO操作。fork是指redis通过创建子进程来进行RDB操作，cow指的是**copy on write**，子进程创建后，父子进程共享数据段，父进程继续提供读写服务，写脏的页面数据会逐渐和子进程分离开来。

RDB的缺点：最后一次持久化后的数据可能丢失，所以一般我们都是用RDB做镜像全量持久化，AOF做增量持久化，进行这样一个配合使用。

fork：作用就是复制一个与当前进程一样的进程。新的进程的所有数据都是和原进程一致的，但是是一个全新的进程，并作为原进程的子进程。

### 如何触发RDB快照？

（1）配置文件(默认使用出厂配置的)

```redis
################################ SNAPSHOTTING  ################################
# Unless specified otherwise, by default Redis will save the DB:
#   * After 3600 seconds (an hour) if at least 1 key changed
#   * After 300 seconds (5 minutes) if at least 100 keys changed
#   * After 60 seconds if at least 10000 keys changed
#
# You can set these explicitly by uncommenting the three following lines.
# 广商默认一下三种（可以修改）
# 1、1小时内，key改动一次，就进行触发快照机制
# 2、5分钟内，key改动100次，则触发快照机制
# 3、1分钟内，key改动10000次，则触发RDB快照机制
save 3600 1
save 300 100
save 60 10000
```

设置完值之后，会产生`dump.rdb`文件，值得注意的是：产生的dump.rdb文件，一定一定备份到另一台备份机器上，否则你宕机了，文件就丢失，还有就是第一次触发完RDB后，第二次触出发RDB时，会将第一次产生的dump.rdb覆盖掉。

（2）如何恢复呢？

从备份机器将dump.rdb移动到Redis安装目录，并且启动服务器即可恢复。(这样就可以从新读会内存中)

（3）命令save或者bgsave都会迅速生产dump.rdb文件备考

* 使用save时，它只管保存，其他的不管，全部阻塞
* 使用bgsave时，Redis会在后台异步进行快照(也就是说，进行快照同时还可以响应客户端请求)
* save 和 bgsave区别？save命令会阻塞服务器，而 bgsave命令不会

### RDB优缺点？

优点：适合大规模的数据恢复；对数据完整性和一致性要求不高

缺点：会丢失最后一次快照后的所有修改；Fork时，内存中的数据被克隆一份，2倍的膨胀性能。



## AOF(Append Only File)

AOF持久化是通过保存Redis服务器所执行的写命令来记录数据库状态，只允许追加文件但不可以修改文件。redis重启后根据保存的记录指令执行完成数据的恢复。

配置文件

```redis
############################## APPEND ONLY MODE ###############################
# ......

# AOF 和 RDB 的持久性可以同时启用，无需问题
# AOF and RDB persistence can be enabled at the same time without problems.
# If the AOF is enabled on startup Redis will load the AOF, that is the file
# with the better durability guarantees.
#
# Please check https://redis.io/topics/persistence for more information.

# 1、默认AOF持久化是关闭的
# appendonly no
# 开启
appendonly yes

# 2、默认AOF的名字是appendonly.aof
# The name of the append only file (default: "appendonly.aof")
appendfilename "appendonly.aof"

# The fsync() call tells the Operating System to actually write data on disk
# instead of waiting for more data in the output buffer. Some OS will really flush
# data on disk, some other OS will just try to do it ASAP.
#
# Redis supports three different modes:
#
# no: don't fsync, just let the OS flush the data when it wants. Faster.
# always: fsync after every write to the append only log. Slow, Safest.
# everysec: fsync only one time every second. Compromise.
#
# 默认每秒
# The default is "everysec", as that's usually the right compromise between
# speed and data safety. It's up to you to understand if you can relax this to
# "no" that will let the operating system flush the output buffer when
# it wants, for better performances (but if you can live with the idea of
# some data loss consider the default persistence mode that's snapshotting),
# or on the contrary, use "always" that's very slow but a bit safer than
# everysec.
#
# More details please check the following article:
# http://antirez.com/post/redis-persistence-demystified.html
#
# If unsure, use "everysec".

# 3、appendfsync
# 3.1 always：同步持久化每次发生数据变更会被立即立即到磁盘，这种方式性能较差但保证数据的完整性
# 3.2 everysec：出厂默认推荐，异步操作，每秒记录，如果一秒内宕机，有数据丢失
# 3.3 no
# appendfsync always
appendfsync everysec
# appendfsync no

# When the AOF fsync policy is set to always or everysec, and a background
# saving process (a background save or AOF log background rewriting) is
# performing a lot of I/O against the disk, in some Linux configurations
# Redis may block too long on the fsync() call. Note that there is no fix for
# this currently, as even performing fsync in a different thread will block
# our synchronous write(2) call.
#
# In order to mitigate this problem it's possible to use the following option
# that will prevent fsync() from being called in the main process while a
# BGSAVE or BGREWRITEAOF is in progress.
#
# This means that while another child is saving, the durability of Redis is
# the same as "appendfsync none". In practical terms, this means that it is
# possible to lose up to 30 seconds of log in the worst scenario (with the
# default Linux settings).
#
# If you have latency problems turn this to "yes". Otherwise leave it as
# "no" that is the safest pick from the point of view of durability.

no-appendfsync-on-rewrite no

# Automatic rewrite of the append only file.
# Redis is able to automatically rewrite the log file implicitly calling
# BGREWRITEAOF when the AOF log size grows by the specified percentage.
#
# This is how it works: Redis remembers the size of the AOF file after the
# latest rewrite (if no rewrite has happened since the restart, the size of
# the AOF at startup is used).
#
# This base size is compared to the current size. If the current size is
# bigger than the specified percentage, the rewrite is triggered. Also
# you need to specify a minimal size for the AOF file to be rewritten, this
# is useful to avoid rewriting the AOF file even if the percentage increase
# is reached but it is still pretty small.
#
# Specify a percentage of zero in order to disable the automatic AOF
# rewrite feature.

auto-aof-rewrite-percentage 100
auto-aof-rewrite-min-size 64mb

# An AOF file may be found to be truncated at the end during the Redis
...
```

开启配置文件后，就会生成appendonly.aof，当你宕机时，重启redis服务器，会重新执行appendonly.aof里面所有记录恢复数据。

### AOF 和 RDB共存，是如何加载的？

首先加载appendonly.aof文件，如果appendonly.aof出现错误/异常，那么我们可以执行：`redis-check-aof --fix appendonly.aof `即可恢复。

### AOF 的配置策略是什么？

```redis
# 默认每秒
# The default is "everysec", as that's usually the right compromise between
# speed and data safety. It's up to you to understand if you can relax this to
# "no" that will let the operating system flush the output buffer when
# it wants, for better performances (but if you can live with the idea of
# some data loss consider the default persistence mode that's snapshotting),
# or on the contrary, use "always" that's very slow but a bit safer than
# everysec.
#
# More details please check the following article:
# http://antirez.com/post/redis-persistence-demystified.html
#
# If unsure, use "everysec".

# 3、appendfsync
# 3.1 always：同步持久化每次发生数据变更会被立即到磁盘，这种方式性能较差但保证数据的完整性
# 3.2 everysec：出厂默认推荐，异步操作，每秒记录，如果一秒内宕机，有数据丢失
# 3.3 no
# appendfsync always
appendfsync everysec
# appendfsync no
```

### ReWrite

（1）什么是ReWrite？

aof采用文件追加方式，文件会越来越大为了避免出现此情况，新增一种机制。当AOF文件的大小超过所设定的阈值，Redis就会启动AOF文件的内容压缩，只保留可以恢复数据的最小指令集，可以使用命令`bgrewriteaof`。

（2）重写原理

AOF 文件持续增长而过大时，还是会fork出一条新进程来将文件重写，然后遍历新进程的内存中数据，每一条记录有一条的set语句。

（3）触发机制

Redis会记录上一次重写时的AOF大小，默认配置时当AOF文件大小是上次rewrite后大小的一倍且文件大于默认64M时触发。

```redis
############################## APPEND ONLY MODE ###############################
...
# Specify a percentage of zero in order to disable the automatic AOF
# rewrite feature.

# 默认64mb，大型互联网是根本不够用，会很大的
auto-aof-rewrite-percentage 100
auto-aof-rewrite-min-size 64mb
```

### AOF优缺点

优势：

* 每秒同步：appendfsync always 同步持久化，每次发生数据变更会被立即记录到磁盘，性能差但数据完整性好
* 每修改同步：appendfsync everysec 异步操作，每秒记录，如果1秒内宕机，有数据丢失
* 不同步：appendfsync no 从不同步

劣势：

* 相同数据集的数据而言，AOF文件要远大于RDB文件，恢复速度慢于RDB
* AOF运行效率慢于RDB，每秒同步策略效率较好，不同步效率和RDB相同



### RDB 和 AOF 总结

（1）RDB持久化方式，能够在指定时间间隔内将数据进行快照

（2）AOF持久化方式，记录每次对Redis服务器写的操作，当重启服务器时会重写执行这些命令来恢复原始的数据。

（3）同步开启两种方式：同时开启时，Redis重启时候会优先加载AOF文件来恢复原始数据(先加载的原因：通常情况下AOF文件保存的数据比RDB要更加完整性)。

（4）可以只使用AOF？不建议，因为RDB更适合用于备份数据库快速重启，而且不会有AOF可能潜在bug，

### 常见缓存淘汰策略

常见缓存淘汰算法：FIFO先进先出、LRU最近最少使用、LFU最近使用频率最低

几种策略：



**更多文章已被Github收录：https://github.com/niutongg/JavaLeague**