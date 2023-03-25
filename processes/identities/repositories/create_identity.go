package repositories

import (
	"errors"

	"github.com/Jamify-app/auth/processes/identities/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *repository) CreateIdentity(ctx *gin.Context, identity *models.Identity) error {
	collection := r.store.Database("auth").Collection("identities")

	result, err := collection.InsertOne(ctx.Request.Context(), identity)
	if err != nil {
		return err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		identity.ID = oid
	} else {
		return errors.New("could not cast result Inserted ID as primitive Object ID")
	}

	return nil
}

func (r *repository) GetIdentity(ctx *gin.Context, email string) (*models.Identity, error) {
	collection := r.store.Database("auth").Collection("identities")

	identity := new(models.Identity)

	err := collection.FindOne(ctx.Request.Context(), bson.D{{Key: "email", Value: email}}).Decode(identity)
	if err != nil {
		return nil, err
	}

	return identity, nil
}
