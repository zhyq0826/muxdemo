package home

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

// InitRoutes init domain app route
func InitRoutes(router *mux.Router) *mux.Router {
	appRouter := router.PathPrefix("/").Subrouter()
	appRouter.HandleFunc("/index", home)
	appRouter.HandleFunc("/", home)
	return router
}

func home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/index.html"))
	t.Execute(w, "hello world")
}
