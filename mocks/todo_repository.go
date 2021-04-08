package mocks

import (
	"github.com/csolarz-ml/gqlgen-todos/graph/model"
	"github.com/stretchr/testify/mock"
)

type MockTodoRepository struct {
	mock.Mock
}

func (s *MockTodoRepository) Save(m *model.Todo) *model.Todo {
	args := s.Called(m)
	return args.Get(0).(*model.Todo)
}

func (s *MockTodoRepository) Find() []*model.Todo {
	args := s.Called()
	return args.Get(0).([]*model.Todo)
}
