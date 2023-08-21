package auth

import (
	"bistleague-be/model/config"
	"bistleague-be/model/dto"
	"bistleague-be/services/middleware/guard"
	"bistleague-be/services/usecase/auth"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Router struct {
	cfg     *config.Config
	usecase *auth.Usecase
	vld     *validator.Validate
}

func New(cfg *config.Config, usecase *auth.Usecase, vld *validator.Validate) *Router {
	return &Router{
		cfg:     cfg,
		usecase: usecase,
		vld:     vld,
	}
}

func (r *Router) RegisterRoute(app *fiber.App) {
	app.Post("/login", guard.DefaultGuard(r.SignInUser))
	app.Post("/register", guard.DefaultGuard(r.SignUpUser))
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *Router) SignInUser(g *guard.GuardContext) error {
	req := dto.SignInUserRequestDTO{}
	err := g.FiberCtx.BodyParser(&req)
	if err != nil {
		fmt.Println("error", err)
		return g.ReturnError(http.StatusBadRequest, "cannot find json body")
	}
	err = r.vld.StructCtx(g.FiberCtx.Context(), &req)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, err.Error())
	}
	resp, err := r.usecase.SignInUser(g.FiberCtx.Context(), req)
	if err != nil {
		return g.ReturnError(http.StatusNotFound, "wrong username or password")
	}
	return g.ReturnSuccess(resp)
}

func (r *Router) SignUpUser(g *guard.GuardContext) error {
	req := dto.SignUpUserRequestDTO{}
	err := g.FiberCtx.BodyParser(&req)
	if err != nil {
		fmt.Println("error", err)
		return g.ReturnError(http.StatusBadRequest, "cannot find json body")
	}
	err = r.vld.StructCtx(g.FiberCtx.Context(), &req)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, err.Error())
	}
	if req.RePassword != req.Password {
		return g.ReturnError(http.StatusBadRequest, "password does not match")
	}
	resp, err := r.usecase.InsertNewUser(g.FiberCtx.Context(), req)
	if err != nil {
		fmt.Println(err)
		return g.ReturnError(http.StatusInternalServerError, "cannot register user")
	}
	return g.ReturnSuccess(resp)
}
