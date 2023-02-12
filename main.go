package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	router := gin.Default()
	router.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})
	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		fmt.Printf("Server could not start :: %+v", err)
	}
}
