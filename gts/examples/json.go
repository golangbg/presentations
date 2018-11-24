package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// START OMIT
type jsonRequest struct {
	Name string `json:"name"`
}

type jsonResponse struct {
	Capitalized string `json:"capitalized"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	var req jsonRequest
	json.NewDecoder(r.Body).Decode(&req)

	res := jsonResponse{Capitalized: strings.ToUpper(req.Name)}
	json.NewEncoder(w).Encode(res)
}

//END OMIT

func main() {
	r := mux.NewRouter() // mux is the Gorilla router
	r.HandleFunc("/", rootHandler)

	// Custom server
	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      r, // Here we link the router to our custom server
	}

	// Bring up that server!
	log.Printf("starting server on %q", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
