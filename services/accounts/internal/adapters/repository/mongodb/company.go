package mongodb

import (
	"context"
	"time"

	"github.com/avalonprod/eliteeld/accounts/internal/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		db: db.Collection(userCollection),
	}
}

func (u *UserRepository) Create(ctx context.Context, input model.User) error {
	nCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := u.db.InsertOne(nCtx, input)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	nCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user model.User
	filter := bson.M{"email": email}

	res := u.db.FindOne(nCtx, filter)

	err := res.Err()
	if err != nil {
		return user, err
	}
	if err := res.Decode(&user); err != nil {
		return user, err
	}

	return user, err
}

func (u *UserRepository) GetUserById(ctx context.Context, id string) (model.User, error) {
	nCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user model.User

	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return user, err
	}

	filter := bson.M{"_id": objectID}

	res := u.db.FindOne(nCtx, filter)

	err = res.Err()
	if err != nil {
		return user, err
	}
	if err := res.Decode(&user); err != nil {
		return user, err
	}

	return user, err
}

func (u *UserRepository) IsDuplicate(ctx context.Context, email string) (bool, error) {
	nCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	filter := bson.M{"email": email}

	count, err := u.db.CountDocuments(nCtx, filter)
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil

}
