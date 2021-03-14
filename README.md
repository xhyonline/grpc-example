# grpc-example 

这是一个 Golang 使用 grpc 通信的示例

## 文件介绍

1. client.go 是客户端
2. server.go 是服务端
3. proto 目录是定义 `.proto` 文件的路径
4. gen 目录是通过 .proto 编译生成的 go 代码

## protoc 生成代码说明

请进入 proto 目录下进行命令操作

```
cd proto
protoc --go_out=plugins=grpc,paths=source_relative:. ./user/*.proto
```

```
命令解释:
其中，--go_out 参数是用来指定 protoc-gen-go 插件的工作方式 和 go 代码目录架构的生成位置，可以向 --go_out 传递很多参数，见 golang/protobuf 文档 。
主要的两个参数为 plugins 和 paths ，代表 生成 go 代码所使用的插件 和 生成的 go 代码的目录怎样架构。--go_out 参数的写法是，参数之间用逗号隔开，最后加上冒号来指定代码目录架构的生成位置，例如：--go_out=plugins=grpc,paths=import:. 。paths 参数有两个选项，import 和 source_relative 。默认为 import ，代表按照生成的 go 代码的包的全路径去创建目录层级，source_relative 代表按照 proto 源文件的目录层级去创建 go 代码的目录层级，如果目录已存在则不用创建。
```

因此当执行上面的命令后,将会在 user 目录下生成 `phone.pb.go`文件与 `user.pb.go` 文件,我们需要将这两个文件移动至 gen/user 下。

同理再执行命令

```
protoc --go_out=plugins=grpc,paths=source_relative:. ./basic/*.proto
```

生成的文件放入 gen/basic 下