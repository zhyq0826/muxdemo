package router

import (
	"github.com/gorilla/mux"
	"github.com/zhyq0826/muxdemo/gitlab"
	"github.com/zhyq0826/muxdemo/home"
)

// InitRouter init all app route
func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router = home.InitRoutes(router)
	router = gitlab.InitRoutes(router)
	return router
}
