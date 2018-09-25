package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// START OMIT
type GreetHandler struct {
	Greeting string
}

func (g GreetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>%s</h1>", g.Greeting)
}

func main() {
	// Link the root path to rootHandler
	http.Handle("/", GreetHandler{"Blah blah blah дрън дрън дрън"})

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
