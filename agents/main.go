package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "agents ok\n")
	})

	log.Print("agents listening on :4001")
	log.Fatal(http.ListenAndServe(":4001", mux))
}
