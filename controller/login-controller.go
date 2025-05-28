package controller

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/dto"
	"gilab.com/pragmaticreviews/golang-gin-poc/service"
	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jwtService   service.JWTService
}

// Login implements LoginController.
func (controller *loginController) Login(ctx *gin.Context) string {
	var credentials dto.Credentials
	err := ctx.ShouldBindJSON(&credentials) // 解析 json

	if err != nil {
		return ""
	}

	isAuthenticated := controller.loginService.Login(credentials.Username, credentials.Password)

	if isAuthenticated {
		return controller.jwtService.GenerateToken(credentials.Username, true)
	}

	return ""
}

func NewLoginController(loginService service.LoginService, jwtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}
