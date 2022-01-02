### 哨兵(sentinel)模式

第一步：首先创建：touch sentinel.conf文件

第二步：编辑哨兵文件

vim sentinel.conf，内容如下：

```redis
# host6379是被监控主机名称(自己取名) IP 端口 1
sentinel monitor host6379 127.0.0.1 6379 1
```

注：上面的1，表示主机挂掉后slave投票，谁票多谁就成为主机。

第三步：启动哨兵

进入redis安装bin目录执行该命令：

```redis
redis-sentinel /usr/local/etc/sentinel.conf
```

结果如下：

```redis
7513:X 20 Sep 2021 09:45:13.410 # oO0OoOo Redis is starting oO0OoO0Oo
7513:X 20 Sep 2021 09:45:13.410 # Redis version=6.2.5, bits=64, commit=000000, modified=0, pid=7915, just started
7513:X 20 Sep 2021 09:45:13.410 # Configuration loaded
7513:X 20 Sep 2021 09:45:13.411 * Increased maximum number of open files to 10032 (it was originally set to 256).
7513:X 20 Sep 2021 09:45:13.411 * monotonic clock: POSIX clock_gettime
                _._
           _.-``__ ''-._
      _.-``    `.  `_.  ''-._           Redis 6.2.5 (00000000/0) 64 bit
  .-`` .-```.  ```\/    _.,_ ''-._
 (    '      ,       .-`  | `,    )     Running in sentinel mode
 |`-._`-...-` __...-.``-._|'` _.-'|     Port: 26379
 |    `-._   `._    /     _.-'    |     PID: 7622
  `-._    `-._  `-./  _.-'    _.-'
 |`-._`-._    `-.__.-'    _.-'_.-'|
 |    `-._`-._        _.-'_.-'    |           https://redis.io
  `-._    `-._`-.__.-'_.-'    _.-'
 |`-._`-._    `-.__.-'    _.-'_.-'|
 |    `-._`-._        _.-'_.-'    |
  `-._    `-._`-.__.-'_.-'    _.-'
      `-._    `-.__.-'    _.-'
          `-._        _.-'
              `-.__.-'

7513:X 20 Sep 2021 09:45:13.413 # Sentinel ID is 2b8f846df1ec353hggjk25c34hg6212d4hgfghj82e7
# 监控6379主机
7513:X 20 Sep 2021 09:45:13.413 # +monitor master host6379 127.0.0.1 6379 quorum 1
# 巡逻6380 和 6381从机，一旦上面6379挂掉了，就投票选出新的master
7513:X 20 Sep 2021 09:45:13.420 * +slave slave 127.0.0.1:6380 127.0.0.1 6380 @ host6379 127.0.0.1 6379
7513:X 20 Sep 2021 09:45:13.421 * +slave slave 127.0.0.1:6381 127.0.0.1 6381 @ host6379 127.0.0.1 6379

```

第五步：故意让6379挂掉，执行shutdown

### 哨兵缺点

由于所有的写操作先在master上操作，然后同步更新到slave上，所以从master同步到Slave有一定的延迟，当系统很繁忙时，延迟问题更加严重。

