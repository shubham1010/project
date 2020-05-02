package routers

import (
	"github.com/gorilla/mux"
	"github.com/shubham1010/project/backend/users/controllers"
)

func SetUsersRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/demo", controllers.Index)
	router.HandleFunc("/login", controllers.Login)
	router.HandleFunc("/user", controllers.User)
	router.HandleFunc("/newuser", controllers.NewUser)
	router.HandleFunc("/adduser", controllers.AddUser)
	return router
}
