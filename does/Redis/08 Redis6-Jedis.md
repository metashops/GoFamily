只要把前面的知识点掌握了，以下就非常顺手了

### 1、打开IDEA新建一个空的项目就可以，只是练习使用Redis

![9DD3B9AA-50F2-4C99-9B73-574F532BE05F.png](http://ww1.sinaimg.cn/large/006FuVcvgy1gumyd6p5jwj618e11atbg02.jpg)

然后下一步，创建完项目后，连接即可使用。

> 注：创建空项目，如果运行报错，记得把JDK版本改为你自己的版本，有时候默认就是JDK1.5版本的。

改JDK版本有三个地方：

一：

![6315CFCD-63B4-4AF0-A3C2-4E461925393D.png](http://ww1.sinaimg.cn/large/006FuVcvgy1gumyk6a257j612m0g8n0502.jpg)

二：

![2D615382-7291-45C2-BEE9-62C796764EF2.png](http://ww1.sinaimg.cn/large/006FuVcvgy1gumyllio9ij60l60guabx02.jpg)

![5D143257-1F2B-4257-B6F9-2DFA76EF5C99.png](http://ww1.sinaimg.cn/large/006FuVcvgy1gumyn1nysuj61di124jxe02.jpg)

三：

![AF4ACAE5-306B-4FB9-83C1-2CC66EBB6AD7.png](http://ww1.sinaimg.cn/large/006FuVcvgy1gumypfudkij60va0pawi402.jpg)

![7B159A69-FEA0-420C-B995-86C52F2600EE.png](http://ww1.sinaimg.cn/large/006FuVcvgy1gumyqseinjj61do12e79002.jpg)

### 2、连接

![00E78B6E-4A19-4F06-8C56-36D1EB206E2A.png](http://ww1.sinaimg.cn/large/006FuVcvgy1gumyftxhybj61nw0iwwhw02.jpg)

### 3、开始我们的5种数据结构

### 4、Redis-事务

（1）日常

```java
public class RedisTestTX {
    public static void main(String[] args) {
        Jedis jedis = new Jedis("127.0.0.1",6379);
        //
        Transaction transaction = jedis.multi();
        transaction.set("k4","v4");
        transaction.set("k5","v5");
        // 执行提交
        transaction.exec();
    }
}
```



（2）加锁

现在我们先存钱

```java
127.0.0.1:6379> set balance 100
OK
127.0.0.1:6379> set debt 0
OK
```

执行Java代码：

```java
public class RedisTestTX {
    public boolean transMethod() {
        Jedis jedis = new Jedis("127.0.0.1",6379);
        int balance; //可用余额
        int debt; //欠额
        int amtToSubtract = 10; //模拟实刷额度
        jedis.watch("balance");
        balance = Integer.parseInt(jedis.get("balance"));
        if (balance < amtToSubtract) {
            jedis.unwatch(); // 从watch变成unwatch
            System.out.println("modify");
            return false;
        } else {
            System.out.println("------transaction------");
            Transaction transaction = jedis.multi(); // 开启事务
            transaction.decrBy("balance",amtToSubtract); //假如你刷10元，就减10元
            transaction.incrBy("debt",amtToSubtract); //增加
            transaction.exec(); // 提交
            balance = Integer.parseInt(jedis.get("balance"));
            debt = Integer.parseInt(jedis.get("debt"));
            System.out.println("balance:" + balance);
            System.out.println("debt:" + debt);
            return true;
        }
    }
    public static void main(String[] args) {
        RedisTestTX test = new RedisTestTX();
        boolean retValue = test.transMethod();
        System.out.println("main retValue" + retValue);
    }
}
```

运行后，我们去查询

```redis
127.0.0.1:6379> get balance
"90"
127.0.0.1:6379> get debt
"10"
```

总结：

（1）watch命令是标记一个键（监控这个值是否被别人修改过）

（2）在提交事务前，如果该键别人修改过，那么事务就失败（在程序中会重新再次尝试）

（3）上面案例，我们首先标记balance，然后检查余额是否足够，不足就取消，并不做加减操作

（4）如果在此期间键balance被其他人修改，那正在提交事务（exec）时会报错（在程序中也是在此会尝试，直到成功）

### 5、Redis-主从复制

了解即可，一般不会使用代码来进行配置的。

```java
public class RedisTestMS {
    public static void main(String[] args) {
        Jedis jedisM = new Jedis("127.0.0.1", 6379);
        Jedis jedisS1 = new Jedis("127.0.0.1", 6380);
        Jedis jedisS2 = new Jedis("127.0.0.1", 6381);
        jedisS1.slaveof("127.0.0.1",6379); //设置6379为主节点
        jedisS2.slaveof("127.0.0.1",6379);
        jedisM.set("class", "hello redis");
        String set = jedisS1.get("class");
        System.out.println("result:" + set);
    }
}

```



### 6、Redis-Jedis-JedisPool

Redis连接池是什么？

首先Redis也是一种数据库，它基于C/S模式，因此如果需要使用必须建立连接，稍微熟悉网络的人应该都清楚地知道为什么需要建立连接，C/S模式本身就是一种远程通信的交互模式，因此Redis服务器可以单独作为一个数据库服务器来独立存在。假设Redis服务器与客户端分处在异地，虽然基于内存的Redis数据库有着超高的性能，但是底层的网络通信却占用了一次数据请求的大量时间，因为每次数据交互都需要先建立连接，假设一次数据交互总共用时30ms，超高性能的Redis数据库处理数据所花的时间可能不到1ms，也即是说前期的连接占用了29ms，连接池则可以实现在客户端建立多个链接并且不释放，当需要使用连接的时候通过一定的算法获取已经建立的连接，使用完了以后则还给连接池，这就免去了数据库连接所占用的时间。
如何使用连接池？

（1）新建一个类，进行封装

```java
public class RedisJedisPool {
    // 双向检索
    private static volatile JedisPool jedisPool = null;
    private RedisJedisPool(){}
    public static JedisPool getJedisPoolInstance() {
        if (null == jedisPool) {
            synchronized (RedisJedisPool.class) {
                if (null == jedisPool) {
                    //Jedis连接池的配置
                    JedisPoolConfig poolConfig = new JedisPoolConfig();
                    poolConfig.setMaxTotal(1000);
                    poolConfig.setMaxIdle(32);//空闲（32空闲就要换）
                    poolConfig.setMaxWaitMillis(100*10000);//最大的等待
                    poolConfig.setTestOnBorrow(true);//获得一个Jedis实例的时候是否检查连接可用性(ping()->true)
                    jedisPool = new JedisPool(poolConfig, "127.0.0.1", 6379);
                }
            }
        }
        return jedisPool;
    }
    public static void release(JedisPool jedisPool, Jedis jedis) {
        if (null != jedis) {
            jedisPool.returnResource(jedis);
        }
    }
}

```

（2）测试使用

```java
public class TestPool {
    public static void main(String[] args) {
        JedisPool jedisPool1 = RedisJedisPool.getJedisPoolInstance();
        JedisPool jedisPool2 = RedisJedisPool.getJedisPoolInstance();
        System.out.println(jedisPool1 == jedisPool2);
        Jedis jedis = null;
        try {
            jedis = jedisPool1.getResource();//从池拿出来
            jedis.set("aa","11");
        } catch (Exception e) {
            e.printStackTrace();
        } finally {
            RedisJedisPool.release(jedisPool1,jedis);
        }
    }
}

```

