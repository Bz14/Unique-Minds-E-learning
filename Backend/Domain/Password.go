package domain

type PasswordServiceInterface interface {
	HashPassword(password string) (string, error)
	UnHashPassword(hashedPassword string, password string) error
}