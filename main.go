package main

import (
	util "gorl/pkg"
	db "gorl/pkg/db"
	"gorl/pkg/middleware"
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

	assets := http.FileServer(http.Dir("assets"))
	router.Handle("GET /assets/", http.StripPrefix("/assets/", assets))

	router.HandleFunc("GET /", routers.Index)
	router.HandleFunc("GET /admin", routers.Admin)
	router.HandleFunc("GET /{name}", routers.HandleRedirect)

	router.HandleFunc("POST /api/v1/create_link", routers.CreateLink)
	router.HandleFunc("POST /api/v1/create_random_link", routers.CreateRandomLink)

	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.Authentication,
	)

	server := http.Server{
		Addr:    addr,
		Handler: stack(router),
	}

	log.Printf("INFO: Starting server on %s", addr)
	server.ListenAndServe()
}
