package main

import (
	"log"
	"os"

	"github.com/Kozzen890/project2-group2-glng-ks-08/databases"
	routers "github.com/Kozzen890/project2-group2-glng-ks-08/router"
)

func main() {
	port := os.Getenv("PORT")
	log.Printf("Starting server on port %s\n", port)
	databases.StartDB()
	routers.InitRouter().Run(port)
}