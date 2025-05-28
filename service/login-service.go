package service

type LoginService interface {
	Login(username, password string) bool
}

type loginService struct {
	authorizedUsername string
	authorizedPassword string
}

// Login implements LoginService.
func (service *loginService) Login(username string, password string) bool {
	return service.authorizedUsername == username && service.authorizedPassword == password
}

func NewLoginService() LoginService {
	return &loginService{
		authorizedUsername: "admin",
		authorizedPassword: "123456",
	}
}
