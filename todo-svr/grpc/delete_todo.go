package grpc

import (
	"context"
	"oncekey/go-todo/todo-svr/endpoints"
	"oncekey/go-todo/todo-svr/proto"
)

func (s *gRPCServer) DeleteTodo(ctx context.Context, req *proto.DeleteTodoRequest) (*proto.DeleteTodoReply, error) {
	_, rep, err := s.deleteTodo.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*proto.DeleteTodoReply), nil
}

func DecodeDeleteTodoRequest(_ context.Context, in interface{}) (interface{}, error) {
	req := in.(*proto.DeleteTodoRequest)
	return &endpoints.DeleteByIdRequest{Id: req.Id}, nil

}
func DecodeDeleteTodoResponse(_ context.Context, in interface{}) (interface{}, error) {
	resp := &proto.DeleteTodoReply{
		Code: 0,
		Msg:  "成功",
	}
	tmp := in.(*endpoints.DeleteByIdResponse)
	if tmp.Err != nil {
		resp.Code = 1
		resp.Msg = tmp.Err.Error()
		return tmp, nil
	}
	return resp, nil
}
