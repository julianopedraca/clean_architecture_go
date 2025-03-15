package routes

import (
	"net/http"

	apiInterface "julianopedraca/clean_architecture_go/interfaces"

	"github.com/gin-gonic/gin"
)

func SearchLine(c *gin.Context, api apiInterface.SptransApiInterface, lineNumber int) {

	respBody, err := api.SearchLine(lineNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": string(respBody)})
}
