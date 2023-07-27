package auth

import (
	"bistleague-be/model/config"
	"bistleague-be/model/entity"
	"bistleague-be/services/middleware/guard"
	"bistleague-be/services/usecase/auth"
	"bistleague-be/services/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

type Router struct {
	cfg     *config.Config
	usecase *auth.Usecase
}

func New(cfg *config.Config, usecase *auth.Usecase) *Router {
	return &Router{
		cfg:     cfg,
		usecase: usecase,
	}
}

func (r *Router) RegisterRoute(app *fiber.App) {
	app.Post("/login", guard.DefaultGuard(r.SignInUser))
	app.Get("/token", guard.DefaultGuard(r.CreateSign))
	app.Post("/register", guard.DefaultGuard(r.SignUpUser))
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenAuth struct {
	Token             string `json:"token"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

func (r *Router) CreateSign(g *guard.GuardContext) error {
	claims := entity.CustomClaim{
		TeamID: "s08d8d",
		UserID: "skss",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "rest",
			Subject:   "",
			ExpiresAt: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token, err := utils.CreateJWTToken(r.cfg.Secret.JWTSecret, claims)
	if err != nil {
		return g.ReturnError(500, err.Error())
	}
	return g.ReturnSuccess(map[string]interface{}{
		"token": token,
	})
}

func (r *Router) SignInUser(g *guard.GuardContext) error {
	return nil
}

func (r *Router) SignUpUser(g *guard.GuardContext) error {
	sql, err := r.usecase.InsertNewUser(g.FiberCtx.Context())
	if err != nil {
		return g.ReturnError(http.StatusInternalServerError, err.Error())
	}
	return g.ReturnSuccess(sql)
}
