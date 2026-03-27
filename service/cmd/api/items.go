package main

import (
	"fmt"
	"net/http"
	"time"

	"agenticitemsapi.arryn.net/internal/data"
	"agenticitemsapi.arryn.net/internal/validator"
)

func (app *application) createItemHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name              string     `json:"name"`
		Code              string     `json:"code"`
		Description       string     `json:"description"`
		SellPriceCents    data.Cents `json:"sell_price"`
		PurchaseCostCents data.Cents `json:"purchase_cost"`
	}

	err := app.readJson(w, r, &input)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	item := &data.Item{
		Name:              input.Name,
		Code:              input.Code,
		Description:       input.Description,
		SellPriceCents:    input.SellPriceCents,
		PurchaseCostCents: input.PurchaseCostCents,
	}

	v := validator.New()

	if data.ValidateItem(v, item); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showItemHandler(w http.ResponseWriter, r *http.Request) {
	// extract url params from request context
	id, err := app.readIdParam(r)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	item := data.Item{
		ID:                id,
		Name:              "Test Item",
		CreatedAt:         time.Now().UTC(),
		Code:              "A0001",
		Description:       "A test item for testing stuff",
		SellPriceCents:    1224,
		PurchaseCostCents: 500,
	}

	itemEnvelope := envelope{"item": item}

	err = app.writeJsonResponse(w, http.StatusOK, itemEnvelope, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
