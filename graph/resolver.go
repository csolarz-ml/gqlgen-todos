package graph

import "github.com/csolarz-ml/gqlgen-todos/repository"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen
type Resolver struct {
	TodoRepository repository.TodoRepository
}
