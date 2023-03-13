package identities

import (
	"github.com/Jamify-app/auth/processes/identities/handlers"
	"github.com/Jamify-app/auth/processes/identities/services"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(group *gin.RouterGroup, service services.IService) {
	group.POST("", handlers.PostIdentities(service))
}
