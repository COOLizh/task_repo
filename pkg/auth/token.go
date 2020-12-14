package auth

import (
	"time"

	config "github.com/COOLizh/task_repo/configs"
	"github.com/COOLizh/task_repo/pkg/models"
	"github.com/dgrijalva/jwt-go"
)

var conf = config.New()

type tokenClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

// CreateToken returns token string with given user's ID
func CreateToken(u *models.User) (string, error) {
	claims := tokenClaims{
		ID: u.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 168).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(conf.JwtSalt))
}

// ParseToken extracts ID from JWT string
func ParseToken(token string) (int, error) {
	tk, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.JwtSalt), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := tk.Claims.(*tokenClaims); ok && tk.Valid {
		return claims.ID, nil
	}
	return 0, err
}
