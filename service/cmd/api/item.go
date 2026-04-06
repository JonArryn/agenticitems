package main

import (
	"errors"
	"fmt"
	"net/http"

	"agenticitemsapi.arryn.net/internal/data"
	"agenticitemsapi.arryn.net/internal/validator"
)

func (app *application) createItemHandler(w http.ResponseWriter, r *http.Request) {
	// marshal json
	var input data.ItemInput

	err := app.readJson(w, r, &input)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// validate input data
	v := validator.New()

	item := data.ValidateInputItem(v, &input)

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// write to db
	err = app.models.Item.Insert(item)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrDuplicateCode):
			app.failedValidationResponse(w, r, map[string]string{
				"code": "an item with this code already exists",
			})
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// add location header
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("v1/items/%d", item.ID))

	// response
	err = app.writeJsonResponse(w, http.StatusCreated, envelope{"item": item}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showItemHandler(w http.ResponseWriter, r *http.Request) {
	// extract url params from request context
	id, err := app.readIdParam(r)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	item, err := app.models.Item.Get(id)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	itemEnvelope := envelope{"item": item}

	err = app.writeJsonResponse(w, http.StatusOK, itemEnvelope, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateItemHandler(w http.ResponseWriter, r *http.Request) {
	// extract url params from request context
	id, err := app.readIdParam(r)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	item, err := app.models.Item.Get(id)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// marshal json
	var input data.ItemInput

	err = app.readJson(w, r, &input)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// validate input data
	v := validator.New()

	validItem := data.ValidateInputItem(v, &input)

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	item.Code = validItem.Code
	item.Name = validItem.Name
	item.Description = validItem.Description
	item.SellPrice = validItem.SellPrice
	item.PurchaseCost = validItem.PurchaseCost

	err = app.models.Item.Update(item)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	itemEnvelope := envelope{"item": item}

	err = app.writeJsonResponse(w, http.StatusOK, itemEnvelope, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
