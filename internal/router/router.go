package router

import (
	"julianopedraca/clean_architecture_go/internal/controller/routes"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	server := gin.Default()
	routes.RegisterRoutes(server)
	return server
}
