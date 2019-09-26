package firestore;

import (
	"context"
	"log"
	"os"
	"cloud.google.com/go/firestore"
)

func Connect() (*firestore.Client, context.Context) {
	log.Printf("Connecting to firestore on: %s", os.Getenv("FIRESTORE_EMULATOR_HOST"))
	log.Printf("Using project: %s", os.Getenv("FIRESTORE_PROJECT_ID"))
	fsctx := context.Background()

	// create new firestore client
	client, err := firestore.NewClient(fsctx, os.Getenv("FIRESTORE_PROJECT_ID"))
	if err != nil {
		log.Printf("Error: Failed to connect to firestore: %+v", err)
		return nil, nil
	}

	return client, fsctx
}
