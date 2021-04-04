package mocks

import (
	"github.com/csolarz-ml/gqlgen-todos/graph/model"
	"github.com/stretchr/testify/mock"
)

type TodoRepositoryMock struct {
	mock.Mock
}

func (r *TodoRepositoryMock) Save(todo *model.Todo) {
}

func (r *TodoRepositoryMock) Find() []*model.Todo {
	args := r.Called()
	return args.Get(0).([]*model.Todo)
}
