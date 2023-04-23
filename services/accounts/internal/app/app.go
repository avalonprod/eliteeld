package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/avalonprod/eliteeld/accounts/internal/adapters/repository"
	"github.com/avalonprod/eliteeld/accounts/internal/config"
	"github.com/avalonprod/eliteeld/accounts/internal/controller"
	"github.com/avalonprod/eliteeld/accounts/internal/domain/service"
	"github.com/avalonprod/eliteeld/accounts/pkg/db/mongodb"
	"github.com/avalonprod/eliteeld/accounts/pkg/hasher"
	"github.com/avalonprod/eliteeld/accounts/pkg/logger"
)

const configsDir = "configs"
const logsFile = "logs/logs.log"

func Run() {
	logger := logger.NewLogger()
	logger.Init(logsFile)
	cfg, err := config.Init(configsDir)
	if err != nil {
		logger.Errorf("error parse config. error: %v", err)
		return
	}

	mongoClient, err := mongodb.NewConnection(&mongodb.Config{
		URL:      cfg.Mongo.URL,
		Username: cfg.Mongo.Username,
		Password: cfg.Mongo.Password,
	})

	if err != nil {
		logger.Errorf("failed to connection mongodb. error: %v", err)
	}
	mongodb := mongoClient.Database(cfg.Mongo.Database)
	hasher := hasher.NewHasher(cfg.Password.PasswordSalt)
	repository := repository.NewRepository(mongodb)
	service := service.NewService(&service.Options{
		Repository: repository,
		Logger:     *logger,
		Hasher:     hasher,
	})
	handler := controller.NewHandler(service)
	srv := NewServer(cfg, handler.InitRoutes(cfg))

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
