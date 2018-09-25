package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// START OMIT
func rootHandler(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	fmt.Fprintf(w, "<h1>Hello %s!</h1>", args["name"])
}

func main() {
	r := mux.NewRouter() // mux is the Gorilla router
	r.HandleFunc("/{name}", rootHandler)

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

//END OMIT
