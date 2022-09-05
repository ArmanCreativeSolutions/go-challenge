package main

import (
	"github.com/gin-gonic/gin"
	"go-challenge/infrastructure/routes"
	"go-challenge/presentation/cronjobs"
	"log"
)

func main() {
	r := gin.Default()
	routes.SetupRouter(r)
	go cronjobs.RunCronJobs()
	err := r.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
