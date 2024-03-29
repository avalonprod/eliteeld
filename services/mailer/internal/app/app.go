package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/avalonprod/eliteeld/mailer/internal/config"
	"github.com/avalonprod/eliteeld/mailer/internal/domain/services"
	httpHandlers "github.com/avalonprod/eliteeld/mailer/internal/interfaces/api/http"
	"github.com/avalonprod/eliteeld/mailer/pkg/email/smtp"
	"github.com/avalonprod/eliteeld/mailer/pkg/logger"
)

const configsDir = "configs"
const logsFile = "logs/logs.log"

func Run() {
	logger := logger.NewLogger()
	logger.Init(logsFile)
	cfg, err := config.Init(configsDir)

	if err != nil {
		logger.Error("failed to parse config.")
		return
	}
	emailSender, err := smtp.NewSMTPSender(cfg.SMTP.From, cfg.SMTP.Password, cfg.SMTP.Host, cfg.SMTP.Port)
	if err != nil {
		logger.Errorf("error: %v", err)

		return
	}

	services := services.NewServices(&services.Deps{
		Logger:      *logger,
		EmailSender: emailSender,
		EmailConfig: cfg.Email,
	})

	handlers := httpHandlers.NewHandlers(services)

	srv := NewServer(cfg, handlers.Init(cfg))

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server %s\n", err.Error())
		}
	}()
	logger.Info("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}

}

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handlers http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + cfg.HTTP.Port,
			Handler:        handlers,
			ReadTimeout:    cfg.HTTP.ReadTimeout,
			WriteTimeout:   cfg.HTTP.WriteTimeout,
			MaxHeaderBytes: cfg.HTTP.MaxHeaderBytes << 25,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
