package routers

import (
	"github.com/gorilla/mux"
	"github.com/shubham1010/project/temp/controllers"
)

func SetUsersRouters(router *mux.Router) *mux.Router{
	router.HandleFunc("/login",controllers.Login).Methods("POST")
	router.HandleFunc("/info",controllers.Info).Methods("GET")
	router.HandleFunc("/signup",controllers.SignUp).Methods("POST")

	return router

}
