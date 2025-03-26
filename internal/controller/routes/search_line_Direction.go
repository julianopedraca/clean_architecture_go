package routes

import (
	"net/http"

	apiInterface "julianopedraca/clean_architecture_go/interfaces"

	"github.com/gin-gonic/gin"
)

func SearchLineDirection(c *gin.Context, api apiInterface.MobilityInterface, line string, direction string) {
	respBody, err := api.SearchLineDirection(line, direction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": respBody})
}
