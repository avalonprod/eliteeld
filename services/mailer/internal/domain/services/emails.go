package services

import (
	"context"
	"fmt"

	"github.com/avalonprod/eliteeld/mailer/internal/config"
	"github.com/avalonprod/eliteeld/mailer/internal/domain/models"
	"github.com/avalonprod/eliteeld/mailer/pkg/email"
	"github.com/avalonprod/eliteeld/mailer/pkg/logger"
)

type EmailsService struct {
	sender email.Sender
	config config.EmailConfig
	logger logger.Logger
}

func NewEmailsService(sender email.Sender, config config.EmailConfig, logger logger.Logger) *EmailsService {
	return &EmailsService{
		sender: sender,
		config: config,
		logger: logger,
	}
}

func (s *EmailsService) SendEmailCompanyRegistration(ctx context.Context, input models.CompanyRegistrationEmailDTO) error {
	subject := fmt.Sprintf(s.config.Subjects.CompanyRegistration)
	sendInput := email.SendEmailInput{To: input.Email, Subject: subject}
	templateInput := models.CompanyRegistrationEmail{Email: input.Email, Name: input.Name}

	err := sendInput.GenerateBodyFromHTML(s.config.Templates.CompanyRegistrationTemplate, templateInput)
	if err != nil {
		return err
	}

	err = s.sender.Send(sendInput)
	if err != nil {
		s.logger.Debugf("failed to send message from email: %s, error: %v", input.Email, err)
		return err
	}
	return nil
}

func (s *EmailsService) SendEmailDriverRegistration(ctx context.Context, input models.DriverRegistrationEmailDTO) error {
	subject := fmt.Sprintf(s.config.Subjects.DriverRegistration, input.CompanyName)
	sendInput := email.SendEmailInput{To: input.Email, Subject: subject}
	templateInput := models.DriverRegistrationEmail{Email: input.Email, Name: input.Name, CompanyName: input.CompanyName, CompanyEmail: input.CompanyEmail}

	err := sendInput.GenerateBodyFromHTML(s.config.Templates.DriverRegistrationTemplate, templateInput)
	if err != nil {
		return err
	}

	err = s.sender.Send(sendInput)
	if err != nil {
		s.logger.Debugf("failed to send message from email: %s, error: %v", input.Email, err)
		return err
	}
	return nil
}
