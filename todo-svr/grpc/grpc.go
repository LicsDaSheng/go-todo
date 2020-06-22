package grpc

import (
	"context"
	gRPCTransport "github.com/go-kit/kit/transport/grpc"
	"go.uber.org/zap"
	"oncekey/go-todo/todo-svr/endpoints"
	"oncekey/go-todo/todo-svr/proto"
)

type gRPCServer struct {
	createTodo gRPCTransport.Handler
	deleteTodo gRPCTransport.Handler
	updateTodo gRPCTransport.Handler
	queryTodo  gRPCTransport.Handler
	findById   gRPCTransport.Handler
}

func NewGRPCServer(endpoint endpoints.Endpoints, log *zap.Logger) proto.TodoServiceServer {
	options := []gRPCTransport.ServerOption{
		gRPCTransport.ServerErrorHandler(newZapLogErrorHandler(log)),
	}
	return &gRPCServer{
		createTodo: gRPCTransport.NewServer(
			endpoint.CreateTodo,
			DecodeCreateTodoRequest,
			DecodeCreateTodoResponse,
			options...,
		),
		deleteTodo: gRPCTransport.NewServer(
			endpoint.DeleteById,
			DecodeDeleteTodoRequest,
			DecodeDeleteTodoResponse,
			options...,
		),
	}
}

func (s *gRPCServer) UpdateTodo(ctx context.Context, req *proto.UpdateTodoRequest) (*proto.UpdateTodoReply, error) {
	_, rep, err := s.updateTodo.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*proto.UpdateTodoReply), nil
}
func (s *gRPCServer) QueryTodo(ctx context.Context, req *proto.QueryTodoRequest) (*proto.QueryTodoReply, error) {
	_, rep, err := s.queryTodo.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*proto.QueryTodoReply), nil
}

func (s *gRPCServer) FindById(ctx context.Context, req *proto.FindByIdRequest) (*proto.FindByIdReply, error) {
	_, rep, err := s.findById.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*proto.FindByIdReply), nil
}

type LogErrorHandler struct {
	logger *zap.Logger
}

func newZapLogErrorHandler(logger *zap.Logger) *LogErrorHandler {
	return &LogErrorHandler{
		logger: logger,
	}
}

func (h *LogErrorHandler) Handle(ctx context.Context, err error) {
	h.logger.Warn("gRPC invoke error", zap.Error(err))
}
