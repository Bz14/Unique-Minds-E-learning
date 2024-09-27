package domain

type ValidatorInterface interface {
	ValidateEmail(email string) error
	ValidatePassword(password string) error
	ValidateRole(role string) error
}