package repository

import (
	"context"

	"github.com/avalonprod/eliteeld/accounts/internal/adapters/repository/mongodb"
	"github.com/avalonprod/eliteeld/accounts/internal/domain/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type User interface {
	Create(ctx context.Context, input model.User) error
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	GetUserById(ctx context.Context, id string) (model.User, error)
	IsDuplicate(ctx context.Context, email string) (bool, error)
}

type Repository struct {
	User User
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		User: mongodb.NewUserRepository(db),
	}
}
