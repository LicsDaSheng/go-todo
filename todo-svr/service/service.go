package service

import (
	"context"
	"errors"
	"go.uber.org/zap"
)

// Service todo service
type Service interface {
	CreateTodo(ctx context.Context, todo TODO) (int64, error)
	DeleteByID(ctx context.Context, id int64) error
	UpdateByID(ctx context.Context, todo TODO) error
	QueryTodo(ctx context.Context, todo TODO) ([]TODO, error)
	FindByID(ctx context.Context, id int64) (TODO, error)
}

// TODO model for TODO
type TODO struct {
	ID         int64  `json:"id"`
	Topic      string `json:"topic"`
	Creator    string `json:"creator"`
	CreateTime string `json:"createTime"`
	Desc       string `json:"desc"`
}

type impl struct {
	logger *zap.Logger
}

// NewService create a new Service
func NewService(logger *zap.Logger) Service {
	return &impl{
		logger: logger,
	}
}

// CreateTodo CreateTodo
func (s impl) CreateTodo(ctx context.Context, todo TODO) (int64, error) {
	panic(errors.New("CreateTodo not implemented"))
}

// DeleteByID DeleteById
func (s impl) DeleteByID(ctx context.Context, id int64) error {
	panic(errors.New("DeleteById not implemented"))
}

//UpdateByID UpdateById
func (s impl) UpdateByID(ctx context.Context, todo TODO) error {
	panic(errors.New("UpdateById not implemented"))
}

// QueryTodoes QueryTodoes
func (s impl) QueryTodo(ctx context.Context, todo TODO) ([]TODO, error) {
	panic(errors.New("QueryTodoes not implemented"))
}

// FindByID FindById
func (s impl) FindByID(ctx context.Context, id int64) (TODO, error) {
	panic(errors.New("FindById not implemented"))
}
