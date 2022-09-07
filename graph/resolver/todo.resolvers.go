package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gqlgen-todos/database"
	"gqlgen-todos/graph/model"
	"math/rand"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID,
	}
	// r.todos = append(r.todos, todo)
	firestoreDao := database.GetFirestoreDao(r.ctx, r.client)
	err := firestoreDao.AddTodo(todo)
	return todo, err
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	// return r.todos, nil
	firestoreDao := database.GetFirestoreDao(r.ctx, r.client)
	data, err := firestoreDao.GetTodos()

	todos := []*model.Todo{}
	for _, v := range data {
		todo := &model.Todo{
			ID:     v["ID"].(string),
			Text:   v["Text"].(string),
			Done:   v["Done"].(bool),
			UserID: v["UserID"].(string),
		}
		todos = append(todos, todo)
	}
	return todos, err
}
