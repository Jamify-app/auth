package repositories

import (
	"github.com/Jamify-app/auth/processes/identities/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate mockery --name=IRepository --inpackage
type IRepository interface {
	CreateIdentity(*gin.Context, *models.Identity) error
	GetIdentity(*gin.Context, string) (*models.Identity, error)
}

type repository struct {
	store *mongo.Client
}

func NewRepository(store *mongo.Client) IRepository {
	return &repository{
		store: store,
	}
}
