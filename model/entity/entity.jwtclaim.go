package entity

import "github.com/golang-jwt/jwt/v5"

type CustomClaim struct {
	TeamID string `json:"team_id"`
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}
