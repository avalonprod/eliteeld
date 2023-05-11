package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/avalonprod/eliteeld/accounts/internal/config"
	"github.com/avalonprod/eliteeld/accounts/internal/core/services"
	"github.com/avalonprod/eliteeld/accounts/internal/infrastructure/emails"
	"github.com/avalonprod/eliteeld/accounts/internal/interfaces/api/rest"
	"github.com/avalonprod/eliteeld/accounts/internal/storages"
	"github.com/avalonprod/eliteeld/accounts/pkg/auth"
	"github.com/avalonprod/eliteeld/accounts/pkg/db/mongodb"
	"github.com/avalonprod/eliteeld/accounts/pkg/hasher"
	"github.com/avalonprod/eliteeld/accounts/pkg/logger"
)

func Run() {

	logger := logger.NewLogger()
	logger.Init("logs/logs.log")

	cfg, err := config.Init("configs")
	if err != nil {
		logger.Error(err)
		return
	}
	mongoClient, err := mongodb.NewConnection(&mongodb.Config{
		URL:      cfg.Mongo.URL,
		Username: cfg.Mongo.Username,
		Password: cfg.Mongo.Password,
	})
	if err != nil {
		logger.Errorf("failed to create new mongo client. error: %v", err)
		return
	}
	mongodb := mongoClient.Database(cfg.Mongo.Database)
	tokenManager, err := auth.NewManager(cfg.Auth.JWT.SigningKey)
	if err != nil {
		logger.Error(err)
		return
	}
	emails := emails.NewEmails(cfg.Emails.ApiUrlCompanyRegistration, cfg.Emails.ApiUrlDriverRegistration)
	hasher := hasher.NewHasher(cfg.Auth.PasswordSalt)
	storages := storages.NewStorages(mongodb)
	services := services.NewServices(&services.Deps{
		CompanyStorage:  storages.Company,
		Hasher:          hasher,
		Logger:          logger,
		TokenManager:    tokenManager,
		AccessTokenTTL:  cfg.Auth.JWT.AccessTokenTTL,
		RefreshTokenTTL: cfg.Auth.JWT.RefreshTokenTTL,
		EmailsService:   emails,
	})
	restHandlers := rest.NewHandler(services, tokenManager)

	// Starting server
	srv := NewServer(cfg, restHandlers.Init(cfg))

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server %s\n", err.Error())
			return
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
	if err := mongoClient.Disconnect(context.Background()); err != nil {
		logger.Error(err.Error())
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
