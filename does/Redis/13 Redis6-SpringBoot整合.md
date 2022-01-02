### Redis与SpringBoot整合

第一步：新建一个SpringBoot项目

第二步：在pom.xml文件中引入redis依赖即可

```java
<!--redis依赖-->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-data-redis</artifactId>
        </dependency>
        <!--spring2.x 集成redis所需要连接池-->
        <dependency>
            <groupId>org.apache.commons</groupId>
            <artifactId>commons-pool2</artifactId>
            <version>2.7.0</version>
        </dependency>
```

第三步：在application.properties配置文件配置redis，根据实际配置即可

```java
mycar.beand=BYD
mycar.price=100000
# redis服务器地址
spring.redis.host=127.0.0.1
# redis服务器端口
spring.redis.port=6379
# redis数据库索引（默认使用0库）
spring.redis.database=0
# redis连接超时时间（毫秒）
spring.redis.timeout=18000000
# redis连接池最大连接数
spring.redis.lettuce.pool.max-active=20
# redis最大阻塞等待时间（负数表示没有限制）
spring.redis.lettuce.pool.max-wait=-1
# redis连接池中的最大空闲连接
spring.redis.lettuce.pool.max-idle=5
# redis连接池最小空闲连接
spring.redis.lettuce.pool.min-idle=0
```

第四步：添加redis配置类，这样在Springboot启动加载可以使用，我们在，在config目录下建一个配置类。

```java
//EnableCaching开启，Configuration表示这个类是配置类
@EnableCaching
@Configuration
public class RedisConfig extends CachingConfigurerSupport {
    @Bean
    public RedisTemplate<String, Object> redisTemplate(RedisConnectionFactory factory) {
        RedisTemplate<String, Object> Template = new RedisTemplate<>();
        StringRedisSerializer redisSerializer = new StringRedisSerializer();
        Jackson2JsonRedisSerializer jackson2JsonRedisSerializer = new Jackson2JsonRedisSerializer(Object.class);
        ObjectMapper om = new ObjectMapper();
        om.setVisibility(PropertyAccessor.ALL, JsonAutoDetect.Visibility.ANY);
        om.enableDefaultTyping(ObjectMapper.DefaultTyping.NON_FINAL);
        jackson2JsonRedisSerializer.setObjectMapper(om);
        Template.setConnectionFactory(factory);
        //Key序列化
        Template.setKeySerializer(redisSerializer);
        //value序列化
        Template.setKeySerializer(redisSerializer);
        return Template;
    }
    @Bean
    public CacheManager cacheManager(RedisConnectionFactory factory) {
        RedisSerializer<String> redisSerializer = new StringRedisSerializer();
        Jackson2JsonRedisSerializer jackson2JsonRedisSerializer = new Jackson2JsonRedisSerializer(Object.class);
        //解决查询缓存转换异常问题
        ObjectMapper om = new ObjectMapper();
        om.setVisibility(PropertyAccessor.ALL, JsonAutoDetect.Visibility.ANY);
        om.enableDefaultTyping(ObjectMapper.DefaultTyping.NON_FINAL);
        jackson2JsonRedisSerializer.setObjectMapper(om);
        //设置序列化（解决乱码问题），过期时间600秒
        RedisCacheConfiguration config = RedisCacheConfiguration.defaultCacheConfig()
                .entryTtl(Duration.ofSeconds(600))
                .serializeKeysWith(RedisSerializationContext.SerializationPair.fromSerializer(redisSerializer))
                .serializeValuesWith(RedisSerializationContext.SerializationPair.fromSerializer(jackson2JsonRedisSerializer))
                .disableCachingNullValues();
        RedisCacheManager cacheManager = RedisCacheManager
                .builder(factory)
                .cacheDefaults(config)
                .build();
        return cacheManager();
    }
}
```

第五步：Controller层添加测试方式

```java
@RestController
@RequestMapping("/redisTest")
public class RedisControllerTest {
    //注入Redis，然后就可以调用啦
    @Autowired
    private RedisTemplate redisTemplate;
    @GetMapping
    public String testRedis(){
        //
        redisTemplate.opsForValue().set("id","1");
        String idBy = (String) redisTemplate.opsForValue().get("id");
        return idBy;
    }
}
```

第六步：启动