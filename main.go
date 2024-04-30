package main

import (
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

	routers.DomainName = util.GetEnv("DOMAIN")
	routers.HttpRegex, err = regexp.Compile(`https?:\/\/[a-zA-Z0-9%]*:?[a-zA-Z0-9%]*\/?[(a-z).\/?]*\/?[^\s]+\.[^\s]{2,}\/?.*`)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	addr := util.GetEnv("ADDR")
	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("./assets"))
	router.Handle("GET /assets/", http.StripPrefix("/assets/", fs))

	router.HandleFunc("POST /api/v1/create_random_link", routers.CreateRandomLink)
	router.HandleFunc("POST /api/v1/create_link", routers.CreateLink)
	router.HandleFunc("GET /{name}", routers.HandleRedirect)
	router.HandleFunc("GET /", routers.Index)

	server := http.Server{
		Addr:    addr,
		Handler: router,
	}

	log.Printf("INFO: Starting server on %s", addr)
	log.Println(server.ListenAndServe())
}
