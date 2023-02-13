package services

import (
	"context"
	"time"

	"github.com/Jamify-app/auth/encryption"
	"github.com/Jamify-app/auth/processes/identities/models"
	"github.com/Jamify-app/auth/processes/identities/repositories"
)

type IService interface {
	CreateIdentity(context.Context, *models.Identity) error
}

type service struct {
	repository repositories.IRepository
}

func NewService(repository repositories.IRepository) IService {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateIdentity(ctx context.Context, identity *models.Identity) error {
	salt, hash, err := encryption.EncryptPassword(identity.Token)
	if err != nil {
		return err
	}

	now := time.Now()

	identity.TokenHash = hash
	identity.TokenSalt = salt
	identity.CreatedAt = &now

	err = s.repository.CreateIdentity(ctx, identity)
	if err != nil {
		return err
	}

	return nil
}
