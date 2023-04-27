package service

import (
	"context"
	"errors"
	"regexp"
	"time"

	"github.com/avalonprod/eliteeld/accounts/internal/adapters/emails"
	"github.com/avalonprod/eliteeld/accounts/internal/adapters/repository"
	apperrors "github.com/avalonprod/eliteeld/accounts/internal/appErrors"
	"github.com/avalonprod/eliteeld/accounts/internal/domain/model"
	"github.com/avalonprod/eliteeld/accounts/pkg/hasher"
	"github.com/avalonprod/eliteeld/accounts/pkg/logger"
)

type UserService struct {
	repository repository.User
	logger     logger.Logger
	hasher     hasher.PasswordHasher
	emails     emails.Emails
}

func NewUserService(repository repository.User, logger logger.Logger, hasher hasher.PasswordHasher, emails emails.Emails) *UserService {
	return &UserService{
		repository: repository,
		logger:     logger,
		hasher:     hasher,
		emails:     emails,
	}
}

func (u *UserService) UserLoginEmail(ctx context.Context, input model.LoginEmailUserInput) (string, error) {
	err := emailValidate(input.Email)
	if err != nil {
		return "", err
	}
	user, err := u.repository.GetUserByEmail(ctx, input.Email)
	if err != nil {
		if errors.Is(err, apperrors.ErrorUserNotFound) {
			return "", err
		}
		return "", err
	}

	return user.Email, nil
}

func (u *UserService) UserLoginPassword(ctx context.Context, input model.LoginPasswordUserInput) (model.UserPayload, error) {
	err := passwordValidate(input.Password)
	if err != nil {
		return model.UserPayload{}, err
	}

	passwordHash, err := u.hasher.Hash(input.Password)
	if err != nil {
		u.logger.Errorf("failed to hash password. error: %v", err)
		return model.UserPayload{}, err
	}
	user, err := u.repository.GetUserByEmail(ctx, input.Email)
	if err != nil {
		u.logger.Errorf("failde to get user by email. error: %v", err)
		return model.UserPayload{}, apperrors.ErrorUserNotFound
	}
	if user.Password != passwordHash {
		return model.UserPayload{}, errors.New("Incorrect password")
	}
	payload := model.UserPayload{
		ID:           user.ID,
		Email:        user.Email,
		Verification: user.Verification,
	}
	return payload, nil
}

func (u *UserService) UserRegister(ctx context.Context, input model.RegisterUserInput) error {
	if err := emailValidate(input.Email); err != nil {

		return err
	}
	if err := passwordValidate(input.Password); err != nil {
		return err
	}

	isDuplicate, err := u.repository.IsDuplicate(ctx, input.Email)
	if err != nil {
		u.logger.Errorf("failed to check user is duplicate. error: %v", err)
		return apperrors.ErrorUserNotFound
	}
	if isDuplicate {
		u.logger.Info("user alredy exist")
		return apperrors.ErrUserAlreadyExists
	}

	passwordHash, err := u.hasher.Hash(input.Password)
	if err != nil {
		u.logger.Errorf("failed to hash password. error: %v", err)
		return err
	}
	user := model.User{
		Name:           input.Name,
		Surname:        input.Surname,
		Usdot:          32,
		Email:          input.Email,
		Password:       passwordHash,
		RegisteredTime: time.Now(),
		LastVisitTime:  time.Now(),
		Verification:   false,
	}

	err = u.repository.Create(ctx, user)
	if err != nil {
		u.logger.Errorf("failed to create user. error: %v", err)
		return err
	}
	u.emails.SendEmailCompanyRegistration(input.Email, input.Name)
	return nil
}

func emailValidate(email string) error {
	if email == "" {
		return errors.New("email cannot be empty")
	}
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("incorrect email format")
	}
	return nil
}

func passwordValidate(password string) error {
	if password == "" {
		return errors.New("password cannot be empty")
	}

	if len(password) < 8 || !regexp.MustCompile(`[A-Z]+`).MatchString(password) || !regexp.MustCompile(`\d+`).MatchString(password) {
		return errors.New("Password must be at least 8 characters long and contain at least one uppercase letter and one digit")
	}
	return nil
}
