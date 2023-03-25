package service_tests

import (
	"testing"

	"github.com/Jamify-app/auth/processes/identities/models"
	"github.com/Jamify-app/auth/processes/identities/repositories"
	"github.com/Jamify-app/auth/processes/identities/services"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestCreateIdentity(t *testing.T) {
	tt := map[string]struct {
		give                        *models.Identity
		want                        error
		mockReturnCreateIdentityErr error
		mockReturnGetIdentity       *models.Identity
		mockReturnGetIdentityErr    error
		shouldSucceed               bool
	}{
		"Should successfully encrypt password and create new identity": {
			give: &models.Identity{
				Email: "test@test.com",
				Token: "password1234!",
			},
			want:                        nil,
			mockReturnCreateIdentityErr: nil,
			mockReturnGetIdentity:       nil,
			mockReturnGetIdentityErr:    mongo.ErrNoDocuments,
			shouldSucceed:               true,
		},
		"Should successfully find existing identity": {
			give: &models.Identity{
				Email: "test@test.com",
				Token: "password1234!",
			},
			want:                        nil,
			mockReturnCreateIdentityErr: nil,
			mockReturnGetIdentity:       nil,
			mockReturnGetIdentityErr:    nil,
			shouldSucceed:               true,
		},
		"Should not encrypt password and fail to create identity": {
			give: &models.Identity{
				Email: "test@test.com",
				Token: "",
			},
			want:                        assert.AnError,
			mockReturnCreateIdentityErr: nil,
		},
		"Should successfully encrypt password, but fail to create identity": {
			give: &models.Identity{
				Email: "test@test.com",
				Token: "password1234!",
			},
			want:                        assert.AnError,
			mockReturnCreateIdentityErr: assert.AnError,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			repo := new(repositories.MockIRepository)

			service := services.NewService(repo)

			repo.
				On("GetIdentity", nil, tc.give.Email).
				Return(tc.mockReturnGetIdentity)

			repo.
				On("CreateIdentity", nil, tc.give).
				Return(tc.mockReturnCreateIdentityErr)

			_, err := service.CreateIdentity(nil, tc.give)

			if tc.shouldSucceed {
				assert.Empty(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}

}
