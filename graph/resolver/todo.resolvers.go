package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gqlgen-todos/firestore"
	"gqlgen-todos/graph/generated"
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
	firestoreDao := firestore.GetFirestoreDao()
	err := firestoreDao.AddTodo(todo)
	return todo, err
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	// return r.todos, nil
	firestoreDao := firestore.GetFirestoreDao()
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
