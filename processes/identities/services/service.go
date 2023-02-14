package services

import (
	"context"

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
