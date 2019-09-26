package store

import (
	"log"
	"context"
	"time"

	storePb "../../../store"
)

func GetLineItem(client storePb.StoreClient, id uint32) (*storePb.LineItem) {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	lineItem, err := client.GetLineItem(ctx, &storePb.LineItemRequest{Id: id})

	if err != nil {
		log.Printf("Error: Could not retreive line item: %v", err)
		return nil
	}

	log.Printf("Retreived line item: %v", lineItem)
	return lineItem
}
