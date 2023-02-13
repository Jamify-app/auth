package models

import "time"

type Identity struct {
	Email     string     `bson:"email"`
	Token     string     `bson:"-"`
	TokenHash []byte     `bson:"token_hash"`
	TokenSalt []byte     `bson:"token_salt"`
	CreatedAt *time.Time `bson:"created_at"`
}
