package service

import (
	"context"

	"github.com/avalonprod/eliteeld/accounts/internal/adapters/repository"
	"github.com/avalonprod/eliteeld/accounts/internal/domain/model"
	"github.com/avalonprod/eliteeld/accounts/pkg/hasher"
	"github.com/avalonprod/eliteeld/accounts/pkg/logger"
)

type User interface {
	UserLoginEmail(ctx context.Context, input model.LoginEmailUserInput) (string, error)
	UserLoginPassword(ctx context.Context, input model.LoginPasswordUserInput) (model.UserPayload, error)
	UserRegister(ctx context.Context, input model.RegisterUserInput) error
}

type Service struct {
	User User
}

type Options struct {
	Repository *repository.Repository
	Logger     logger.Logger
	Hasher     hasher.PasswordHasher
}

func NewService(options *Options) *Service {
	userService := NewUserService(options.Repository.User, options.Logger, options.Hasher)
	return &Service{
		User: userService,
	}
}
