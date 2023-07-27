package utils

import (
	"bistleague-be/model/entity"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWTToken(key string, claims entity.CustomClaim) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key))
}
