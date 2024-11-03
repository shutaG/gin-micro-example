
# gin-micro-example
本项目使用清晰、简单的方式，示例gin开发过程中一些功能的使用，目前完成的功能示例：
- [x] 接入prometheus：统计活跃的请求数与请求耗时




## 接入prometheus
- 使用docker compose实现prometheus的部署
- 在8081端口暴露了采集接口metrics
- 使用了prometheus的Gauge与Summary指标，在gin的中间件中进行统计

可以使用wrk来进行请求，并在127.0.0.1:9090中查看实时指标
```shell
# 进行并发测试
go-wrk -c 20 -d 50 http://localhost:8080/users/hello
```
