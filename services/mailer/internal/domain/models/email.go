package models

type CompanyRegistrationEmail struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type CompanyRegistrationEmailDTO struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type DriverRegistrationEmail struct {
	Email        string `json:"email"`
	Name         string `json:"name"`
	CompanyName  string `json:"companyName"`
	CompanyEmail string `json:"companyEmail"`
}

type DriverRegistrationEmailDTO struct {
	Email        string `json:"email"`
	Name         string `json:"name"`
	CompanyName  string `json:"companyName"`
	CompanyEmail string `json:"companyEmail"`
}
