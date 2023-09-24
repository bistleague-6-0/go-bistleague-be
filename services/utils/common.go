package utils

import (
	"bistleague-be/model/entity"
	"github.com/golang-jwt/jwt/v5"
	"math/rand"
	"time"
)

func createJWTToken(key string, claims entity.CustomClaim) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key))
}

func GenerateAdminJWTToken(key string, userUID string) (string, error) {
	claims := entity.CustomClaim{
		TeamID: "",
		UserID: userUID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "rest",
			Subject:   "",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2 * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	return createJWTToken(key, claims)
}

func GenerateJWTToken(key string, userUID string, teamUID string) (string, error) {
	claims := entity.CustomClaim{
		TeamID: teamUID,
		UserID: userUID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "rest",
			Subject:   "",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	return createJWTToken(key, claims)
}

func GenerateRefreshKey(key string, userUID string) {

}

func GenerateRandomName() string {
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 128)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
