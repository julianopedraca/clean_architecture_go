package routes

import (
	"net/http"

	apiInterface "julianopedraca/clean_architecture_go/interfaces"

	"github.com/gin-gonic/gin"
)

func SearchStops(c *gin.Context, api apiInterface.SptransApiInterface, stop string) {
	respBody, err := api.SearchStops(stop)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": respBody})
}
