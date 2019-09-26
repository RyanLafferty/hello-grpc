package firestore;

import (
	"context"
	"log"
	"cloud.google.com/go/firestore"
)

func FetchLineItem(client *firestore.Client, ctx context.Context, id string) (map[string] interface{}) {
	collection := client.Collection("line_items")
	query := collection.Doc(id)
	lineItem, err := query.Get(ctx)

	if err != nil {
		log.Printf("Error[lib/db/firestore/FetchLineItem]: %+v", err)
		return nil
	}

	lineItemData := lineItem.Data()
	return lineItemData
}
