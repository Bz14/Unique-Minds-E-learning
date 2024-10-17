package usecases

import (
	"errors"
	"time"
	domain "unique-minds/Domain"
	utils "unique-minds/Utils"
)

type UserUseCase struct {
	userRepo        domain.UserRepoInterface
	validator       domain.ValidatorInterface
	passwordService domain.PasswordServiceInterface
}

func NewUserUseCase(repo domain.UserRepoInterface, dataValidator domain.ValidatorInterface, password domain.PasswordServiceInterface) *UserUseCase {
	return &UserUseCase{
		userRepo:        repo,
		validator:       dataValidator,
		passwordService: password,
	}
}

// SignUp implements domain.UserUseCaseInterface.
func (u *UserUseCase) SignUp(user domain.User) error {
	if err := u.validator.ValidateEmail(user.Email); err != nil {
		return err
	}
	if err := u.validator.ValidatePassword(user.Password); err != nil {
		return err
	}
	if err := u.validator.ValidateRole(user.Role); err != nil {
		return err
	}

	err := u.userRepo.FindUserByEmail(user.Email)
	if err == nil {
		return errors.New("email already exists")
	}
	password, err := u.passwordService.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = password
	user.IsVerified = false
	user.Created_at = time.Now()
	user.Updated_at = time.Now()

	err = u.userRepo.SaveUnverifiedUser(&user)
	if err != nil{
		return err
	}
	token, err := utils.GenerateResetToken()
	if err != nil{
		return err
	}
	if err := utils.SendResetPasswordEmail(user.Email, token); err != nil{
		return err
	}
	return err
}

// FindEmail implements domain.UserUseCaseInterface.
func (u *UserUseCase) FindEmail(email string) error {
	err := u.userRepo.FindUserByEmail(email)
	return err
}
