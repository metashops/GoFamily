# GoFamily介绍：
想和大家一起学习Go语言。

以下是我平时学习的笔记，希望能帮助到正在学习Go之旅。

一起拥抱更多开源，希望大家多多支持。（别忘记**Star**哟）

<p align="center">
  <a href="#公众号"><img src="https://img.shields.io/badge/%E5%85%AC%E4%BC%97%E5%8F%B7-metashops-lightgrey.svg" alt="公众号"></a>
  <a href="https://blog.csdn.net/realize_dream?type=blog"><img src="https://img.shields.io/badge/CSDN-博客-critical"></a>
</p>



[Go学习路线，脑图在线编辑地址](https://www.processon.com/view/link/61d69f7f07912973ef0fa690) 



<details> <summary>Go学习路线，展开脑图</summary> <pre><img src="https://github.com/metashops/GoFamily/blob/main/golang-developer-roadmap.png"/></pre> </details>



## Table of Contents(Ctrl+F)

## 计算机基础

八股文笔记在线脑图地址：[计算机网络](https://www.processon.com/view/link/620e0e86079129799619067a) 

计算机网络

1. [浏览器发起http请求背后的流程](https://github.com/metashops/GoFamily/blob/main/network/1-%E6%B5%8F%E8%A7%88%E5%99%A8%E5%8F%91%E8%B5%B7HTTP%E8%AF%B7%E6%B1%82%E8%83%8C%E5%90%8E%E7%9A%84%E5%8E%9F%E7%90%86.md)  
2. [全面剖析http细节](https://github.com/metashops/GoFamily/blob/main/network/2-%E5%85%A8%E9%9D%A2%E5%89%96%E6%9E%90HTTP%E7%BB%86%E8%8A%82.md)  

## Go 基础知识

基础数据结构

* [Go 语言Map底层原理](https://blog.csdn.net/realize_dream/article/details/121999836?spm=1001.2014.3001.5501) 
* [Go 语言切片](https://blog.csdn.net/realize_dream/article/details/121952796?spm=1001.2014.3001.5501) 

* [Go 语言反射](https://blog.csdn.net/realize_dream/article/details/121940315?spm=1001.2014.3001.5501) 

并发控制

* [Mutex 如何解决资源并发访问问题？](https://mp.weixin.qq.com/s?__biz=MzUyODgxNzM0Nw==&amp;mid=2247484323&amp;idx=1&amp;sn=264afadcf4ca1092d891f076c5e200e9&amp;chksm=fa6bce0fcd1c4719c1b92c4827b03488b982d424f43a9508c45eca86499e9564ecfbd5694bce&token=1599394794&lang=zh_CN#rd) 
* [Go语言Mutex 源码分析](https://mp.weixin.qq.com/s?__biz=MzUyODgxNzM0Nw==&amp;mid=2247484304&amp;idx=1&amp;sn=63f463b4d143377d1b1540e99f8cc22d&amp;chksm=fa6bce3ccd1c472a8eedf4a34a43891a38729e2c50460f01aab5826001149003fc1ab91bb1d4&token=1360817128&lang=zh_CN#rd) 
* [Go goroutine 和 channel 入门](https://blog.csdn.net/realize_dream/article/details/121710134?spm=1001.2014.3001.5501)  
* [Go Goroutine调度器及面试精选](https://mp.weixin.qq.com/s?__biz=MzUyODgxNzM0Nw==&amp;mid=2247484244&amp;idx=1&amp;sn=a281424fa74d96e4a16ed256983fbad6&amp;chksm=fa6bcef8cd1c47ee16b628ae3296182c8d5e66a7bb98e39f5c0eabe6e2c0dca811113fb6129e&token=468025259&lang=zh_CN#rd) 
* [彻底明白Go语言的Channel了](https://mp.weixin.qq.com/s?__biz=MzUyODgxNzM0Nw==&amp;mid=2247484271&amp;idx=1&amp;sn=39d42387e9a14ac7a378970e7123b3b0&amp;chksm=fa6bcec3cd1c47d533e2cfa15c51b359d812e4c338f9452daa7ec2ce7902b799ba4e3d846276&token=468025259&lang=zh_CN#rd) 
* [Go 语言 channel 的阻塞问题](https://mp.weixin.qq.com/s?__biz=MzUyODgxNzM0Nw==&amp;mid=2247484283&amp;idx=1&amp;sn=c3de19c27e585a417bd224f92783c48b&amp;chksm=fa6bced7cd1c47c189f99283ce6fe9564c92a5e1f7ba4a28fe224394afb3e9c81167c3f5f246&token=468025259&lang=zh_CN#rd)
* [Go如何保证并发读写的顺序](https://mp.weixin.qq.com/s?__biz=MzUyODgxNzM0Nw==&amp;mid=2247484292&amp;idx=1&amp;sn=2727821035f7ee6b1a21c5feff816a7d&amp;chksm=fa6bce28cd1c473e863afd4add7b54750d77ebb5b08efc5aead897222a0d70c114afa4a90bbc&token=1208516942&lang=zh_CN#rd) 

## gRPC-go

0. [什么是 gRPC](https://github.com/metashops/GoFamily/blob/main/grpc/doc/0-%E4%BB%80%E4%B9%88%E6%98%AFgRPC.md)  

1. [grpc 入门实战](https://github.com/metashops/GoFamily/blob/main/grpc/doc/1-gRPC%E5%85%A5%E9%97%A8%E5%AE%9E%E6%88%98.md) 
2. [grpc服务器端启动时都做了哪些骚动作？](https://github.com/metashops/GoFamily/blob/main/grpc/doc/2-grpc%E6%9C%8D%E5%8A%A1%E5%90%AF%E5%8A%A8%E9%83%BD%E5%81%9A%E5%93%AA%E4%BA%9B%E9%AA%9A%E4%BD%9C.md) 
<<<<<<< HEAD
3. [grpc客户端与服务器端是如何建立链接？](https://github.com/metashops/GoFamily/blob/main/grpc/doc/3-grpc%E5%AE%A2%E6%88%B7%E7%AB%AF%E4%B8%8E%E6%9C%8D%E5%8A%A1%E5%99%A8%E7%AB%AF%E6%98%AF%E5%A6%82%E4%BD%95%E5%BB%BA%E7%AB%8B%E9%93%BE%E6%8E%A5%EF%BC%9F.md) 
=======
3. grpc客户端与服务器端是如何建立链接？
>>>>>>> links
4. 解析器
5. 平衡器
6. 帧接收和发送原理
7. grpc客户端是如何向grpc服务器发起rpc请求以及处理流程
8. 滑动窗口
9. keepalive
10. 多路复用multiplex

## 云原生

* Docker
  * Docker 基础
  * Dockerfile
* Kubernetes 基础
  * Kubernetes 背景
  * 为什么要 Pod
  * Pod 原理
  * Pod 生命周期
* Kubernetes 常用命令
* controller
  * ReplicaSet
  * Deployment
  * StatefulSet
  * DaemonSet
* configuration management
  * ConfigMap
  * Secret
  * ServiceAccount
* Network
  * Flannel
  * NetworkPolicy
  * Ingress
* store
  * Local 存储
  * Ceph
* monitor
  * Prometheus
* DevOps
  * Jenkins
  * Gitlab
* Istio
* Containerd
  * OCI & CRI
* 二次开发基础
  * Client-go 
    * Informer 机制
    * Reflect 源码
    * DeltaFIFO 源码
    * Indexer 源码
    * WorkQueue 源码
  * 编写控制器
    * CRD
    * 代码生成器
  * Operator
    * 什么是 Operator
    * Operator Framework 使用

## 边缘计算

* 为什么要边缘计算
* KubeEdge
* k3s
* 边缘计算落地实现



## 数据库相关

### Redis

* [01 Redis6-5种数据结构](https://github.com/metashops/GoFamily/blob/main/does/Redis/01%20Redis6-5%E7%A7%8D%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84.md) 
* [02 Redis6-配置文件](https://github.com/metashops/GoFamily/blob/main/does/Redis/02%20Redis6-%E9%85%8D%E7%BD%AE%E6%96%87%E4%BB%B6.md) 
* [03 Redis6-持久化](https://github.com/metashops/GoFamily/blob/main/does/Redis/03%20Redis6-%E6%8C%81%E4%B9%85%E5%8C%96.md) 
* [04 Redis6-事务](https://github.com/metashops/GoFamily/blob/main/does/Redis/04%20Redis6-%E4%BA%8B%E5%8A%A1.md) 
* [05 Redis6-发布订阅](https://github.com/metashops/GoFamily/blob/main/does/Redis/05%20Redis6-%E5%8F%91%E5%B8%83%E8%AE%A2%E9%98%85.md) 
* [06 Redis6-复制](https://github.com/metashops/GoFamily/blob/main/does/Redis/06%20Redis6-%E5%A4%8D%E5%88%B6.md) 
* [07 Redis6-哨兵模式](https://github.com/metashops/GoFamily/blob/main/does/Redis/07%20Redis6-%E5%93%A8%E5%85%B5%E6%A8%A1%E5%BC%8F.md) 
* [08 Redis6-Jedis](https://github.com/metashops/GoFamily/blob/main/does/Redis/08%20Redis6-Jedis.md) 
* [09 Redis6-其他数据结构](https://github.com/metashops/GoFamily/blob/main/does/Redis/09%20Redis6-%E5%85%B6%E4%BB%96%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84.md) 
* [10 Redis6-穿透|击穿|雪崩](https://github.com/metashops/GoFamily/blob/main/does/Redis/10%20Redis6-%E7%A9%BF%E9%80%8F%7C%E5%87%BB%E7%A9%BF%7C%E9%9B%AA%E5%B4%A9.md) 
* [11 Redis6-分布式锁](https://github.com/metashops/GoFamily/blob/main/does/Redis/11%20Redis6-%E5%88%86%E5%B8%83%E5%BC%8F%E9%94%81.md) 
* [12 Redis6-新功能介绍](https://github.com/metashops/GoFamily/blob/main/does/Redis/12%20Redis6-%E6%96%B0%E5%8A%9F%E8%83%BD%E4%BB%8B%E7%BB%8D.md) 
* [13 Redis6-SpringBoot整合](https://github.com/metashops/GoFamily/blob/main/does/Redis/13%20Redis6-SpringBoot%E6%95%B4%E5%90%88.md) 
* [14 Redis6-配置参考说明](https://github.com/metashops/GoFamily/blob/main/does/Redis/14%20Redis6-%E9%85%8D%E7%BD%AE%E5%8F%82%E8%80%83%E8%AF%B4%E6%98%8E.md) 
* [15 Redis-布隆过滤器](https://github.com/metashops/GoFamily/blob/main/does/Redis/15%20Redis-%E5%B8%83%E9%9A%86%E8%BF%87%E6%BB%A4%E5%99%A8.md) 
* [16 Redis-常见面试题](https://github.com/metashops/GoFamily/blob/main/does/Redis/16%20Redis-%E5%B8%B8%E8%A7%81%E9%9D%A2%E8%AF%95%E9%A2%98.md) 



## 框架

* [Gin ](https://github.com/gin-gonic/gin) 
* [GORM](https://gorm.io/docs/) 



## 优质博客

* [煎鱼](https://eddycjy.com/) 
* [鸟窝](https://colobu.com/) 
* [陈皓](https://coolshell.cn/) 
* [曹春晖](https://xargin.com/) 



## 不定期分享

* [defer关键字，都能扯很久？](https://mp.weixin.qq.com/s?__biz=MzUyODgxNzM0Nw==&amp;mid=2247484225&amp;idx=1&amp;sn=f41838df530ce9b920f6a180f164b3aa&amp;chksm=fa6bceedcd1c47fb0b103c9dc6fe0364a8e79b0a5658c45d8bab52ad8f72c7214f23758539a7&token=612778377&lang=zh_CN#rd)

## 辅助资料

> 以下是自己学习辅助资料仅供参考，如果你有更好的资料、书籍、文章等等，欢迎分享。

| **书名/课程**            | **备注** |
| :----------------------- | -------- |
| 《趣谈网络协议》         |          |
| 《Go 语言设计与实现》    |          |
| 《MySQL 45 讲》          |          |
| 《Redis 核心技术与实战》 |          |
| 《深入剖析 Kubernetes》  |          |
| 《Kubernetes 权威指南》  |          |



## 福利

> 关于Go语言学习之旅，推荐以下使用过的书籍及优质博客，希望能帮助到你。

**入门推荐**：

* [Effective Go](https://golang.google.cn/doc/effective_go.html)：官方文档，虽然英文，但也要读完

* Mastering GO 中文版：(知识点详细，适合初学者快速掌握)。
* Mastering GO 英文版
* Go 语言圣经：[在线博客](https://www.topgoer.cn/docs/gopl-zh/gopl-zh-1d29qo5qbk3js)
* Go 语言实战：(可以作为辅助)。
* Ultimate Go Notebook 英文版，
* Go 101：不断更新

**进阶推荐**：

* Go 专家编程：(可以辅助学习)，[在线博客](https://www.topgoer.cn/docs/gozhuanjia/gogfjhk)
* Go 语言核心编程
* Go 并发编程实战
* 其他书籍：Distributed Computing with Go 英文版

**推荐博客**：

* [Go中文网](http://c.biancheng.net/golang/)
* [地鼠文档](https://www.topgoer.cn/)：资料非常多，非常方便
* [Go 语言设计与实现](https://draveness.me/golang/)：非常棒，从使用到底层都给你讲一遍

**推荐视频**：

* 从来没有接触过编程：推荐B站韩顺平老师(课时很长需要坚持)
* 推荐国外课程：Ardan labs 可保持下来或 [阿里云盘在线观看](https://www.aliyundrive.com/s/ZNWXUhVVKxr)

PDF下载：

所有书籍点击 [下载入口](https://pan.baidu.com/s/1n3YQPU1baIkWhfExJIj0tg)，密码： 3fat      或关注回复：「电子书」

以上学习Go入门到入坑所参考的书籍，请供参考！



## 其他推荐

* 算法与数据结构体系课，[下载链接](https://pan.baidu.com/s/1t6lbcsU_N1iXotcm7Zoz4g)，密码：ls1u 
* 后端校招面试突击课，4年本科基础大复盘助力进大厂(2021版)，[下载链接](https://pan.baidu.com/s/1Av95FA22i-t4uzI2i9zZtw)，密码：3a5a

* 基于Spring Cloud微服务架构广告系统设计与实现(2020版)，[下载链接](https://pan.baidu.com/s/1U2oQNt1XmK2M79VsOmemLw)，密码：hc5s
* RPC框架核心源码深度解析，[下载链接](https://pan.baidu.com/s/1TQz5xc8sX6uVsngpwBQz9A)，密码：9odo
* 对标大厂 Redis6视频讲解，[下载链接](https://pan.baidu.com/s/1N2Af0V18EfoOFy6mN1bwhA)，密码：0475
* 大厂-JUC并发编程与源码分析，[下载链接](https://pan.baidu.com/s/1J_V1i_FAxdH9uVs6kbcJHQ)，密码：0rvq 

> 若部分链接失效，后台留言或者私信。

![](https://images.weserv.nl/?url=https://article.biliimg.com/bfs/article/113334bfc50d6a52e4a38e87c74c43a8c9f0685a.png)

---

扫码加微信，进群。。。

  <a name="公众号"></a>

![](http://ww1.sinaimg.cn/large/006FuVcvgy1gv0vb8cmr3j60by0by74s02.jpg)

---



新建了一个知识星球，用于免费分享技术文章、职场心得、面经等等，欢迎加入。

![](https://images.weserv.nl/?url=https://article.biliimg.com/bfs/article/79e99b22c32616c3e4c34b08487ed32c96c80e39.png)

