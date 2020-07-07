package config

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}

func RunServer(router *mux.Router) {
	port := ReadEnv("routerPort", "8080")
	fmt.Println("Starting Web Server at port : " + port)
	err := http.ListenAndServe(": "+port, router)
	if err != nil {
		log.Fatal(err)
	}
}