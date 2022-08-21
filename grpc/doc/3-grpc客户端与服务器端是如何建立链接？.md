### 1、gRPC 客户端与服务端建立连接的流程

**1.1 服务端在进行TCP连接前需要做什么事情？**

* 启动gRPC服务器时，主要做了初始化设置，拦截器设置，加密认证设置等等。
* 将提供的服务，如examples案例中的SayHello注册到 gRPC服务器里头。
* gRPC 服务器端启动监听端口，监听 gRPC 客户端发起的连接请求，如果没有请求，就会一直阻塞。

**1.2 TCP 连接之前，客户端也做啦哪些事情？**

* 启动 gRPC 客户端时，也是主要做啦初始化设置，比如服务连接地址设置，拦截器设置，是否阻塞式连接，连接安全性设置，加密认证设置。
* 解析器的最终目的，根据设置的服务连接地址，对此地址进行解析，最终获得改服务器后端对应的地址列表。
* 平衡器目的是，拿到解析器获取到的后端服务连接地址，从启动时注册的平衡器中获取链接策略，然后开始触发 TCP 链接。

**1.3 客户端与服务端真正建立连接**

* 客户端调用 go 原生 net包 Dialer 对象，向 gRPC 服务器发起连接请求。
* gRPC 服务器接收到 gRPC 客户端发起的连接请求后，专门创建一个 goroutine 来处理此请求，这里客户端的每一次请求，服务器端都会单独创建一个协程来处理。

**1.4 帧设置阶段**

* gRPC 采用 http2协议，因此在TCP链接建立后，需要进行帧的设置，双方同步一下信息，比如帧大小的设置，初始化窗口大小设置及PRI校验等。
* 帧设置好后，接下来 gRPC 客户端就可以使用新创建的 TCP 链接向 gRPC 服务器端发送 RPC请求啦，如调用SayHello 方法。

### 2、gRPC 服务器源码分析

**2.1 服务器端是如何启动和注册 gRPC服务？**

以官网例子为切入点讲解

```go
func main() {
	// 创建一个监听，端口号是port，协议是TCP
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
  // 创建 gRPC 服务器
	s := grpc.NewServer()
  // 将能够提供的服务，注册到 gRPC服务器里头
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
  // 启动 gRPC 服务器进行监听
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```

**2.2 服务器端是如何知道客户端的 TCP 请求？**

源码位置：grpc-go/server.go

```go
func (s *Server) Serve(lis net.Listener) error {
  // 表示服务器启动
	s.serve = true
  // 监听的进行存储
	ls := &listenSocket{Listener: lis}
  // 使用循环阻塞式监听 gRPC 客户端的请求
	for {
    // 阻塞方式监听客户端请求，如果没有请求时会一直阻塞在这里
		rawConn, err := lis.Accept()
    // 当监听到客户端请求来了，会创建一个协程来处理
		go func() {
			s.handleRawConn(lis.Addr().String(), rawConn)
			s.serveWG.Done()
		}()
	}
}
```

**2.3 服务器又是如何处理客户端的请求？**

```go
func (s *Server) handleRawConn(lisAddr string, rawConn net.Conn) {
  // 基本校验
	if s.quit.HasFired() {
		rawConn.Close()
		return
	}
  // 设置 TCP 链路的deadline
	rawConn.SetDeadline(time.Now().Add(s.opts.connectionTimeout))

  // HTTP2完成握手，此时跟客户端交换帧初始化信息，帧大小，窗口大小等
	// Finish handshaking (HTTP2)
	st := s.newHTTP2Transport(rawConn)
	rawConn.SetDeadline(time.Time{})
	if st == nil {
		return
	}

	if !s.addConn(lisAddr, st) {
		return
	}
  // 创建协程来处理流
	go func() {
		s.serveStreams(st)
		s.removeConn(lisAddr, st)
	}()
}
```

**以上，服务端与客户端已经建立起了 gRPC 链接，接下来就是处理客户端发送过来的数据。**

**2.4 服务器端和客户端帧握手阶段是如何实现？**

```go
func (s *Server) newHTTP2Transport(c net.Conn) transport.ServerTransport {
	// 省略。。。
	st, err := transport.NewServerTransport(c, config)
	
	return st
}
```

