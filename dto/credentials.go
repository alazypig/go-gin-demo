package dto

type Credentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type JWT struct {
	Token string `json:"token"`
}
