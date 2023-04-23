package model

import "time"

type User struct {
	ID                string    `json:"id" bson:"_id,omitempty"`
	Name              string    `json:"name" bson:"name"`
	Surname           string    `json:"surname" bson:"surname"`
	Phone             string    `json:"phone" bson:"phone"`
	Usdot             int       `json:"usdot" bson:"usdot"`
	State             string    `json:"state" bson:"state"`
	City              string    `json:"city" bson:"city"`
	TimeZone          string    `json:"timeZone" bson:"timeZone"`
	ZipCode           int       `json:"zipCode" bson:"zipCode"`
	FleetSize         string    `json:"fleetSize" bson:"fleetSize"`
	CarrierName       string    `json:"carrierName" bson:"carrierName"`
	MainOfficeAddress string    `json:"mainOfficeAddress" bson:"mainOfficeAddress"`
	Ein               int       `json:"ein" bson:"ein"`
	Email             string    `json:"email" bson:"email"`
	Password          string    `json:"password" bson:"password"`
	RegisteredTime    time.Time `json:"registeredTime" bson:"registeredTime"`
	LastVisitTime     time.Time `json:"lastVisitTime" bson:"lastVisitTime"`
	Verification      bool      `json:"verification" bson:"verification"`
}

type LoginEmailUserInput struct {
	Email string `json:"email"`
}

type LoginPasswordUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserInput struct {
	Name              string `json:"name" bson:"name"`
	Surname           string `json:"surname" bson:"surname"`
	Phone             string `json:"phone" bson:"phone"`
	Usdot             int    `json:"usdot" bson:"usdot"`
	State             string `json:"state" bson:"state"`
	City              string `json:"city" bson:"city"`
	TimeZone          string `json:"timeZone" bson:"timeZone"`
	ZipCode           int    `json:"zipCode" bson:"zipCode"`
	FleetSize         string `json:"fleetSize" bson:"fleetSize"`
	CarrierName       string `json:"carrierName" bson:"carrierName"`
	MainOfficeAddress string `json:"mainOfficeAddress" bson:"mainOfficeAddress"`
	Ein               int    `json:"ein" bson:"ein"`
	Email             string `json:"email" bson:"email"`
	Password          string `json:"password" bson:"password"`
}

type UserPayload struct {
	Email        string `json:"email"`
	ID           string `json:"id" bson:"_id,omitempty"`
	Verification bool   `json:"verification" bson:"verification"`
}
