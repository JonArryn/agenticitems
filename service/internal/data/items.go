package data

import "time"

type Item struct {
	ID 					int 	  `json:"id"`
	CreatedAt 			time.Time `json:"created_at"`
	Name 				string 	  `json:"name"`
	Code 				string    `json:"code"`
	Description         string    `json:"description"`
	SellPriceCents      uint64    `json:"sell_price_cents"`
	PurchaseCostCents   uint64    `json:"purchase_cost_cents"`
	DeletedAt           time.Time `json:"deleted_at,omitzero"`
}