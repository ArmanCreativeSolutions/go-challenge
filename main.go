package main

import (
	"github.com/gin-gonic/gin"
	"go-challenge/infrastructure/routes"
)

func main() {
	r := gin.Default()
	routes.SetupRouter(r)
	r.Run("localhost:8080")
}
