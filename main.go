package main

import (
	"log"

	"github.com/Kozzen890/project2-group2-glng-ks-08/databases"
	routers "github.com/Kozzen890/project2-group2-glng-ks-08/router"
)

func main() {
	port := ":8080"
	log.Printf("Starting server on port %s\n", port)
	databases.StartDB()
	routers.InitRouter().Run(port)
}