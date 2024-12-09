package domain

// ResetPasswordRequest struct
type ResetPasswordRequest struct {
	Email string `json:"email" binding:"required"`
}