package data

import (
	"time"

	"agenticitemsapi.arryn.net/internal/validator"
)

type Item struct {
	ID                int       `json:"id"`
	CreatedAt         time.Time `json:"created_at"`
	Name              string    `json:"name"`
	Code              string    `json:"code"`
	Description       string    `json:"description"`
	SellPriceCents    Cents     `json:"sell_price_cents"`
	PurchaseCostCents Cents     `json:"purchase_cost_cents"`
	DeletedAt         time.Time `json:"deleted_at,omitzero"`
}

func ValidateItem(v *validator.Validator, input *Item) {
	v.Check(input.Code != "", "code", "must be provided")
	v.Check(len(input.Code) <= 25, "code", "must not be more than 25 bytes long")

	v.Check(input.Name != "", "name", "must be provided")
	v.Check(len(input.Name) <= 99, "name", "must not be more than 99 bytes long")

	v.Check(input.SellPriceCents > 0, "sell_price", "must be a positive value")
	v.Check(input.PurchaseCostCents > 0, "purchase_cost", "must be a positive value")

}
