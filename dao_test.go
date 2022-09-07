package main

import (
	"fmt"
	"gqlgen-todos/database"
	"gqlgen-todos/graph/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	user := &model.User{
		ID:   "xxx",
		Name: "yyy",
	}
	ctx, client := database.GetContextAndClient()
	firestoreDao := database.GetFirestoreDao(ctx, client)
	err := firestoreDao.AddUser(user)
	assert.Nil(t, err)
}

func TestGetUser(t *testing.T) {
	ctx, client := database.GetContextAndClient()
	firestoreDao := database.GetFirestoreDao(ctx, client)
	data, _ := firestoreDao.GetUsers()
	for _, v := range data {
		fmt.Println(v)
	}
}
