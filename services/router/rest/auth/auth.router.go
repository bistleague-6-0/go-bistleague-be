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
	app.Post("/refresh", guard.DefaultGuard(r.RefreshToken))
	app.Post("/forget/password", guard.DefaultGuard(r.ForgetPasswordRoute))
	app.Get("/forget/password", guard.DefaultGuard(r.ValidateForgetPasswordRoute))
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

func (r *Router) RefreshToken(g *guard.GuardContext) error {
	req := dto.RefreshTokenRequestDTO{}
	err := g.FiberCtx.BodyParser(&req)
	if err != nil {
		fmt.Println("error", err)
		return g.ReturnError(http.StatusBadRequest, "cannot find json body")
	}
	err = r.vld.StructCtx(g.FiberCtx.Context(), &req)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, err.Error())
	}
	resp, err := r.usecase.RefreshToken(g.FiberCtx.Context(), req)
	if err != nil {
		return g.ReturnError(http.StatusNotFound, "wrong refresh key")
	}
	return g.ReturnSuccess(resp)
}

func (r *Router) ForgetPasswordRoute(g *guard.GuardContext) error {
	var req struct {
		Email string `json:"email" validate:"required,email"`
	}
	err := g.FiberCtx.BodyParser(&req)
	if err != nil {
		fmt.Println("error", err)
		return g.ReturnError(http.StatusBadRequest, "cannot find json body")
	}
	err = r.vld.StructCtx(g.FiberCtx.Context(), &req)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, err.Error())
	}
	err = r.usecase.ForgetPasswordUsecase(g.FiberCtx.Context(), req.Email)
	if err != nil {
		return g.ReturnError(http.StatusInternalServerError, "internal server error")
	}
	return g.ReturnSuccess("success")
}

func (r *Router) ValidateForgetPasswordRoute(g *guard.GuardContext) error {
	token := g.FiberCtx.Query("token")
	if token == "" {
		return g.ReturnError(http.StatusBadRequest, "token is invalid")
	}
	err := r.usecase.ValidateForgetPasswordTokenUsecase(g.FiberCtx.Context(), token)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, "token is invalid")
	}
	return g.ReturnSuccess(token)
}
