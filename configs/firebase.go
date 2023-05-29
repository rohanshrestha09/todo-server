package configs

import (
	"context"
	"log"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var Bucket *storage.BucketHandle

func InitializeFirebaseApp() {
	var err error

	config := &firebase.Config{
		StorageBucket: "dev-synapse-0.appspot.com",
	}
	opt := option.WithCredentialsFile("./dev-synapse-0-firebase-adminsdk-nld47-c1ebb0b796.json")
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	Bucket, err = client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}
}
