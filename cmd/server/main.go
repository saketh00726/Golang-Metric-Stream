package main

import (
	"log"

	"metrics-app/internal/alert"
	"metrics-app/internal/api"
	"metrics-app/internal/database"
	"metrics-app/internal/grpcserver"
)

func main() {

	db := database.InitDB()

	alert.StartAlertEngine(db, 20)

	go grpcserver.StartGRPCServer(db)

	go api.StartRESTServer(db)

	log.Println("Application started successfully")

	select {}
}
