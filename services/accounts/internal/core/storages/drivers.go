package storages

import (
	"context"

	"github.com/avalonprod/eliteeld/accounts/internal/core/models"
)

type Drivers interface {
	Create(ctx context.Context, input models.Driver) error
	GetByCredentials(ctx context.Context, email, password string) (models.Driver, error)
	IsDuplicate(ctx context.Context, email string) (bool, error)
	// SetSession(ctx context.Context, driverID string, session models.Session) error
	// GetByRefreshToken(ctx context.Context, refreshToken string) (models.Driver, error)
	// ChangePassword(ctx context.Context, userID, password, newPassword string) error
}
