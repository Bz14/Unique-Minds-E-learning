package usecases

import (
	"errors"
	domain "unique-minds/Domain"
)

type UserUseCase struct {
	userRepo domain.UserRepoInterface
	validator domain.ValidatorInterface
	passwordService domain.PasswordServiceInterface
}

func NewUserUseCase(repo domain.UserRepoInterface, dataValidator domain.ValidatorInterface, password domain.PasswordServiceInterface) *UserUseCase {
	return &UserUseCase{
		userRepo : repo,
		validator: dataValidator,
		passwordService: password,
	}
}


// SignUp implements domain.UserUseCaseInterface.
func (u *UserUseCase) SignUp(signUpRequest domain.SignUpRequest) error {
	if err := u.validator.ValidateEmail(signUpRequest.Email); err != nil{
		return err
	}
	if err := u.validator.ValidatePassword(signUpRequest.Password); err != nil{
		return err
	}
	if err := u.validator.ValidateRole(signUpRequest.Role); err != nil{
		return err	
	}

	err := u.userRepo.FindUserByEmail(signUpRequest.Email)
	if err == nil{
		return errors.New("email already exists")
	}
	password, err := 
	// hash password
	// create user
	return nil
}