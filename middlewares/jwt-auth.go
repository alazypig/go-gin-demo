package middlewares

import (
	"log"
	"net/http"

	"gilab.com/pragmaticreviews/golang-gin-poc/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]

		token, err := service.NewJWTService().ValidateToken(tokenString)

		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[Name]", claims["name"])
			log.Println("Claims[Admin]", claims["admin"])
			log.Println("Claims[Issuer]", claims["issuer"])
			log.Println("Claims[IssuedAt]", claims["iat"])
			log.Println("Claims[ExpiresAt]", claims["exp"])
		} else {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
