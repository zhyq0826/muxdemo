package gitlab

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zhyq0826/muxdemo/util"
)

// InitRoutes init domain app route
func InitRoutes(router *mux.Router) *mux.Router {
	fmt.Println("\033[32m gitlab subrouter init")
	fmt.Println("\033[37m")
	appRouter := router.PathPrefix("/v1/gitlab").Subrouter()
	appRouter.Use(util.Middleware)
	appRouter.HandleFunc("/list", home)
	return router
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message": "ok"}`))
}
