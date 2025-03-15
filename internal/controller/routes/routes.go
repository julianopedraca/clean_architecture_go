package routes

import (
	sptrans "julianopedraca/clean_architecture_go/external/sptrans_api"
	"log/slog"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	api := &sptrans.SptransApi{}

	server.GET("/authentication", func(c *gin.Context) {
		Authentication(c, api)
	})
	server.GET("/search-line", func(c *gin.Context) {
		lineNameString := c.Query("lineName")
		lineNameNumber, err := strconv.Atoi(lineNameString)
		if err != nil {
			slog.Error("Failed to convert string to integer", "error", err.Error())
		}
		SearchLine(c, api, lineNameNumber)
	})
}
