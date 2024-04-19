package main

import (
	"encoding/json"
	"fmt"
	util "gorl/pkg"
	db "gorl/pkg/db"
	"log"
	"net/http"
	"os"
	"strings"
)

type Request struct {
	Link string `json:"link"`
}

func init() {
	db.Init()
}

func main() {
	addr := fmt.Sprintf("%s", util.GetEnv("ADDR"))
	domain_name := util.GetEnv("DOMAIN")

	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)

		if r.URL.Path != "/" {
			name := strings.TrimPrefix(r.URL.Path, "/")
			link, err := db.DB.GetLink(name)

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			http.Redirect(w, r, link, 301)
			return
		}

		html, _ := os.ReadFile("index.html")
		fmt.Fprintf(w, string(html))
	})

	router.HandleFunc("/api/v1/get_link", func(w http.ResponseWriter, r *http.Request) {
		var request Request

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		link, err := db.DB.GenLink(request.Link)

		if err != nil {
			log.Fatal(err)
		}

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "{\"link\": \"%s/%s\"}", domain_name, link)
	})

	server := http.Server{
		Addr:    addr,
		Handler: router,
	}

	log.Printf("INFO: Starting server on %s", addr)
	log.Println(server.ListenAndServe())
}
