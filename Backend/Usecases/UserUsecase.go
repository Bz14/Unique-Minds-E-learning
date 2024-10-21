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
func (u *UserUseCase) SignUp(user domain.User) (bool, error) {
	if err := u.validator.ValidateEmail(user.Email); err != nil {
		return false, err
	}
	if err := u.validator.ValidatePassword(user.Password); err != nil {
		return false, err
	}
	if err := u.validator.ValidateRole(user.Role); err != nil {
		return false, err
	}

	err := u.userRepo.FindUserByEmail(user.Email)
	if err == nil {
		return false, errors.New("email already exists")
	}
	password, err := u.passwordService.HashPassword(user.Password)
	if err != nil {
		return false, err
	}
	token, err := utils.GenerateResetToken()
	if err != nil{
		return false, err
	}
	
	unverifiedUser, err := u.userRepo.FindUnverifiedUserByEmail(user.Email)
	if err == nil{
		if time.Now().After(unverifiedUser.VerificationTokenExpire) {
			err := u.userRepo.UpdateUnverifiedUser(unverifiedUser.Email, time.Now(), token, time.Now().Add(time.Hour * 10))
			if err != nil{
				return false, err
			}
			if err := utils.SendVerificationEmail(unverifiedUser.FullName,unverifiedUser.Email, token); err != nil{
				return false, err
			}
		}
		return true, nil
	}

	user.Password = password
	user.IsVerified = false
	user.Created_at = time.Now()
	user.Updated_at = time.Now()
	user.VerificationToken = token
	user.VerificationTokenExpire = time.Now().Add(time.Hour * 10)
	err = u.userRepo.SaveUnverifiedUser(&user)
	if err != nil{
		return false, err
	}
	
	if err := utils.SendVerificationEmail(user.FullName, user.Email, token); err != nil{
		return false, err
	}
	return false, err
}

// FindEmail implements domain.UserUseCaseInterface.
func (u *UserUseCase) FindEmail(email string) error {
	err := u.userRepo.FindUserByEmail(email)
	return err
}

func (u *UserUseCase) VerifyEmail(token string) error{
	user, err := u.userRepo.FindUserByToken(token)
	if err != nil{
		return err
	}
	if time.Now().After(user.VerificationTokenExpire){
		return errors.New("verification token has expired")
	}
	user.IsVerified = true
	user.Created_at = time.Now()
	user.Updated_at = time.Now()
	user.VerificationToken = ""
	user.VerificationTokenExpire = time.Now()
	err = u.userRepo.SignUpUser(user)
	if err != nil{
		return err
	}
	return nil
}


