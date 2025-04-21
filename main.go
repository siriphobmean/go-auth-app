package main

import (
	"go-auth-app/config"
	"go-auth-app/database"
	"go-auth-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.ConnectDB()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}