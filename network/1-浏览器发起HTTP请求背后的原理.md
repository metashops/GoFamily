### 1、浏览器发起HTTP请求背后流程

流程图：



* 浏览器从URL解析出域名；根据域名查询DNS，获取到域名对应的IP地址。
* 浏览器使用IP地址与服务器建立TCP连接。如果使用HTTPS，那么会完成TLS/SSL握手。
* 完成建立连接后，需要构建HTTP的请求，构造过程是需要填充HTTP头部，如上下文所需要的。
* 通过连接发起HTTP请求，服务器完成资源的表述，把HTML页面，CSS等作为包体作为HTTP响应给浏览器。
* 浏览器在渲染引擎中，会去解析这个响应，根据这个响应如超链接，图片等资源。
* 浏览器再次发送新HTTP请求。



### 2、什么是Hypertext Transfer Protocol（HTTP）协议？

官方定义：[https://www.rfc-editor.org/rfc/rfc7230.html](https://www.rfc-editor.org/rfc/rfc7230.html)

是一种无状态、应用层、请求/应答方式运行的协议，它使用扩展的语义和自描述消息格式，与其于网络的超文本信息系统灵活互动。