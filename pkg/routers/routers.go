package routers

import (
	"encoding/json"
	"fmt"
	. "gorl/pkg"
	"gorl/pkg/db"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Request struct {
	Link string `json:"link"`
}

var HttpRegex *regexp.Regexp

func GetLink(w http.ResponseWriter, r *http.Request, domain_name string) {
	log.Printf("INFO: url := %s", r.URL)
}

func GetRandLink(w http.ResponseWriter, r *http.Request, domain_name string) {
	log.Printf("INFO: url := %s", r.URL)
	var request Request

	log.Printf("Get request from: %s", r.URL)

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !HttpRegex.Match([]byte(request.Link)) {
		http.Error(w, "You must supply valid http link", http.StatusBadRequest)
	}

	link, err := db.DB.GenLink(RandStringRunes(5), request.Link)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"link\": \"%s/%s\"}", domain_name, link)
}

func Index(w http.ResponseWriter, r *http.Request) {
	log.Printf("INFO: url := %s", r.URL)
	if r.URL.Path != "/" {

		name := strings.TrimPrefix(r.URL.Path, "/")
		link, err := db.DB.GetLink(name)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		http.Redirect(w, r, link, http.StatusMovedPermanently)
		return
	}

	html, _ := os.ReadFile("index.html")
	fmt.Fprintf(w, string(html))
}
