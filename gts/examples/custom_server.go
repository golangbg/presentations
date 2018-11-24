package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Handlers are the link between routes and functionality
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello Plovdiv!</h1>")
}

// START OMIT
func main() {
	// Link the root path to rootHandler
	http.HandleFunc("/", rootHandler)

	// Custom server
	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Bring up that server!
	log.Printf("starting server on %q", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}

//END OMIT
