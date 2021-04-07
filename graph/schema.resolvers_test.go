package graph

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/csolarz-ml/gqlgen-todos/graph/generated"
	"github.com/csolarz-ml/gqlgen-todos/graph/model"
	"github.com/csolarz-ml/gqlgen-todos/mocks"
	"github.com/stretchr/testify/require"
)

func TestMutationResolverCreateTodo(t *testing.T) {
	repositoryMock := new(mocks.MockTodoRepository)
	resolvers := Resolver{TodoRepository: repositoryMock}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers})))

	todo := model.Todo{Text: "go summercamp", UserID: "1"}

	repositoryMock.On("Save", &todo).Return(&todo)

	query := `
		mutation {
			createTodo(input: { text: "go summercamp", userId: "1" })
			{
			  id
			  text
			  done
			  user { id, name }
			}
		  }`

	var resp struct {
		CreateTodo struct {
			ID   string `json:"id"`
			Text string `json:"text"`
			Done bool   `json:"done"`
			User struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"user"`
		} `json:"createTodo"`
	}

	c.MustPost(query, &resp)

	repositoryMock.AssertExpectations(t)

	require.Equal(t, "", resp.CreateTodo.ID)
	require.Equal(t, "go summercamp", resp.CreateTodo.Text)
	require.Equal(t, false, resp.CreateTodo.Done)
	require.NotNil(t, resp.CreateTodo.User)
	require.Equal(t, "1", resp.CreateTodo.User.ID)
	require.Equal(t, "user 1", resp.CreateTodo.User.Name)
}

func TestQueryResolverTodos(t *testing.T) {

}
