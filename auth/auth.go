package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var SECRET = []byte("super-secret-auth-key")

type JWTClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(20 * time.Minute)
	claims := &JWTClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SECRET)
}

func TokenValid(ctx *gin.Context) error {
	bearer := ctx.Request.Header.Get("Authorization")
	if bearer != "" {
		jwtParts := strings.Split(bearer, " ")
		if len(jwtParts) == 2 {
			jwtEncoded := jwtParts[1]
			token, err := jwt.Parse(jwtEncoded, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signin method %v", token.Header["alg"])
				}
				return SECRET, nil
			})
			if err != nil {
				return err
			}
			_, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				ctx.JSON(http.StatusInternalServerError, gin.H{"message": "unable to parse claims"})
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("token not found")

}
