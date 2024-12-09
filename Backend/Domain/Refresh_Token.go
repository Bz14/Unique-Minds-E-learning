package domain

// RefreshTokenRequest struct
type RefreshTokenRequest struct {
	RefreshToken string `form:"refreshToken" binding:"required"`
}

// RefreshTokenResponse struct
type RefreshTokenResponse struct {
	AccessToken string `json:"accessToken"`
}
