package main

import (
	"github.com/gin-gonic/gin"
	"rest.com/main/db"
	"rest.com/main/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
