package routes

import (
	"fmt"
	"net/http"

	apiInterface "julianopedraca/clean_architecture_go/interfaces"

	"github.com/gin-gonic/gin"
)

func SearchLine(c *gin.Context, api apiInterface.SptransApiInterface, line string) {
	respBody, err := api.SearchLine(line)
	fmt.Println("🚀 ~ file: search_line.go ~ line 13 ~ funcSearchLine ~ respBody : ", respBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": string(respBody)})
}
