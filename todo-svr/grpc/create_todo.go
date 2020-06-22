package grpc

import (
	"context"
	"oncekey/go-todo/todo-svr/endpoints"
	"oncekey/go-todo/todo-svr/proto"
	"oncekey/go-todo/todo-svr/service"
	"time"
)

func (s *gRPCServer) CreateTodo(ctx context.Context, req *proto.CreateTodoRequest) (*proto.CreateTodoReply, error) {
	_, rep, err := s.createTodo.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*proto.CreateTodoReply), nil

}

func DecodeCreateTodoRequest(ctx context.Context, in interface{}) (interface{}, error) {
	tmp := in.(*proto.CreateTodoRequest)

	todo := &service.TODO{
		Topic:      tmp.Topic,
		Creator:    tmp.Creator,
		CreateTime: time.Unix(tmp.CreateTime, 0).Format("2006-01-02 15:04:05"),
		Desc:       tmp.Desc,
	}
	return &endpoints.CreateTodoRequest{Todo: *todo}, nil
}

func DecodeCreateTodoResponse(_ context.Context, in interface{}) (interface{}, error) {
	tmp := in.(*endpoints.CreateTodoResponse)

	out := &proto.CreateTodoReply{}
	if tmp.Err != nil {
		out.Code = 1
		out.Msg = tmp.Err.Error()
	}
	return &proto.CreateTodoReply{
		Code: 0,
		Msg:  "成功",
		Id:   tmp.I,
	}, nil
}
