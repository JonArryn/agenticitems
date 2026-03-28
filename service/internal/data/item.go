package data

import (
	"database/sql"
	"errors"
	"time"

	"agenticitemsapi.arryn.net/internal/validator"
)

type Item struct {
	ID                int       `json:"id"`
	CreatedAt         time.Time `json:"created_at"`
	Name              string    `json:"name"`
	Code              string    `json:"code"`
	Description       string    `json:"description"`
	SellPriceCents    Cents     `json:"sell_price"`
	PurchaseCostCents Cents     `json:"purchase_cost"`
	DeletedAt         *time.Time `json:"deleted_at,omitempty"`
}

func ValidateItem(v *validator.Validator, input *Item) {
	v.Check(input.Code != "", "code", "must be provided")
	v.Check(len(input.Code) <= 25, "code", "must not be more than 25 bytes long")

	v.Check(input.Name != "", "name", "must be provided")
	v.Check(len(input.Name) <= 99, "name", "must not be more than 99 bytes long")

	v.Check(input.SellPriceCents > 0, "sell_price", "must be a positive value")
	v.Check(input.PurchaseCostCents > 0, "purchase_cost", "must be a positive value")

}

type ItemModel struct {
	DB *sql.DB
}

func (i ItemModel) Insert(item *Item) error {
	query := `
	INSERT INTO items (code, name, description, sell_price_cents, purchase_cost_cents)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, code, created_at
	`

	args := []any{item.Code, item.Name, item.Description, item.SellPriceCents, item.PurchaseCostCents}

	return i.DB.QueryRow(query, args...).Scan(&item.ID, &item.Code, &item.CreatedAt)
}

func (i ItemModel) Get(id int) (*Item, error) {
	// validate id
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	// write query
	query := `
	SELECT id, code, name, description, sell_price_cents, purchase_cost_cents, created_at, deleted_at
	FROM items 
	WHERE id=$1
	`
	// create an empty item struct
	var item Item

	// hit the db
	err := i.DB.QueryRow(query, id).Scan(
		&item.ID,
		&item.Code,
		&item.Name,
		&item.Description,
		&item.SellPriceCents,
		&item.PurchaseCostCents,
		&item.CreatedAt,
		&item.DeletedAt,
	)

	// handle db or retrieval errors
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	// return item
	return &item, nil
}

func (i ItemModel) Update(item *Item) error {
	return nil
}

func (i ItemModel) Delete(id int) error {
	return nil
}
