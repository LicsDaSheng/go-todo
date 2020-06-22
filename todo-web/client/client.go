package client

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"oncekey/go-todo/todo-svr/service"
)

func NewGRPCClient(conn *grpc.ClientConn, log *zap.Logger) service.Service {

}
