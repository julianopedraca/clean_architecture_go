package main

import (
	"julianopedraca/clean_architecture_go/internal/router"
)

func main() {
	server := router.Setup()
	server.Run(":8080")
}
