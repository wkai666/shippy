# Golang 微服务学习总结（shippy 项目）

#### 项目案例：

[港口货物管理平台（shippy）](https://github.com/wkai666/shippy)

#### 微服务特点：

分布式运行、单一代码库、可独立运行测试和部署

#### 微服务优势：

1.  降低代码复杂性
2.  提高代码扩展性

#### RPC:

Remote Producedure Call 远程过程调用

#### gRPC:

谷歌开源的轻量级 RPC 通信框架，其中的通信协议基于二进制数据流，故此性能优异。并且 gRPC 支持 HTTP 2.0 协议，使用二进制帧进行数据传输，还可为通信双方建立持续的双向数据流。

#### protobuf:

gRPC 用于基于 HTTP 2.0 通信的二进制格式

#### 微服务开发流程：

1. 定义 protobuf 通信协议文件
2. 生成协议代码
3. 服务端与客户端实现协议
4. RPC 函数调用

#### Docker:

容器技术的一种实现，用以解决微服务架构中的缺陷

#### docker-compose:

用 docker-compose.yaml 来编排管理各个容器，同时设置容器的环境变量。

#### MongoDB: clone() 与 copy() 的区别：

+ clone: 新会话重用了主会话的 socket, 避免了资源浪费
- copy:  创建全新的会话

#### 踩坑总结：


1. gRPC 客户端版本差异(v2.9.1)：

![](@attachment/Clipboard_2021-12-02-17-48-56.png)

```
-- go-micro/v2/client/rpc_client.go

func newRpcClient(opt ...Option) Client {
	opts := NewOptions(opt...)

	p := pool.NewPool(
		pool.Size(opts.PoolSize),
		pool.TTL(opts.PoolTTL),
		pool.Transport(opts.Transport),
	)

	rc := &rpcClient{
		opts: opts,
		pool: p,
		seq:  0,
	}
	rc.once.Store(false)

	c := Client(rc)

	// wrap in reverse
	for i := len(opts.Wrappers); i > 0; i-- {
		c = opts.Wrappers[i-1](c)
	}

	return c
}

```

```
-- go-micro/v2/client/grpc/grpc.go

func newClient(opts ...client.Option) client.Client {
	options := client.NewOptions()
	// default content type for grpc
	options.ContentType = "application/grpc+proto"

	for _, o := range opts {
		o(&options)
	}

	rc := &grpcClient{
		opts: options,
	}
	rc.once.Store(false)

	rc.pool = newPool(options.PoolSize, options.PoolTTL, rc.poolMaxIdle(), rc.poolMaxStreams())

	c := client.Client(rc)

	// wrap in reverse
	for i := len(options.Wrappers); i > 0; i-- {
		c = options.Wrappers[i-1](c)
	}

	return c
}

```


##### 需要补习的知识：

1. protobuf 语法
2. docker 容器常用基础操作
3. docker network 容器连接
3. Dockerfile、Makefile 等
4. docker-compose 容器编排
5. MongoDB NoSQL 知识


#### 相关链接：

+ [微服务框架Go-Micro集成Nacos实战之服务注册与发现](https://segmentfault.com/a/1190000038287804)
+ [Golang 微服务教程](https://segmentfault.com/a/1190000015135650?utm_source=sf-similar-article)
+ [go-micro 微服务实践](https://www.jianshu.com/p/ec6d9c55809f)
+ [go-micro 微服务常见问题(- malformed HTTP response)](https://blog.csdn.net/hanyren/article/details/110943212)

