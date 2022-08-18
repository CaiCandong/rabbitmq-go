# rabbitMQ-golang 
## docker安装
``` shell 
    docker run -d --hostname my-rabbit --name rabbit -p 15672:15672 -p 5672:5672 -e RABBITMQ_DEFAULT_USER=center -e RABBITMQ_DEFAULT_PASS=123qwe -e RABBITMQ_DEFAULT_VHOST=center  rabbitmq:management
```
> - 账户：center
> - 密码：123qwe
> - v-host: center
> - 15672：控制台端口号
> - 5672：应用访问端口号
> - 管理web 的端口：http://localhost:15672

原文：https://github.com/rabbitmq/rabbitmq-tutorials/tree/master/go
## chapter01 Hello World 
使用队列完成Hello world
```
    go run .\chapter01\sender\send.go
    go run .\chapter01\reveiver\receive.go
```
## chapter02  Work Queues

## chapter03 Publish/Subscribe
```
    go run .\chapter03\sender\sender.go hello world
    go run .\chapter03\reveiver\receive.go //可以启动多个
```
使用`fanout`类型的exchange进行消息广播,发送方只需要将消息发送到指定的交换机即可   
- 默认交换机: 默认交换器，该交换器由空字符串（`""`）标识，类型为direct
- 临时队列: 空字符串作为队列名称,系统将自动分配队列名称
## chapter04 Routing
``` 
go run .\chapter04\receiver\receive.go info warn
go run .\chapter03\sender\sender.go warn "a warning"
```


