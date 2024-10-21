package controller

import (
	"net/http"
	domain "unique-minds/Domain"

	"github.com/gin-gonic/gin"
)

type OauthController struct {
	oauthUseCase domain.OauthUseCaseInterface
}

type OauthControllerInterface interface{
	GoogleAuth() (string, error)
	GoogleCallback() (string, error)
}

func NewOauthController(useCase domain.OauthUseCaseInterface) *OauthController {
	return &OauthController{
		oauthUseCase: useCase,
	}
}


func (oc *OauthController) GoogleAuth(ctx *gin.Context){
	URL, isSuccess := oc.oauthUseCase.GoogleAuth()
	if isSuccess {
		url, ok := URL.(string)
		if !ok {
			ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{
				Message: "Invalid URL type",
				Status: http.StatusInternalServerError,
			})
			return
		}
		ctx.Redirect(http.StatusTemporaryRedirect, url)
	}else{
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "Failed to get URL",
			Status: http.StatusInternalServerError,
		})
		return
	}

}

func (oc *OauthController) GoogleCallback() (string, error) {}