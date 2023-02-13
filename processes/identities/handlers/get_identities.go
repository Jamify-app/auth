package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIdentities(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "get identities")
}
