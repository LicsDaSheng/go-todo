package endpoints

import (
	"context"
	"go.uber.org/zap"
	"oncekey/go-todo/todo-svr/service"

	"github.com/go-kit/kit/endpoint"
)

type CreateTodoRequest struct {
	Todo service.TODO
}
type CreateTodoResponse struct {
	I   int64
	Err error
}

func NewEndpoints(srv service.Service, logger *zap.Logger) Endpoints {
	return Endpoints{
		CreateTodo: MakeCreateTodoEndpoint(srv),
		DeleteById: MakeDeleteByIdEndpoint(srv),
		UpdateById: MakeUpdateByIdEndpoint(srv),
		QueryTodo:  MakeQueryTodoEndpoint(srv),
		FindById:   MakeFindByIdEndpoint(srv),
	}
}
func MakeCreateTodoEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateTodoRequest)
		i, err := s.CreateTodo(ctx, req.Todo)
		return CreateTodoResponse{I: i, Err: err}, nil
	}
}

type DeleteByIdRequest struct {
	Id int64
}
type DeleteByIdResponse struct {
	Err error
}

func MakeDeleteByIdEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteByIdRequest)
		err := s.DeleteByID(ctx, req.Id)
		return DeleteByIdResponse{Err: err}, nil
	}
}

type UpdateByIdRequest struct {
	Todo service.TODO
}
type UpdateByIdResponse struct {
	Err error
}

func MakeUpdateByIdEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateByIdRequest)
		err := s.UpdateByID(ctx, req.Todo)
		return UpdateByIdResponse{Err: err}, nil
	}
}

type QueryTodoRequest struct {
	Todo service.TODO
}
type QueryTodoResponse struct {
	S   []service.TODO
	Err error
}

func MakeQueryTodoEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(QueryTodoRequest)
		slice1, err := s.QueryTodo(ctx, req.Todo)
		return QueryTodoResponse{S: slice1, Err: err}, nil
	}
}

type FindByIdRequest struct {
	Id int64
}
type FindByIdResponse struct {
	T   service.TODO
	Err error
}

func MakeFindByIdEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FindByIdRequest)
		T, err := s.FindByID(ctx, req.Id)
		return FindByIdResponse{T: T, Err: err}, nil
	}
}

type Endpoints struct {
	CreateTodo endpoint.Endpoint
	DeleteById endpoint.Endpoint
	UpdateById endpoint.Endpoint
	QueryTodo  endpoint.Endpoint
	FindById   endpoint.Endpoint
}
