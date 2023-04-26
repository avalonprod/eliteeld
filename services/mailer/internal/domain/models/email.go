package models

type CompanyRegistrationEmail struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type CompanyRegistrationEmailDTO struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
