package controller

import (
	"fmt"
	"net/http"
	"strings"
	domain "unique-minds/Domain"
	infrastructures "unique-minds/Infrastructures"

	"github.com/gin-gonic/gin"
)

type OauthController struct {
	oauthUseCase domain.OauthUseCaseInterface
	config     infrastructures.Config
}

type OauthControllerInterface interface{
	GoogleAuth() (string, error)
	GoogleCallback() (string, error)
}

func NewOauthController(useCase domain.OauthUseCaseInterface, config infrastructures.Config) *OauthController {
	return &OauthController{
		oauthUseCase: useCase,
		config:     config,
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
		return
	}else{
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "Failed to get URL",
			Status: http.StatusInternalServerError,
		})
		return
	}

}

func (oc *OauthController) GoogleCallback(ctx *gin.Context){
	state := ctx.Query("state")

	if strings.ToLower(state) != oc.config.State {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Invalid oauth state",
			Status: http.StatusBadRequest,
		})
		return
	}

	code := ctx.Query("code")
	if code == "" {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Code not found",
			Status: http.StatusBadRequest,
		})
		return
	}
	user, err , isNew := oc.oauthUseCase.GoogleCallback(code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: err.Message,
			Status: http.StatusInternalServerError,
		})
		return
	}
	if isNew{
		redirectUrl := fmt.Sprintf("%s?email=%s", oc.config.RoleRedirect,user.Email)
		ctx.Redirect(http.StatusTemporaryRedirect, redirectUrl)
		return
	}
	ctx.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "User already exists",
		Status: http.StatusOK,
	})
}
