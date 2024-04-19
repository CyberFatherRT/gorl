package main

import (
	"fmt"
	util "gorl/pkg"
	db "gorl/pkg/db"
	routers "gorl/pkg/routers"
	"log"
	"net/http"
	"regexp"
)

func init() {
	var err error
	db.Init()
	routers.HttpRegex, err = regexp.Compile(`https?:\/\/[a-zA-Z0-9%]*:?[a-zA-Z0-9%]*\/?[(a-z).\/?]*\/?[^\s]+\.[^\s]{2,}\/?.*`)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	addr := fmt.Sprintf("%s", util.GetEnv("ADDR"))
	domain_name := util.GetEnv("DOMAIN")

	router := http.NewServeMux()

	router.HandleFunc("/", routers.Index)
	router.HandleFunc("/api/v1/get_random_link", func(w http.ResponseWriter, r *http.Request) {
		routers.GetRandLink(w, r, domain_name)
	})
	router.HandleFunc("/api/v1/get_link", func(w http.ResponseWriter, r *http.Request) {
		routers.GetRandLink(w, r, domain_name)
	})

	server := http.Server{
		Addr:    addr,
		Handler: router,
	}

	log.Printf("INFO: Starting server on %s", addr)
	log.Println(server.ListenAndServe())
}
