package main

import (
	"julianopedraca/clean_architecture_go/internal/controller/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
