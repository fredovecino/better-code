package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//Command-line flags
	port := flag.String("p", "3000", "port to serve on")
	flag.Parse()

	router := mux.NewRouter()

	// Handle API routes
	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/", testHandle).Methods("GET")

	// Serve static files from public directory
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./public/"))))

	log.Println("Starting server on port " + *port)
	log.Fatal(http.ListenAndServe(":"+*port, router))

}

func testHandle(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Ok")
}
