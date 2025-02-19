package domain

// SignUpRequest struct
type SignUpRequest struct {
	UserName        string `bson:"username" json:"username" validate:"required,min=5,max=255"`
	Email           string `bson:"email" json:"email" validate:"required"`
	Password        string `bson:"password" json:"password" validate:"required"`
	ConfirmPassword string `bson:"confirm_password" json:"confirm_password" validate:"required"`
	UserType        string `bson:"user_type" json:"user_type" validate:"required"`
}