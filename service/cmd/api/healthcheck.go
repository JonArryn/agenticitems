package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request){

	data := envelope{
		"status": "available",
		"system_info" :map[string]string{
			"environment": app.config.env,
			"version": version,
		},
	}

	err := app.writeJsonResponse(w, http.StatusOK, data, nil)
	if err != nil {
		app.errorResponse(w, r, 500, err)
	}
}