package main

import (
	"anime/config"
	"anime/routes"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()
	routes := routes.RegisterRoutes()

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", routes))

}
