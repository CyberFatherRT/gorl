package routers

import (
	"gorl/pkg/db"
	"net/http"
)

func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	link, err := db.GetLink(name)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, link, http.StatusMovedPermanently)
}

func Handle401(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/401.html")
}

func Handle403(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/403.html")
}

func Handle404(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/404.html")
}

func Admin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/admin.html")
}

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}
