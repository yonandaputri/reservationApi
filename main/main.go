package main

import (
	"hotelApp/config"
	"hotelApp/master"
)

func main() {
	db, _ := config.ConnectionDB()
	router := config.CreateRouter()
	master.Init(router, db)
	config.RunServer(router)
}