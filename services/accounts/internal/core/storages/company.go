package storages

import (
	"context"

	"github.com/avalonprod/eliteeld/accounts/internal/core/models"
)

type Company interface {
	Create(ctx context.Context, input models.Company) error
	GetByCredentials(ctx context.Context, email, password string) (models.Company, error)
	IsDuplicate(ctx context.Context, email string) (bool, error)
	SetSession(ctx context.Context, companyID string, session models.Session) error
	GetByRefreshToken(ctx context.Context, refreshToken string) (models.Company, error)
}
