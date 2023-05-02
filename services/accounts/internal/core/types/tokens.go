package types

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type RefreshToken struct {
	Token string `json:"token" binding:"required"`
}
