package service_tests

import (
	"context"
	"testing"

	"github.com/Jamify-app/auth/mocks"
	"github.com/Jamify-app/auth/processes/identities/models"
	"github.com/Jamify-app/auth/processes/identities/services"
	"github.com/stretchr/testify/assert"
)

func TestCreateIdentity(t *testing.T) {
	ctx := context.Background()

	tt := map[string]struct {
		give          *models.Identity
		want          error
		mockReturn    error
		shouldSucceed bool
	}{
		"Should successfully encrypt password and create identity": {
			give: &models.Identity{
				Email: "test@test.com",
				Token: "password1234!",
			},
			want:          nil,
			mockReturn:    nil,
			shouldSucceed: true,
		},
		"Should not encrypt password and fail to create identity": {
			give: &models.Identity{
				Email: "test@test.com",
				Token: "",
			},
			want:       assert.AnError,
			mockReturn: nil,
		},
		"Should successfully encrypt password, but fail to create identity": {
			give: &models.Identity{
				Email: "test@test.com",
				Token: "password1234!",
			},
			want:       assert.AnError,
			mockReturn: assert.AnError,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			repo := new(mocks.IRepository)

			service := services.NewService(repo)

			repo.
				On("CreateIdentity", ctx, tc.give).
				Return(tc.mockReturn)

			err := service.CreateIdentity(ctx, tc.give)

			if tc.shouldSucceed {
				assert.Empty(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}

}
