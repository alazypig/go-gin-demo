package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(name string, admin bool) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

// GenerateToken implements JWTService.
func (service *jwtService) GenerateToken(username string, admin bool) string {
	claims := &jwtCustomClaims{
		username,
		admin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(service.secretKey))

	if err != nil {
		panic(err)
	}

	return t
}

// ValidateToken implements JWTService.
func (service *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(service.secretKey), nil
	})
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "edward-secret-key"
	}

	return secret
}

func NewJWTService() JWTService {
	return &jwtService{secretKey: getSecretKey(), issuer: "edward@qq.com"}
}
