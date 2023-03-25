package services

import (
	"github.com/Jamify-app/auth/processes/identities/models"
	"github.com/Jamify-app/auth/processes/identities/repositories"
	"github.com/gin-gonic/gin"
)

//go:generate mockery --name=IService --inpackage
type IService interface {
	CreateIdentity(*gin.Context, *models.Identity) (*models.Identity, error)
}

type service struct {
	repository repositories.IRepository
}

func NewService(repository repositories.IRepository) IService {
	return &service{
		repository: repository,
	}
}
