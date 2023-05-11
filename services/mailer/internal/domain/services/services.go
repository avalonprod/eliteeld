package services

import (
	"context"

	"github.com/avalonprod/eliteeld/mailer/internal/config"
	"github.com/avalonprod/eliteeld/mailer/internal/domain/models"
	"github.com/avalonprod/eliteeld/mailer/pkg/email"
	"github.com/avalonprod/eliteeld/mailer/pkg/logger"
)

type EmailsI interface {
	SendEmailCompanyRegistration(ctx context.Context, input models.CompanyRegistrationEmailDTO) error
	SendEmailDriverRegistration(ctx context.Context, input models.DriverRegistrationEmailDTO) error
}

type Services struct {
	Emails EmailsI
}

type Deps struct {
	Logger      logger.Logger
	EmailSender email.Sender
	EmailConfig config.EmailConfig
}

func NewServices(deps *Deps) *Services {
	return &Services{
		Emails: NewEmailsService(deps.EmailSender, deps.EmailConfig, deps.Logger),
	}
}
