package main

import (
	"github.com/xhyonline/grpc-example/gen/user"
	"github.com/xhyonline/grpc-example/server"
	"log"
	"net"
	"google.golang.org/grpc"
)

func main() {
	// 新建 grpc 服务
	grpc := grpc.NewServer()
	// server 中实现了 proto 定义的接口
	s:=server.NewServer()
	// 二者做绑定
	user.RegisterUserServerServer(grpc, s)
	l, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	}
	// 服务启动
	err=grpc.Serve(l)
	if err!=nil {
		panic(err)
	}
}
