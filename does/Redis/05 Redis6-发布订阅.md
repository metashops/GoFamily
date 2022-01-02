### Redis发布订阅（了解即可实际开发不用）

是进程间的一种消息通信模式：发送者(pub)发送消息，订阅者(sub)接收消息/

如何使用？

（1）可以一次性订阅多个

```
# 假设订阅三个频道如下
127.0.0.1:6379> subscribe c1 c2 c3
```

（2）发布

```
# 假设给c1发布如下，那么C1就会收到
127.0.0.1:6379> publish c1 hello-redis
(integer) 1
```





**更多文章已被Github收录：https://github.com/niutongg/JavaLeague**

![4B544609-FA1C-4FE1-B99A-041A94838FAF.png](http://ww1.sinaimg.cn/large/006FuVcvgy1gulvyk74wcj61kc1mgdox02.jpg)