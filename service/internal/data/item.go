package data

import (
	"database/sql"
	"errors"
	"time"

	"agenticitemsapi.arryn.net/internal/validator"
	"github.com/shopspring/decimal"
)

type Item struct {
	ID           int             `json:"id"`
	CreatedAt    time.Time       `json:"created_at"`
	Name         string          `json:"name"`
	Code         string          `json:"code"`
	Description  string          `json:"description"`
	SellPrice    decimal.Decimal `json:"sell_price"`
	PurchaseCost decimal.Decimal `json:"purchase_cost"`
	DeletedAt    *time.Time      `json:"deleted_at,omitempty"`
}

type RawItem struct {
	Name         string `json:"name"`
	Code         string `json:"code"`
	Description  string `json:"description"`
	SellPrice    string `json:"sell_price"`
	PurchaseCost string `json:"purchase_cost"`
}

func ValidateInputItem(v *validator.Validator, input *RawItem) *Item {

	v.Check(validator.NotBlank(input.Code), "code", "must be provided")
	v.Check(len(input.Code) <= 25, "code", "must not be more than 25 bytes long")

	v.Check(validator.NotBlank(input.Name), "name", "must be provided")
	v.Check(len(input.Name) <= 99, "name", "must not be more than 99 bytes long")

	v.Check(validator.MaxDecimalPlaces(input.SellPrice, 4), "sell_price", "Sell price must have 4 decimal places or less")
	v.Check(validator.MaxDecimalPlaces(input.PurchaseCost, 4), "purchase_cost", "Purchase cost must have 4 decimal places or less")

	v.Check(validator.NotBlank(input.SellPrice), "sell_price", "Sell price cannot be blank")
	v.Check(validator.NotBlank(input.PurchaseCost), "purchase_cost", "Purchase cost cannot be blank")

	sellPriceDecimal, err := decimal.NewFromString(input.SellPrice)

	if err != nil {
		v.AddError("sell_price", "Must be a valid number")
		return nil //no point in continuing if a number can't be parsed
	}

	purchaseCostDecimal, err := decimal.NewFromString(input.PurchaseCost)

	if err != nil {
		v.AddError("purchase_cost", "Must be a valid number")
		return nil

	}

	v.Check(sellPriceDecimal.IsPositive(), "sell_price", "must be a positive value")
	v.Check(purchaseCostDecimal.IsPositive(), "purchase_cost", "must be a positive value")

	if !v.Valid() {
		return nil
	}

	item := &Item{
		Name:         input.Name,
		Code:         input.Code,
		Description:  input.Description,
		SellPrice:    sellPriceDecimal,
		PurchaseCost: purchaseCostDecimal,
	}

	return item
}

type ItemModel struct {
	DB *sql.DB
}

func (i ItemModel) Insert(item *Item) error {
	query := `
	INSERT INTO items (code, name, description, sell_price, purchase_cost)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, code, created_at
	`

	args := []any{item.Code, item.Name, item.Description, item.SellPrice, item.PurchaseCost}

	return i.DB.QueryRow(query, args...).Scan(&item.ID, &item.Code, &item.CreatedAt)
}

func (i ItemModel) Get(id int) (*Item, error) {
	// validate id
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	// write query
	query := `
	SELECT id, code, name, description, sell_price, purchase_cost, created_at, deleted_at
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
		&item.SellPrice,
		&item.PurchaseCost,
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
