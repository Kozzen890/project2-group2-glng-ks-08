package main

import (
	"log"
	"os"

	"github.com/Kozzen890/project2-group2-glng-ks-08/databases"
	routers "github.com/Kozzen890/project2-group2-glng-ks-08/router"
)

func main() {
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	log.Printf("Starting server on port %s\n", PORT)
	databases.StartDB()
	routers.InitRouter().Run(":" + PORT)
}