package routers

import (
	"net/http"
)

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
