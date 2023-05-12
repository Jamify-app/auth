package services

import (
	"errors"
	"time"

	"github.com/Jamify-app/auth/common/encryption"
	"github.com/Jamify-app/auth/processes/identities/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *service) CreateIdentity(ctx *gin.Context, identity *models.Identity) (*models.Identity, error) {
	salt, hash, err := encryption.EncryptPassword(identity.Token)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	identity.TokenHash = hash
	identity.TokenSalt = salt
	identity.CreatedAt = &now

	foundIdentity, err := s.repository.GetIdentity(ctx, identity.Email)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		}
	}

	if foundIdentity != nil {
		isPassSame, err := encryption.IsPasswordSame(identity.Token, foundIdentity.TokenHash, foundIdentity.TokenSalt)
		if err != nil {
			return nil, err
		}

		if isPassSame {
			return foundIdentity, nil
		}

		if foundIdentity.Email != "" {
			return nil, errors.New("identity exists, incorrect password")
		}
	}

	err = s.repository.CreateIdentity(ctx, identity)
	if err != nil {
		return nil, err
	}

	return identity, nil
}
