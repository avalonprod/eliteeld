package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/avalonprod/eliteeld/accounts/internal/core/models"
	"github.com/avalonprod/eliteeld/accounts/internal/core/storages"
	"github.com/avalonprod/eliteeld/accounts/internal/core/types"
	"github.com/avalonprod/eliteeld/accounts/internal/infrastructure/emails"
	"github.com/avalonprod/eliteeld/accounts/pkg/auth"
	"github.com/avalonprod/eliteeld/accounts/pkg/hasher"
)

const (
	ErrCompanyAlreadyExists = "user with such email already exists"
	ErrCompanyNotFound      = "user doesn't exists"
)

type CompanyService struct {
	companyStorage  storages.Company
	driversService  Drivers
	hasher          hasher.PasswordHasher
	tokenManager    auth.TokenManager
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
	emailsService   *emails.Emails
}

func NewCompanyService(companyStorage storages.Company, hasher hasher.PasswordHasher, tokenManager auth.TokenManager, accessTokenTTL time.Duration, refreshTokenTTL time.Duration, emailsService *emails.Emails, driversService Drivers) *CompanyService {
	return &CompanyService{
		companyStorage:  companyStorage,
		hasher:          hasher,
		tokenManager:    tokenManager,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
		emailsService:   emailsService,
		driversService:  driversService,
	}
}

func (s *CompanyService) CompanySignUp(ctx context.Context, input types.CompanySignUpDTO) error {
	err := validateCredentials(input.Email, input.Password)
	if err != nil {
		return err
	}
	passwordHash, err := s.hasher.Hash(input.Password)
	if err != nil {
		return err
	}
	company := models.Company{
		Name:              input.Name,
		Surname:           input.Surname,
		Phone:             input.Phone,
		Usdot:             input.Usdot,
		State:             input.State,
		City:              input.City,
		TimeZone:          input.TimeZone,
		ZipCode:           input.ZipCode,
		FleetSize:         input.FleetSize,
		CarrierName:       input.CarrierName,
		MainOfficeAddress: input.MainOfficeAddress,
		Ein:               input.Ein,
		Email:             input.Email,
		Password:          passwordHash,
		RegisteredTime:    time.Now(),
		LastVisitTime:     time.Now(),
		Verification:      false,
	}
	isDuplicate, err := s.companyStorage.IsDuplicate(ctx, input.Email)
	if err != nil {
		return err

	}
	if isDuplicate {
		return errors.New(ErrCompanyAlreadyExists)
	}

	if err := s.companyStorage.Create(ctx, company); err != nil {
		return err
	}
	go func() {
		s.emailsService.SendEmailCompanyRegistration(input.Email, input.Name)
	}()
	return nil
}

func (s *CompanyService) CompanySignIn(ctx context.Context, input types.CompanySignIpDTO) (types.Tokens, error) {
	err := validateCredentials(input.Email, input.Password)
	if err != nil {
		return types.Tokens{}, err
	}
	passwordHash, err := s.hasher.Hash(input.Password)
	if err != nil {
		return types.Tokens{}, err
	}
	company, err := s.companyStorage.GetByCredentials(ctx, input.Email, passwordHash)
	if err != nil {
		if errors.Is(err, errors.New(ErrCompanyNotFound)) {
			return types.Tokens{}, err
		}
		return types.Tokens{}, err
	}
	return s.createSession(ctx, company.ID)
}

func (s *CompanyService) createSession(ctx context.Context, companyID string) (types.Tokens, error) {
	var (
		res types.Tokens
		err error
	)

	res.AccessToken, err = s.tokenManager.NewJWT(companyID, s.accessTokenTTL)
	if err != nil {
		return res, err
	}

	res.RefreshToken, err = s.tokenManager.NewRefreshToken()
	if err != nil {
		return res, err
	}

	session := models.Session{
		RefreshToken: res.RefreshToken,
		ExpiresAt:    time.Now().Add(s.refreshTokenTTL),
	}

	err = s.companyStorage.SetSession(ctx, companyID, session)

	return res, err
}

func (s *CompanyService) RefreshTokens(ctx context.Context, refreshToken string) (types.Tokens, error) {
	company, err := s.companyStorage.GetByRefreshToken(ctx, refreshToken)
	if err != nil {
		return types.Tokens{}, err
	}

	return s.createSession(ctx, company.ID)
}

func (s *CompanyService) CompanyChangePassword(ctx context.Context, userID string, input types.ChangePasswordDTO) error {
	err := validatePassword(input.Password)
	if err != nil {
		return err
	}
	err = validatePassword(input.NewPassword)
	if err != nil {
		return err
	}
	if strings.Trim(input.Password, " ") == strings.Trim(input.NewPassword, " ") {
		return errors.New("the new password cannot be the same as the old password")
	}
	passwordHash, err := s.hasher.Hash(input.Password)
	if err != nil {
		return err
	}
	newPasswordHash, err := s.hasher.Hash(input.NewPassword)
	if err != nil {
		return err
	}
	err = s.companyStorage.ChangePassword(ctx, userID, passwordHash, newPasswordHash)
	if err != nil {
		if errors.Is(err, errors.New(ErrCompanyNotFound)) {
			return err
		}
		return err
	}
	return nil
}

func (s *CompanyService) CompanyForgotPassword(ctx context.Context, email string) {

}

func (s *CompanyService) CompanyCreateNewDriver(ctx context.Context, input types.DriverSignUpDTO, companyID string) error {
	company, err := s.companyStorage.GetByCompanyID(ctx, companyID)
	if err != nil {
		if errors.Is(err, errors.New(ErrCompanyNotFound)) {
			return err
		}
		return err
	}

	err = s.driversService.DriversSignUp(ctx, input, company.CarrierName, company.Email)
	if err != nil {
		return err
	}
	return nil
}
