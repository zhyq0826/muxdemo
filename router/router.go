package router

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/zhyq0826/muxdemo/gitlab"
	"github.com/zhyq0826/muxdemo/home"
)

// InitRouter init all app route
func InitRouter() *mux.Router {
	fmt.Println("new router init")
	router := mux.NewRouter()
	router = home.InitRoutes(router)
	router = gitlab.InitRoutes(router)
	return router
}
