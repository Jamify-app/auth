package repositories

import (
	"context"

	"github.com/Jamify-app/auth/processes/identities/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type IRepository interface {
	CreateIdentity(context.Context, *models.Identity) error
}

type repository struct {
	store *mongo.Client
}

func NewRepository(store *mongo.Client) IRepository {
	return &repository{
		store: store,
	}
}
