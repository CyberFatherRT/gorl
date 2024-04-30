package routers

import (
	"encoding/json"
	"fmt"
	. "gorl/pkg"
	"gorl/pkg/db"
	"log"
	"net/http"
	"regexp"
)

type Request struct {
	Link string `json:"link"`
}

var HttpRegex *regexp.Regexp
var DomainName string

func CreateLink(w http.ResponseWriter, r *http.Request) {
	log.Printf("INFO: url := %s", r.URL)
}

func CreateRandomLink(w http.ResponseWriter, r *http.Request) {
	var request Request

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if !HttpRegex.Match([]byte(request.Link)) {
		http.Error(w, "You must supply valid http link", http.StatusBadRequest)
	}

	link, err := db.DB.GenLink(RandString(5), request.Link)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"link\": \"%s/%s\"}", DomainName, link)
}

func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	link, err := db.DB.GetLink(name)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, link, http.StatusMovedPermanently)
}

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
