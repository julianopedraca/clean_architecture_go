package routes

import (
	sptrans "julianopedraca/clean_architecture_go/external/sptrans_api"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	api := &sptrans.SptransApi{}

	server.GET("/authentication", func(c *gin.Context) {
		Authentication(c, api)
	})
	server.GET("/search-line", func(c *gin.Context) {
		line := c.Query("line")
		SearchLine(c, api, line)
	})
}
