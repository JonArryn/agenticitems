package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"agenticitemsapi.arryn.net/internal/data"
)

func (app *application) createItemHandler(w http.ResponseWriter, r *http.Request){
	var input struct {
		Name 				string 	  `json:"name"`
		Code 				string    `json:"code"`
		Description         string    `json:"description"`
		SellPriceCents      uint64    `json:"sell_price_cents"`
		PurchaseCostCents   uint64    `json:"purchase_cost_cents"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application)showItemHandler(w http.ResponseWriter, r *http.Request){
	// extract url params from request context
	id, err := app.readIdParam(r)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	item := data.Item{
		ID: id,
		Name: "Test Item",
		CreatedAt: time.Now().UTC(),
		Code:         "A0001",
		Description:  "A test item for testing stuff",
		SellPriceCents:    1224,
		PurchaseCostCents: 500,
	}

	itemEnvelope := envelope{"item": item}

	err = app.writeJsonResponse(w, http.StatusOK, itemEnvelope, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}