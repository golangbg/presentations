package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"html/template"

	"github.com/gorilla/mux"
)

// START OMIT
func rootHandler(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)

	tpl, err := template.New("").Parse("<html><body><h1>Hello, {{ .Name }}</h1></body></html>")
	if err != nil {
		fmt.Fprintf(w, "Parsing error: %v", err)
		return
	}

	data := map[string]interface{}{"Name": args["name"]}
	if err := tpl.Execute(w, data); err != nil {
		fmt.Fprintf(w, "Execution error: %v", err)
		return
	}
}

//END OMIT

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
