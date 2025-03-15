package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	apiInterface "julianopedraca/clean_architecture_go/interfaces"
)

func Authentication(c *gin.Context, api apiInterface.SptransApiInterface) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	respBody, err := api.Authentication()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": string(respBody)})
}
