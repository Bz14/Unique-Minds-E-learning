package domain

// ResetPassword struct
type ResetPassword struct {
	NewPassword string `json:"new_password" validate:"required"`
}