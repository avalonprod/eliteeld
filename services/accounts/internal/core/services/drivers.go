package services

import (
	"context"
	"errors"
	"time"

	"github.com/avalonprod/eliteeld/accounts/internal/core/models"
	"github.com/avalonprod/eliteeld/accounts/internal/core/storages"
	"github.com/avalonprod/eliteeld/accounts/internal/core/types"
	"github.com/avalonprod/eliteeld/accounts/internal/infrastructure/emails"
	"github.com/avalonprod/eliteeld/accounts/pkg/hasher"
)

const (
	ErrDriverAlreadyExists = "user with such email already exists"
	ErrDriverNotFound      = "user doesn't exists"
)

type DriversService struct {
	driversStorage storages.Drivers
	hasher         hasher.PasswordHasher
	emailsService  *emails.Emails
}

func NewDriversService(driversStorage storages.Drivers, hasher hasher.PasswordHasher, emailsService *emails.Emails) *DriversService {
	return &DriversService{
		driversStorage: driversStorage,
		hasher:         hasher,
		emailsService:  emailsService,
	}
}

func (s *DriversService) DriversSignUp(ctx context.Context, input types.DriverSignUpDTO, companyName string, companyEmail string) error {
	err := validateCredentials(input.Email, input.Password)
	if err != nil {
		return err
	}
	passwordHash, err := s.hasher.Hash(input.Password)
	if err != nil {
		return err
	}
	driver := models.Driver{
		Name:           input.Name,
		Surname:        input.Surname,
		Phone:          input.Phone,
		Email:          input.Email,
		Password:       passwordHash,
		RegisteredTime: time.Now(),
		LastVisitTime:  time.Now(),
		Verification:   false,
	}
	isDuplicate, err := s.driversStorage.IsDuplicate(ctx, input.Email)
	if err != nil {
		return err

	}
	if isDuplicate {
		return errors.New(ErrDriverAlreadyExists)
	}

	if err := s.driversStorage.Create(ctx, driver); err != nil {
		return err
	}
	go func() {
		s.emailsService.SendEmailDriverRegistration(input.Email, input.Name, companyName, companyEmail)
	}()
	return nil
}

// func (s *CompanyService) CompanySignIn(ctx context.Context, input types.CompanySignIpDTO) (types.Tokens, error) {
// 	err := validateCredentials(input.Email, input.Password)
// 	if err != nil {
// 		return types.Tokens{}, err
// 	}
// 	passwordHash, err := s.hasher.Hash(input.Password)
// 	if err != nil {
// 		return types.Tokens{}, err
// 	}
// 	company, err := s.companyStorage.GetByCredentials(ctx, input.Email, passwordHash)
// 	if err != nil {
// 		if errors.Is(err, errors.New(ErrCompanyNotFound)) {
// 			return types.Tokens{}, err
// 		}
// 		return types.Tokens{}, err
// 	}
// 	return s.createSession(ctx, company.ID)
// }

// func (s *CompanyService) createSession(ctx context.Context, companyID string) (types.Tokens, error) {
// 	var (
// 		res types.Tokens
// 		err error
// 	)

// 	res.AccessToken, err = s.tokenManager.NewJWT(companyID, s.accessTokenTTL)
// 	if err != nil {
// 		return res, err
// 	}

// 	res.RefreshToken, err = s.tokenManager.NewRefreshToken()
// 	if err != nil {
// 		return res, err
// 	}

// 	session := models.Session{
// 		RefreshToken: res.RefreshToken,
// 		ExpiresAt:    time.Now().Add(s.refreshTokenTTL),
// 	}

// 	err = s.companyStorage.SetSession(ctx, companyID, session)

// 	return res, err
// }
