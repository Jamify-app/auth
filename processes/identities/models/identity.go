package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Identity struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Email     string             `bson:"email"`
	Token     string             `bson:"-"`
	TokenHash []byte             `bson:"token_hash"`
	TokenSalt []byte             `bson:"token_salt"`
	CreatedAt *time.Time         `bson:"created_at"`
}
