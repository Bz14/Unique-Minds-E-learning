package utils

type PasswordService struct{}

type PasswordServiceInterface interface {
	HashPassword(password string) (string, error)
	UnHashPassword(hashedPassword string, password string) error
}

func NewPasswordService() *PasswordService {
	return &PasswordService{}
}

func (ps *PasswordService) HashPassword(password string) (string, error) {
	return "", nil
}

func (ps *PasswordService) UnHashPassword(hashedPassword string, password string) error {
	return nil
}