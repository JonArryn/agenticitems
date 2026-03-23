package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler{
	// initialize router
	router := httprouter.New()
	
	// register handler methods and routes
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/item", app.createItemHandler)
	router.HandlerFunc(http.MethodGet, "/v1/item/:id", app.showItemHandler)

	return router
}