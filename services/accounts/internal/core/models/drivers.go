package models

import "time"

type Driver struct {
	ID             string    `json:"id" bson:"_id,omitempty"`
	Name           string    `json:"name" bson:"name"`
	Surname        string    `json:"surname" bson:"surname"`
	Phone          string    `json:"phone" bson:"phone"`
	CompanyID      string    `json:"companyID" bson:"_companyID"`
	Email          string    `json:"email" bson:"email"`
	Password       string    `json:"password" bson:"password"`
	RegisteredTime time.Time `json:"registeredTime" bson:"registeredTime"`
	LastVisitTime  time.Time `json:"lastVisitTime" bson:"lastVisitTime"`
	Verification   bool      `json:"verification" bson:"verification"`
}
