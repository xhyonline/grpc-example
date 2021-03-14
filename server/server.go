package server

import (
	"github.com/xhyonline/grpc-example/rpc"
)

// server 是 rpc 服务的集合,它继承了所有的方法
type server struct {
	rpc.UserServer
}

// NewServer
func NewServer() *server{
	return new(server)
}


