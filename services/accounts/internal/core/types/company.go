package types

type CompanySignUpDTO struct {
	Name              string `json:"name"`
	Surname           string `json:"surname"`
	Phone             string `json:"phone"`
	Usdot             int    `json:"usdot"`
	State             string `json:"state"`
	City              string `json:"city"`
	TimeZone          string `json:"timeZone"`
	ZipCode           int    `json:"zipCode"`
	FleetSize         string `json:"fleetSize"`
	CarrierName       string `json:"carrierName"`
	MainOfficeAddress string `json:"mainOfficeAddress"`
	Ein               int    `json:"ein"`
	Email             string `json:"email"`
	Password          string `json:"password"`
}

type CompanySignIpDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
