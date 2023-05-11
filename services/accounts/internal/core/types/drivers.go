package types

type DriverSignUpDTO struct {
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	CompanyID string `json:"companyID" bson:"_companyID"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type DriverSignIpDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
