
# gin-micro-example
本项目使用清晰、简单的方式，示例gin开发过程中一些功能的使用，目前完成的功能示例：
- [x] 使用wire解决初始化的依赖问题
- [x] 接入prometheus：统计活跃的请求数与请求耗时
- [x] 增加微服务服务端示例
- [ ] 增加微服务客户端端示例
- [ ] 使用etcd完成服务注册与发现

## 接入prometheus
- 使用docker compose实现prometheus的部署
- 在8081端口暴露了采集接口metrics
- 使用了prometheus的Gauge与Summary指标，在gin的中间件中进行统计

可以使用wrk来进行请求，并在127.0.0.1:9090中查看实时指标
```shell
# 进行并发测试
go-wrk -c 20 -d 50 http://localhost:8080/users/hello
```

## 增加微服务服务端示例
增加了一个service同时被web与grpc调用的示例，启动后可以通过apipost的grpc方式调用/users/:id接口。目录及代码比较分散的原因是为了提高可扩展性能。
- 在/internal中增加service层,增加了一个获取用户信息的service
- 在/api/proto/user/v1中增加了proto文件，定义了GRPC可调用接口
- 使用buf简化proto文件的管理，buf配置在根目录的protobuf.gen.yaml中
- 在/grpc/user.go中实现了buf生成的接口，其依赖于第一步中的service层
- 在/pkg/grpcx/grpc.go中封装了grpc的启动，便于启动多个微服务的服务端
- 在根目录的app.go中封装了app,便于使用wire对启动的服务进行管理及扩展
- 通过wire解决grpc相关的依赖
> 相关软件的安装及命令
> **protobuf的安装**
> win：https://github.com/protocolbuffers/protobuf/rel
> mac：brew install protobuf
> **go和grpc插件的安装**
> go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
> go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
> **buf的安装**
> https://buf.build/docs/installation
> 生成命令：
> buf generate api/proto

## 增加微服务客户端调用
- 在ioc中增加了远程调用的初始化
- 在web层增加了rpc调用的server
- 修改了web层调用的参数


