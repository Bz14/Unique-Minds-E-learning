package domain

type SignUpRequest struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"userType"`
}