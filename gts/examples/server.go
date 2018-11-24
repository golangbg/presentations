package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handlers are the link between routes and functionality
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello Plovdiv!</h1>")
}

func main() {
	// Link the root path to rootHandler
	http.HandleFunc("/", rootHandler)

	// Bring up that server!
	log.Fatal(http.ListenAndServe(":8080", nil))
}
