### Bitmaps

Bitmaps本身不是实际数据类型，它只是对字符串的位图进行操作。（适用于活动量多的）

### HyperLogLog

在实际中，我们经常会遇到与统计相关的功能需求，比如统计网站PV(pageView页面访问量)，可以使用Redis的incr和incrby就可以搞定的。但是redis提供了HyperLogLog用来做基数统计的算法，当我们在输入元素的数量或者体积很大的时候，计算基数所需要的空间是固定的。每个HyperLogLog键只需要12KB内存，就可以计算按近2^64个不同的基数。(去重功能)

**命令**：

* Pfadd/pfcount（添加/统计）

```redis
127.0.0.1:6379> pfadd hyper1 "hadoop"
(integer) 1
127.0.0.1:6379> pfadd hyper1 "spark"
(integer) 1
127.0.0.1:6379> pfadd hyper1 "kafka"
(integer) 1
127.0.0.1:6379> pfadd hyper1 "kafka"
(integer) 0
127.0.0.1:6379> pfcount hyper1
(integer) 3
```

* Pfmerge（合并），如下hyper1有3个元素和hyper2有2个元素，合并后hyper3就有5个了

```redis
127.0.0.1:6379> pfadd hyper1 "hadoop"
(integer) 1
127.0.0.1:6379> pfadd hyper1 "spark"
(integer) 1
127.0.0.1:6379> pfadd hyper1 "kafka"
(integer) 1
127.0.0.1:6379> pfadd hyper2 "java"
(integer) 1
127.0.0.1:6379> pfadd hyper2 "python"
(integer) 1
127.0.0.1:6379> pfmerge hyper3 hyper1 hyper2
OK
127.0.0.1:6379> pfcount hyper3
(integer) 5
```

总结：用于解决基数问题



### Geospatial

Redis 3.2中增加了对GEO类型的支持，GEO(Geograhic)，地理信息的缩写。该类型就是2维坐标，在地图就是经纬度。Redis基于该类型，提供了经纬度设置、查询、范围查询、距离查询、经纬度hash等常见操作。

**命令**：

* GEOADD

```redis
127.0.0.1:6379>GEOADD Sicily 13.361389 38.115556 "Palermo" 15.087269 37.502669 "Catania"
```

* geodist：km单位或者m都可以

```redis
127.0.0.1:6379>geodist Sicily Palermo Catania km
```

* GEORADIUS：以给定的经纬度为中心，找出某一半经内的元素