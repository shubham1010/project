package main

import (
	"net/http"

	"github.com/shubham1010/project/temp/routers"
	"github.com/shubham1010/project/temp/common"

	log "github.com/sirupsen/logrus"

)

func main() {
	log.Info("[MAIN]: Starting App")
	common.StartUp()
	log.Info("[COMMON]: Configuration is Done")

	router := routers.InitRouting()
	log.Info("[ROUTES]: Initailized")

	server := &http.Server {
        Addr: common.Env.Server,
        Handler: router,
    }

	log.Info("[LISTENING]: Port 8080")
    server.ListenAndServe()

}
