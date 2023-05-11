package types

type ChangePasswordDTO struct {
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}
