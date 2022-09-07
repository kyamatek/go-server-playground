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

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		ID:   fmt.Sprintf("T%d", rand.Int()),
		Name: input.Name,
	}
	// r.users = append(r.users, user)
	firestoreDao := database.GetFirestoreDao(r.ctx, r.client)
	err := firestoreDao.AddUser(user)
	return user, err
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	// return r.users, nil
	firestoreDao := database.GetFirestoreDao(r.ctx, r.client)
	data, err := firestoreDao.GetUsers()

	users := []*model.User{}
	for _, v := range data {
		user := &model.User{
			ID:   v["ID"].(string),
			Name: v["Name"].(string),
		}
		users = append(users, user)
	}
	return users, err

}
