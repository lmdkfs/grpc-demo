# 测试grpc服务

## 打包
make linux

## 生成镜像

make docker

> 默认端口为 50051

# grpc 测试工具igrpcurl

## mac install

`brew install grpcurl`

常用命令

- 查看 RPC 服务: grpcurl -plaintext 127.0.0.1:50051 list
- 查看 RPC 接口: grpcurl -plaintext 127.0.0.1:50051 list <service_name>
- 描述定义: grpcurl -plaintext 127.0.0.1:50051 list <message/rpc-interface>
- 调用方法: grpcurl -plaintext -d '{json_body}' 127.0.0.1:50051 <rpc-interface> 比如: grpcurl -plaintext -d '{"type": "1"}' 127.0.0.1:50051 libproto.EventRPC.ListEvents -d 后面加上 @ 支持从标准输入中读取流参数。
- 添加 metadata -H 'key:value'

# demo测试

` grpcurl -d '{"input":"1"}' -plaintext 127.0.0.1:50051   greet.GrpcService.grpcService`
