package routers

import (
	"github.com/gorilla/mux"
)

func InitRouting() *mux.Router{
	router := mux.NewRouter()
	router = SetUsersRouters(router)

	return router
}
