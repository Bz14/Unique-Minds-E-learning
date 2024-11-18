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
	tokenService    domain.TokenServiceInterface
}

func NewUserUseCase(repo domain.UserRepoInterface, dataValidator domain.ValidatorInterface, password domain.PasswordServiceInterface, tokenService domain.TokenServiceInterface) *UserUseCase {
	return &UserUseCase{
		userRepo:        repo,
		validator:       dataValidator,
		passwordService: password,
		tokenService: tokenService,
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

	_, err := u.userRepo.FindUserByEmail(user.Email)
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
	user.GoogleID = ""
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

func (u *UserUseCase) Login(loginRequest domain.LoginRequest)(domain.LoginResponse, error){
	if err := u.validator.ValidateEmail(loginRequest.Email); err != nil {
		return domain.LoginResponse{}, err
	}
	if err := u.validator.ValidatePassword(loginRequest.Password); err != nil {
		return domain.LoginResponse{}, err
	}
	user, err := u.userRepo.FindUserByEmail(loginRequest.Email)
	if err != nil {
		return domain.LoginResponse{}, errors.New("user not signed up")
	}
	if !user.IsVerified{
		return domain.LoginResponse{}, errors.New("user not verified")
	}

	if user.GoogleID != ""{
		return domain.LoginResponse{}, errors.New("user signed up with google. Try using google login")
	}
	err = u.passwordService.UnHashPassword(user.Password, loginRequest.Password)
	if err != nil{
		return domain.LoginResponse{}, errors.New("incorrect password")
	}

	accessToken, err := u.tokenService.GenerateAccessToken(user)
	if err != nil{
		return domain.LoginResponse{}, err
	}

	refreshToken, err := u.tokenService.GenerateResetToken(user)
	if err != nil{
		return domain.LoginResponse{}, err
	}
	return domain.LoginResponse{
		AccessToken: accessToken, 
		RefreshToken: refreshToken,
	},nil

}

func (u *UserUseCase) UpdateRole(email string, role string) error{
	if err := u.validator.ValidateEmail(email); err != nil {
		return err
	}
	if err := u.validator.ValidateRole(role); err != nil {
		return err
	}
	_, err := u.userRepo.FindUserByEmail(email)
	if err != nil {
		return errors.New("user not signed up")
	}
	err = u.userRepo.UpdateRole(email, role)
	if err != nil{
		return err
	}
	return nil
}

