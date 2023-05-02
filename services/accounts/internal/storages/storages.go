package storages

import (
	"github.com/avalonprod/eliteeld/accounts/internal/core/storages"
	"github.com/avalonprod/eliteeld/accounts/internal/storages/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

type Storages struct {
	Company storages.Company
}

func NewStorages(db *mongo.Database) *Storages {
	return &Storages{
		Company: mongodb.NewCompanyStorage(db),
	}
}
