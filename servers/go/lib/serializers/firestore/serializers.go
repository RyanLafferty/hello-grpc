package firestore

import (
	storePb "../../../store"
)

func SerializeLineItem(lineItem map[string]interface{}) (*storePb.LineItem) {
	product := lineItem["product"].(map[string]interface {})
	supplier := lineItem["supplier"].(map[string]interface {})

	return &storePb.LineItem{
		Id: uint32(lineItem["id"].(int64)),
		Quantity: uint32(lineItem["quantity"].(int64)),
		TotalCost: lineItem["total_cost"].(float64),
		Source: "golang",
		Product: &storePb.Product{
			Name:  product["name"].(string),
			Cost: product["cost"].(float64),
		},
		Supplier: &storePb.Supplier{
			Name:  supplier["name"].(string),
			Country: supplier["country"].(string),
		},
	}
}
