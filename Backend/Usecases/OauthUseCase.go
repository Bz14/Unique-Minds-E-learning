package usecases

import (
	"context"
	"time"
	domain "unique-minds/Domain"
	infrastructures "unique-minds/Infrastructures"

	oauth2Service "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

type OauthUseCase struct {
	userRepo     domain.UserRepoInterface
	config       infrastructures.Config
	oauthService domain.OauthConfigInterface
}


func NewOauthUseCase(repo domain.UserRepoInterface, config infrastructures.Config, oauth domain.OauthConfigInterface) *OauthUseCase {
	return &OauthUseCase{
		userRepo:     repo,
		config:       config,
		oauthService: oauth,
	}
}

func (o *OauthUseCase) GoogleAuth() (interface{}, bool){
	credentials, err := o.oauthService.InitialConfig()
	if err != nil{
		return domain.ErrorResponse{
			Message: err.Error(),
			Status: 500,
		}, false
	}
	url := credentials.AuthCodeURL("state")
	return url, true
}

func (o *OauthUseCase) GoogleCallback(code string) (*domain.ErrorResponse, bool){
	credentials, err := o.oauthService.InitialConfig()
	if err != nil{
		return &domain.ErrorResponse{
			Message: err.Error(),
			Status: 500,
		}, false
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	token, err := credentials.Exchange(ctx, code)
	if err != nil {
		return &domain.ErrorResponse{
			Message: err.Error(),
			Status: 500,
		}, false
	}
	client := credentials.Client(ctx, token)
	oauth2Service, err :=  oauth2Service.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return &domain.ErrorResponse{
			Message: err.Error(),
			Status: 500,
		}, false
	}
	userinfo, err := oauth2Service.Userinfo.V2.Me.Get().Do()
	if err != nil {
		return &domain.ErrorResponse{
			Message: err.Error(),
			Status: 500,
		}, false
	}
	existing_user, err := o.userRepo.FindUserByEmail(userinfo.Email)
	if err != nil {
		var newUser domain.User
		newUser.Email = userinfo.Email
		newUser.FullName = userinfo.Name
		newUser.Created_at = time.Now()
		newUser.Updated_at = time.Now()
		newUser.IsVerified = true
		newUser.GoogleID = userinfo.Id

		err = o.userRepo.CreateUser(&newUser)
		if err != nil {
			return &domain.ErrorResponse{
				Message: err.Error(),
				Status: 500,
			}, false
		}
		return nil, true
	}
	if existing_user.GoogleID == ""{
		return &domain.ErrorResponse{
			Message: "User not signed up with google",
			Status: 200,
		}, false
	}

	return nil, false
}