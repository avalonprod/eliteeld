package mongodb

import (
	"context"
	"time"

	"github.com/avalonprod/eliteeld/accounts/internal/core/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DriversStorage struct {
	db *mongo.Collection
}

func NewDriversyStorage(db *mongo.Database) *DriversStorage {
	return &DriversStorage{
		db: db.Collection(driversCollection),
	}
}

func (s *DriversStorage) Create(ctx context.Context, input models.Driver) error {
	nCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := s.db.InsertOne(nCtx, input)
	if err != nil {
		return err
	}

	return nil
}

func (s *DriversStorage) IsDuplicate(ctx context.Context, email string) (bool, error) {
	nCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	filter := bson.M{"email": email}

	count, err := s.db.CountDocuments(nCtx, filter)
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil

}
