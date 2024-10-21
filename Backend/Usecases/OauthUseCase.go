package usecases

import (
	domain "unique-minds/Domain"
	infrastructures "unique-minds/Infrastructures"
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