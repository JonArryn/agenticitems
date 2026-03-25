package main

import (
	"fmt"
	"net/http"
	"time"

	"agenticitemsapi.arryn.net/internal/data"
)

func (app *application) createItemHandler(w http.ResponseWriter, r *http.Request){
fmt.Fprintln(w, "create a new movie")
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