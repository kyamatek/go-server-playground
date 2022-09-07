package main

import (
	"fmt"
	"gqlgen-todos/firestore"
	"gqlgen-todos/graph/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	user := &model.User{
		ID:   "xxx",
		Name: "yyy",
	}
	firestoreDao := firestore.GetFirestoreDao()
	err := firestoreDao.AddUser(user)
	assert.Nil(t, err)
}

func TestGetUser(t *testing.T) {
	firestoreDao := firestore.GetFirestoreDao()
	data, _ := firestoreDao.GetUsers()
	for _, v := range data {
		fmt.Println(v)
	}
}
