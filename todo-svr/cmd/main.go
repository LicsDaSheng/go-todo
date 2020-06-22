package main

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/hwholiday/learning_tools/all_packaged_library/logtool"
	"go.uber.org/zap"
	gRPC "google.golang.org/grpc"
	"net"
	"oncekey/go-todo/todo-svr/endpoints"
	"oncekey/go-todo/todo-svr/grpc"
	"oncekey/go-todo/todo-svr/proto"
	"oncekey/go-todo/todo-svr/service"
	"os"
)

var logger *zap.Logger

func main() {
	// 初始化日志
	newLoggerServer()
	server := service.NewService(logger)
	edps := endpoints.NewEndpoints(server, logger)
	gRPCServer := grpc.NewGRPCServer(edps, logger)
	grpcListener, err := net.Listen("tcp", ":8888")
	if err != nil {
		logger.Warn("Listen", zap.Error(err))
		os.Exit(0)
	}
	baseServer := gRPC.NewServer(gRPC.UnaryInterceptor(grpctransport.Interceptor))

	proto.RegisterTodoServiceServer(baseServer, gRPCServer)
	if err = baseServer.Serve(grpcListener); err != nil {
		logger.Warn("Serve", zap.Error(err))
		os.Exit(0)
	}
}

func newLoggerServer() {
	logger = logtool.NewLogger(
		logtool.SetAppName("todo-svr"),
		logtool.SetDevelopment(true),
		logtool.SetLevel(zap.DebugLevel),
	)
}
