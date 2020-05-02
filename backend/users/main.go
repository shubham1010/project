package main

import (
	"log"
	"net/http"
	"github.com/shubham1010/project/backend/users/common"
	"github.com/shubham1010/project/backend/users/routers"
	"github.com/shubham1010/project/backend/users/controllers"

)

func main() {
	log.Println("Starting...")

	common.StartUp()
	log.Println("Started StartUp()...")

	controllers.Init()

	router := routers.InitRoutes()
	log.Println("Started InitRoutes()...")

	server := &http.Server {
		Addr: common.AppConfig.Server,
		Handler: router,
	}

	log.Println("Listening...")
	server.ListenAndServe()
}
