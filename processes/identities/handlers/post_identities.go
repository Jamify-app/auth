package handlers

import (
	"net/http"

	"github.com/Jamify-app/auth/processes/identities/models"
	"github.com/Jamify-app/auth/processes/identities/services"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostIdentitiesRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type PostIdentitiesResponse struct {
	UserIdentifier primitive.ObjectID `json:"user_identifier"`
}

func (r *PostIdentitiesRequest) Decode(ctx *gin.Context) error {
	return ctx.Bind(r)
}

func PostIdentities(service services.IService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := new(PostIdentitiesRequest)
		if err := request.Decode(ctx); err != nil {
			ctx.Error(errors.Wrapf(err, "Could not decode request :: %v"))
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Email and password are required",
			})
			return
		}

		identity := &models.Identity{
			Email: request.Email,
			Token: request.Password,
		}

		identity, err := service.CreateIdentity(ctx, identity)

		if err != nil {
			ctx.Error(errors.Wrapf(err, "Could not create identity :: %v"))
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Could not create identity",
			})
			return
		}

		ctx.JSON(http.StatusCreated, PostIdentitiesResponse{
			UserIdentifier: identity.ID,
		})
	}
}
