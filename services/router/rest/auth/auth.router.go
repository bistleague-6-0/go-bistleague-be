package auth

import (
	"bistleague-be/model/config"
	"bistleague-be/model/dto"
	"bistleague-be/model/entity"
	"bistleague-be/services/middleware/guard"
	"bistleague-be/services/usecase/auth"
	"bistleague-be/services/utils"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

type Router struct {
	cfg     *config.Config
	usecase *auth.Usecase
	vld     *validator.Validate
}

func New(cfg *config.Config, usecase *auth.Usecase) *Router {
	vld := validator.New() //MARK: move it to common resource
	return &Router{
		cfg:     cfg,
		usecase: usecase,
		vld:     vld,
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
	req := dto.CreateUserRequestDTO{}
	err := g.FiberCtx.BodyParser(&req)
	if err != nil {
		fmt.Println("error", err)
		return g.ReturnError(http.StatusInternalServerError, "cannot decode json body")
	}
	err = r.vld.StructCtx(g.FiberCtx.Context(), &req)
	if err != nil {
		return g.ReturnError(http.StatusInternalServerError, err.Error())
	}
	if req.RePassword != req.Password {
		return g.ReturnError(http.StatusBadRequest, "password does not match")
	}
	resp, err := r.usecase.InsertNewUser(g.FiberCtx.Context(), req)
	if err != nil {
		return g.ReturnError(http.StatusInternalServerError, err.Error())
	}
	return g.ReturnSuccess(resp)
}
