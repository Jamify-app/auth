package repositories

import (
	"context"

	"github.com/Jamify-app/auth/processes/identities/models"
)

func (r *repository) CreateIdentity(ctx context.Context, identity *models.Identity) error {
	collection := r.store.Database("auth").Collection("identities")

	_, err := collection.InsertOne(context.TODO(), identity)
	if err != nil {
		return err
	}

	return nil
}
