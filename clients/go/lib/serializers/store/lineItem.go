package firestore

import (
	storePb "../../../store"
)

type LineItem struct {
	Id uint32 `json:"id,omitempty"`
	Quantity uint32 `json:"quantity,omitempty"`
	TotalCost float64 `json:"total_cost,omitempty"`
	Product *storePb.Product `json:"product,omitempty"`
	Supplier *storePb.Supplier `json:"supplier,omitempty"`
	Source string `json:"source,omitempty"`
	Client string `json:"client,omitempty"`
}

func SerializeLineItem(lineItem *storePb.LineItem) (LineItem) {
	return LineItem{
		Id: lineItem.Id,
		Quantity: lineItem.Quantity,
		TotalCost: lineItem.TotalCost,
		Product: lineItem.Product,
		Supplier: lineItem.Supplier,
		Source: lineItem.Source,
		Client: "golang",
	}
}
