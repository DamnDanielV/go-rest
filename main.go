package main

import (
	"log"
	"net/http"
	"time"

	"github.com/DamnDanielV/go-rest/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// r.HandleFunc("/{text}")
	r.HandleFunc("/users", routes.GetPosts).Methods("GET")
	r.HandleFunc("/users", routes.CreatePost).Methods("POST")

	// r.HandleFunc("/products", ProductsHandler)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
