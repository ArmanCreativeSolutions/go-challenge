package main

import (
	"github.com/gin-gonic/gin"
	"go-challenge/infrastructure/dbconfig"
	"go-challenge/infrastructure/routes"
)

func main() {
	dbconfig.InitDB()
	r := gin.Default()
	routes.SetupRouter(r)
	r.Run(":8080")
}
