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

type LinkRequest struct {
	Link string `json:"long_link"`
}

var HttpRegex *regexp.Regexp
var DomainName string

func CreateLinkRouter(w http.ResponseWriter, r *http.Request) {
	log.Printf("INFO: url := %s", r.URL)
}

func GenerateRandomLinkRouter(w http.ResponseWriter, r *http.Request) {
	var request LinkRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Failed parsing Request", http.StatusBadRequest)
		return
	}

	if !HttpRegex.Match([]byte(request.Link)) {
		http.Error(w, "You must supply valid http link", http.StatusBadRequest)
	}

	link, err := db.GenerateRandomLink(RandString(4), request.Link)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"short_link\": \"%s/%s\"}", DomainName, link)
}
