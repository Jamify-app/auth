package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})
	err := router.Run()
	if err != nil {
		fmt.Printf("Server could not start :: %+v", err)
	}
}
