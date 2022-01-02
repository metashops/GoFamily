问题：现有50亿个电话号码，现在有10万电话号码，如何快速准确的判断这些电话号码是否已经存在？

布隆过滤器（Bloom Filter）是1970年由布隆的提出的。它实际上是一个很长的二进制数组和一系列Hash映射函数。用于判断一个元素是否在集合中。它的优点是空间效率和查询时间都远远超过一般的算法，缺点是有一定的误识别率和删除困难。

使用场景：解决缓存穿透/去重问题，比如推荐系统，每次推荐时都要去掉那些看过的内容的。当用户量很大，系统并发量很高时，如果记录都存在关系数据库，那么数据库抗压肯定有问题的，你会想到用缓存，但是历史记录非常龙大，那就非常浪费存储空间，历史记录是随着时间线性增长的，缓存也撑不了多久。此时布隆过滤器(Bloom Filter)登场。

布隆过滤器是什么？

你可以理解为布隆过滤器是一个可以去重的，但是不怎么精确的set结构，误判概率很小的。

命令：bf.add (添加1个元素) 和 bf.madd (批量添加元素)、bf.exists(查询1个元素是否存在)、bf.mexists批量查询。

**原理：**

布隆过滤器的原理是，当一个元素被加入集合时，通过K个散列函数将这个元素映射成一个位数组中的K个点，把它们置为1。检索时，我们只要看看这些点是不是都是1就（大约）知道集合中有没有它了：如果这些点有任何一个0，则被检元素一定不在；如果都是1，则被检元素很可能在。这就是布隆过滤器的基本思想。

**实现**

要使用BloomFilter，需要引入guava包

```pom
        <dependency>
            <groupId>com.google.guava</groupId>
            <artifactId>guava</artifactId>
            <version>19.0</version>
        </dependency>
```



```java
public class BloomFilter_ {
    private static int total = 1000000;
//    private static BloomFilter<Integer> bf = BloomFilter.create(Funnels.integerFunnel(), total);
    private static BloomFilter<Integer> bf = BloomFilter.create(Funnels.integerFunnel(), total, 0.001);

    public static void main(String[] args) {
        // 初始化1000000条数据到过滤器中
        for (int i = 0; i < total; i++) {
            bf.put(i);
        }

        // 匹配已在过滤器中的值，是否有匹配不上的
        for (int i = 0; i < total; i++) {
            if (!bf.mightContain(i)) {
                System.out.println("有坏人逃脱了~~~");
            }
        }

        // 匹配不在过滤器中的10000个值，有多少匹配出来
        int count = 0;
        for (int i = total; i < total + 10000; i++) {
            if (bf.mightContain(i)) {
                count++;
            }
        }
        System.out.println("误伤的数量：" + count);
    }
}

```

