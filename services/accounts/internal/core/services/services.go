package services

import (
	"context"
	"time"

	"github.com/avalonprod/eliteeld/accounts/internal/core/storages"
	"github.com/avalonprod/eliteeld/accounts/internal/core/types"
	"github.com/avalonprod/eliteeld/accounts/internal/infrastructure/emails"
	"github.com/avalonprod/eliteeld/accounts/pkg/auth"
	"github.com/avalonprod/eliteeld/accounts/pkg/hasher"
	"github.com/avalonprod/eliteeld/accounts/pkg/logger"
)

type Company interface {
	CompanySignUp(ctx context.Context, input types.CompanySignUpDTO) error
	CompanySignIn(ctx context.Context, input types.CompanySignIpDTO) (types.Tokens, error)
	RefreshTokens(ctx context.Context, refreshToken string) (types.Tokens, error)
}

type Deps struct {
	CompanyStorage  storages.Company
	Hasher          hasher.PasswordHasher
	Logger          *logger.Logger
	TokenManager    auth.TokenManager
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
	EmailsService   *emails.Emails
}

type Services struct {
	Company Company
}

func NewServices(deps *Deps) *Services {
	companyService := NewCompanyService(deps.CompanyStorage, deps.Hasher, deps.TokenManager, deps.AccessTokenTTL, deps.RefreshTokenTTL, deps.EmailsService)
	return &Services{
		Company: companyService,
	}
}
