### 什么是 gGRPC？

**1、什么是 RPC**

远程过程调用（Remote Procedure Call，缩写为 RPC）是一个计算机通信协议。

**2、什么是 gRPC**

gRPC 是一个高性能、通用的开源RPC框架，其由 Google 主要面向移动应用开发并基于

HTTP/2 协议标准而设计，基于 ProtoBuf(Protocol Buffers) 序列化协议开发，且支持众多

开发语言。

gRPC 基于 HTTP/2 标准设计，带来诸如双向流、流控、头部压缩、单 TCP 连接上的多复用请求等特性。这些特性使得其在移动设备上表现更好，在一定的情况下更节省空间占用。

**3、gRPC 调用模型**

gRPC 模型官方图如下：

![](http://www.grpc.io/img/grpc_concept_diagram_00.png)

1. 客户端（gRPC Stub）在程序中调用某方法，发起 RPC 调用。
2. 对请求信息使用 Protobuf 进行对象序列化压缩（IDL）。
3. 服务端（gRPC Server）接收到请求后，解码请求体，进行业务逻辑处理并返回。
4. 对响应结果使用 Protobuf 进行对象序列化压缩（IDL）。
5. 客户端接受到服务端响应，解码请求体。回调被调用的 A 方法，唤醒正在等待响应（阻塞）的客户端调用并返回响应结果。

**4、什么是 Protobuf**

Protocol Buffers（Protobuf）是一种与语言、平台无关，可扩展的序列化结构化数据的数据描述语言，我们常常称其为 IDL，常用于通信协议，数据存储等等，相较于 JSON、XML，它更小、更快，因此也更受开发人员的青眯

**5、Proto3 基础语法**

```protobuf
syntax = "proto3";

package helloworld;
option go_package = "./;helloworld";

// The greeting service definition.
service Greeter {
    // Sends a greeting
    rpc SayHello (HelloRequest) returns (HelloReply) {}
  }
  
  // The request message containing the user's name.
  message HelloRequest {
    string name = 1;
  }
  
  // The response message containing the greetings
  message HelloReply {
    string message = 1;
  }
```

> 同一个message的每个字段都有唯一一个编号，并且建议终生这个编号都不要改变。

其中类型可以是以下几种类型：

```protobuf
数字类型： double、float、int32、int64、uint32、uint64、sint32、sint64: 存储长度可变的浮点数、整数、无符号整数和有符号整数
存储固定大小的数字类型：fixed32、fixed64、sfixed32、sfixed64: 存储空间固定
布尔类型: bool
字符串: string
bytes: 字节数组
messageType: 消息类型
enumType:枚举类型
```

**5.1 字段介绍**

**（1）枚举(Enumerations)**

枚举类型适用于提供一组预定义的值，选择其中一个。例如我们将性别定义为枚举类型。

```protobuf
message Student {
  string name = 1;
  enum Gender {
    FEMALE = 0;
    MALE = 1;
  }
  Gender gender = 2;
  repeated int32 scores = 3;
}
```

**（2）使用其他消息类型**

```protobuf
message SearchResponse {
  repeated Result results = 1; 
}

message Result {
  string url = 1;
  string title = 2;
  repeated string snippets = 3;
}
```

或者写成嵌套也可以：

```protobuf
message SearchResponse {
  message Result {
    string url = 1;
    string title = 2;
    repeated string snippets = 3;
  }
  repeated Result results = 1;
}
```

**（3）任意类型(Any)**

Any 可以表示不在 .proto 中定义任意的内置类型。

```protobuf
import "google/protobuf/any.proto";

message ErrorStatus {
  string message = 1;
  repeated google.protobuf.Any details = 2;
}
```

**（4）oneof**

如果你有一组字段，同时最多允许这一组中的一个字段出现，就可以使用`Oneof` 定义这一组字段，这有点Union的意思，但是Oneof允许你设置零各值。`oneof` 字段不能同时使用`repeated`

```protobuf
message SampleMessage {
  oneof test_oneof {
    string name = 4;
    SubMessage sub_message = 9;
  }
}
```

**（5）map**

map类型需要设置键和值的类型。

```protobuf
message MapRequest {
  map<string, int32> points = 1;
}
```

**6、定义服务(Services)**

```protobuf
service SearchService {
  rpc Search (SearchRequest) returns (SearchResponse);
}
```

解释：

- 定义了一个名为 SearchService 的 RPC 服务
- SearchService 服务提供了 `Search`接口
- 入参为 SearchRequest 类型
- 返回参 SearchResponse 类型

> 下一篇：1-gRPC入门实战