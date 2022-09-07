package resolver

import (
	"context"
	"gqlgen-todos/graph/generated"

	"cloud.google.com/go/firestore"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	// todos []*model.Todo
	// users []*model.User
	ctx    context.Context
	client *firestore.Client
}

func GetResolver(ctx context.Context, client *firestore.Client) *Resolver {
	return &Resolver{
		ctx:    ctx,
		client: client,
	}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
