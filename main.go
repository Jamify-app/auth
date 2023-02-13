package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Jamify-app/auth/middleware"
	"github.com/Jamify-app/auth/processes/identities"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx := context.Background()

	identitiesProcess := identities.NewProcess()

	port := os.Getenv("PORT")

	router := gin.New()
	router.Use(middleware.StructuredLogger(&log.Logger))

	identitiesProcess.Start(ctx, router)
	defer identitiesProcess.Stop(ctx)

	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		fmt.Printf("Server could not start :: %+v", err)
	}
}
