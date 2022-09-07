package database

import (
	"context"
	"gqlgen-todos/graph/model"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type FirestoreDao struct {
	ctx    context.Context
	client *firestore.Client
}

func GetFirestoreDao(ctx context.Context, client *firestore.Client) (firestoreDao *FirestoreDao) {
	firestoreDao = &FirestoreDao{
		ctx:    ctx,
		client: client,
	}
	return firestoreDao
}

func (f *FirestoreDao) GetUsers() (data []map[string]interface{}, err error) {
	iter := f.client.Collection("users").Documents(f.ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
			return nil, err
		}
		data = append(data, doc.Data())
	}

	return data, nil
}

func (f *FirestoreDao) AddUser(user *model.User) error {
	_, _, err := f.client.Collection("users").Add(f.ctx, user)
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
	return err
}

func (f *FirestoreDao) GetTodos() (data []map[string]interface{}, err error) {
	iter := f.client.Collection("todos").Documents(f.ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
			return nil, err
		}
		data = append(data, doc.Data())
	}

	return data, nil
}

func (f *FirestoreDao) AddTodo(user *model.Todo) error {
	_, _, err := f.client.Collection("todos").Add(f.ctx, user)
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
	return err
}
