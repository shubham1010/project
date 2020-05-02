package main

import (
	"net/http"
	"fmt"

	"github.com/shubham1010/project/temp/controller"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter();
	router.HandleFunc("/",controller.Hello).Methods("GET")

	http.ListenAndServe(":8080",router)
	fmt.Println("Listening on port 8080..")
}
