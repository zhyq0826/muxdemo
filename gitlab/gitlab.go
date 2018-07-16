package gitlab

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zhyq0826/muxdemo/util"
)

// InitRoutes init domain app route
func InitRoutes(router *mux.Router) *mux.Router {
	appRouter := router.PathPrefix("/v1/gitlab").Subrouter()
	appRouter.Use(util.Middleware)
	appRouter.HandleFunc("/list", home)
	return router
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message": "ok"}`))
}
