package main

import (
	"fmt"
	db "gorl/pkg/db"
	"log"
	"net/http"
	"os"
)

func init() {
	db.Init()
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html, _ := os.ReadFile("index.html")
		fmt.Fprintf(w, string(html))
	})

	router.HandleFunc("/api/v1/get_link", func(w http.ResponseWriter, r *http.Request) {
		link := db.GetLink()
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "{\"link\": \"%s\"}", link)
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("INFO: Starting server on port :8080")
	server.ListenAndServe()
}