```go
func NewServerTransport(conn net.Conn, config *ServerConfig) (_ ServerTransport, err error) {
	// 通过newFramer创建http2的帧
	framer := newFramer(conn, writeBufSize, readBufSize, maxHeaderListSize)
  // 初始化设置帧，类型是[]http2.Setting切片（包括帧最大值，最大流数，初始化窗口大小等）
	// Send initial settings as connection preface to client.
	isettings := []http2.Setting{{
		ID:  http2.SettingMaxFrameSize,
		Val: http2MaxFrameLen,
	}}
	//
	if err := framer.fr.WriteSettings(isettings...); err != nil {
		return nil, connectionErrorf(false, err, "transport: %v", err)
	}
  // 如果需要，调整连接流控制窗口。这里是向客户端发送窗口更新帧
	// Adjust the connection flow control window if needed.
	if delta := uint32(icwz - defaultWindowSize); delta > 0 {
		if err := framer.fr.WriteWindowUpdate(0, delta); err != nil {
			return nil, connectionErrorf(false, err, "transport: %v", err)
		}
	}
	
	
	// Check the validity of client preface.
	preface := make([]byte, len(clientPreface))
	if _, err := io.ReadFull(t.conn, preface); err != nil {
		if err == io.EOF {
			return nil, io.EOF
		}
		return nil, connectionErrorf(false, err, "transport: http2Server.HandleStreams failed to receive the preface from client: %v", err)
	}
	if !bytes.Equal(preface, clientPreface) {
		return nil, connectionErrorf(false, nil, "transport: http2Server.HandleStreams received bogus greeting from client: %q", preface)
	}
  // 读取客户端发过来的帧，并转换为设置帧，然后将设置帧的内容更新到本地
	frame, err := t.framer.fr.ReadFrame()
	if err == io.EOF || err == io.ErrUnexpectedEOF {
		return nil, err
	}
	//
	sf, ok := frame.(*http2.SettingsFrame)
	if !ok {
		return nil, connectionErrorf(false, nil, "transport: http2Server.HandleStreams saw invalid preface type %T from client", frame)
	}
	t.handleSettings(sf)
  // 创建一个协程，来启动帧发送器
	go func() {
		t.loopy = newLoopyWriter(serverSide, t.framer, t.controlBuf, t.bdpEst)
		t.loopy.ssGoAwayHandler = t.outgoingGoAwayHandler
		if err := t.loopy.run(); err != nil {
			if logger.V(logLevel) {
				logger.Errorf("transport: loopyWriter.run returning. Err: %v", err)
			}
		}
		t.conn.Close()
		t.controlBuf.finish()
		close(t.writerDone)
	}()
  // 创建一个协程来启动keepalive，默认是开启的
	go t.keepalive()
	return t, nil
}
```

总结：以上是服务端完成啦跟客户端的帧交互，服务器端自己的帧大小，初始化窗口大小等等信息已经发送给客户端，同时也接收到客户端的帧信息，也在本地进行相应的设置。剩下就是处理流的阶段。

### 3、客户端源码分析

**3.1 gRPC 客户端连接参数设置**

```go
func main() {
  // 通过grpc。Dial跟grpc服务器端建立连接，参数包括：连接地址的设置、链路是否加密、是否阻塞式链接
  // 如grpc.WithBlocak说明是阻塞式链接，必须等到链路链接完成后，才能进行rpc请求，也就是调用SayHello
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 双方建立连接后，gRPC客户端做的事情，如调用SayHello方法
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}

```

Dial 源码

```go
func Dial(target string, opts ...DialOption) (*ClientConn, error) {
	return DialContext(context.Background(), target, opts...)
}
```

* 参数 target 的值，格式如：localhost:30001 or example://lib.example.grpc.io
* 参数 opts值，格式：grpc.WithInsecure() or grpc.WithBlock()
* 进入 DialContext 函数里（/grpc-go/clientconn.go）

**3.2 创建客户端连接器clientConn，并进行相应的初始化**

主要看DialContext这个函数，以下只看重点部分：

（1）首先会创建 ClientConn，初始化相关字段

```go
cc := &ClientConn{
		target:            target,
		csMgr:             &connectivityStateManager{},
		conns:             make(map[*addrConn]struct{}),
		dopts:             defaultDialOptions(),
		blockingpicker:    newPickerWrapper(),
		czData:            new(channelzData),
		firstResolveEvent: grpcsync.NewEvent(),
	}
```

（2）将用户设置的连接参数更新到客户端连接器 ClientConn

```go
for _, opt := range opts {
		opt.apply(&cc.dopts)
	}
```

>  用户所设置的链接参数，如是否链接阻塞，拦截器，超时机制，重试机制，keepalive等更新到cc.dopts中

（3）通过apply函数将参数更新到客户端连接器clientConn

```go
type DialOption interface {
	apply(*dialOptions)
}
```

DialOption 是个接口，定义了apply 函数，并且接收了dialOptions参数

funcDialOption 实现该接口

```go
// 创建结构，内部声明了func(*dialOptions)类型的变量f
type funcDialOption struct {
	f func(*dialOptions)
}
// 重写apply方法，参数为dialOptions
func (fdo *funcDialOption) apply(do *dialOptions) {
	fdo.f(do)
}
// 初始化funcDialOption，
func newFuncDialOption(f func(*dialOptions)) *funcDialOption {
	return &funcDialOption{
		f: f,
	}
}
```

