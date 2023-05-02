package mongodb

import (
	"context"
	"errors"
	"time"

	"github.com/avalonprod/eliteeld/accounts/internal/core/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const ErrUserNotFound = "user doesn't exists"

type CompanyStorage struct {
	db *mongo.Collection
}

func NewCompanyStorage(db *mongo.Database) *CompanyStorage {
	return &CompanyStorage{
		db: db.Collection(companyCollection),
	}
}

func (s *CompanyStorage) Create(ctx context.Context, input models.Company) error {
	nCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := s.db.InsertOne(nCtx, input)
	if err != nil {
		return err
	}

	return nil
}

func (s *CompanyStorage) GetByCredentials(ctx context.Context, email, password string) (models.Company, error) {
	nCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var company models.Company
	filter := bson.M{"email": email, "password": password}

	res := s.db.FindOne(nCtx, filter)

	err := res.Err()
	if err != nil {
		return company, err
	}
	if err := res.Decode(&company); err != nil {
		return company, err
	}

	return company, err
}

func (s *CompanyStorage) IsDuplicate(ctx context.Context, email string) (bool, error) {
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

func (s *CompanyStorage) SetSession(ctx context.Context, companyID string, session models.Session) error {
	nCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	ObjectID, err := primitive.ObjectIDFromHex(companyID)
	if err != nil {
		return err
	}
	_, err = s.db.UpdateOne(nCtx, bson.M{"_id": ObjectID}, bson.M{"$set": bson.M{"session": session, "lastVisitTime": time.Now()}})

	return err
}

func (s *CompanyStorage) GetByRefreshToken(ctx context.Context, refreshToken string) (models.Company, error) {
	var company models.Company
	if err := s.db.FindOne(ctx, bson.M{
		"session.refreshToken": refreshToken,
		"session.expiresAt":    bson.M{"$gt": time.Now()},
	}).Decode(&company); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Company{}, errors.New(ErrUserNotFound)
		}

		return models.Company{}, err
	}

	return company, nil
}
