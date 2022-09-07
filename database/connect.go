package database

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

// func main() {
// 	ctx, client := GetContextAndClient()
// 	add(ctx, client)
// 	defer client.Close()
// }

func GetContextAndClient() (context.Context, *firestore.Client) {

	ctx := context.Background()
	opt := option.WithCredentialsFile("database/secret.json")
	config := &firebase.Config{}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return ctx, client
}
