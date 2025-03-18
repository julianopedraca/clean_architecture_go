package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	apiInterface "julianopedraca/clean_architecture_go/interfaces"
)

func Authentication(c *gin.Context, api apiInterface.SptransApiInterface) {
	respBody, err := api.Authentication()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": string(respBody)})
}
