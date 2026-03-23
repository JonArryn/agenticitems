package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)


type envelope map[string]any

// extracts id named parameter from url path
func (app *application) readIdParam(r *http.Request) (int, error){
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

// completes the response cycle by Marshaling data to json with response code and writing any headers passed in
func (app *application) writeJsonResponse(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	// opted out of MarshalIndent because it's only beneficial for terminals and comes with
	// a significant cost
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// new line for terminals
	js = append(js, '\n')

    // write any optional headers passed in
	for key, values := range headers {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
