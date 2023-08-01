package profile

import (
	"bistleague-be/model/config"
	"bistleague-be/model/dto"
	"bistleague-be/services/middleware/guard"
	"bistleague-be/services/usecase/profile"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Router struct {
	cfg *config.Config
	uc  *profile.Usecase
	vld *validator.Validate
}

func New(cfg *config.Config, uc *profile.Usecase, vld *validator.Validate) *Router {
	return &Router{
		cfg: cfg,
		uc:  uc,
		vld: vld,
	}
}

func (r *Router) Register(app *fiber.App) {
	g := app.Group("/profile")
	g.Get("/:uid", guard.DefaultGuard(r.GetUserProfile))
	g.Put("", guard.AuthGuard(r.cfg, r.UpdateUserProfile)...)
}

func (r *Router) GetUserProfile(g *guard.GuardContext) error {
	uid := g.FiberCtx.Params("uid")
	if uid == "" {
		return g.ReturnError(http.StatusBadRequest, "user id is not provided")
	}
	resp, err := r.uc.GetUserProfile(g.FiberCtx.Context(), uid)
	if err != nil {
		return g.ReturnError(http.StatusNotFound, "cannot find user information")
	}
	return g.ReturnSuccess(resp)
}

func (r *Router) UpdateUserProfile(g *guard.AuthGuardContext) error {
	req := dto.UpdateUserProfileRequestDTO{}
	err := r.uc.UpdateUserProfile(g.FiberCtx.Context(), req, g.Claims.UserID)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, "cannot update user profile")
	}
	return g.FiberCtx.JSON(dto.NoBodyDTOResponseWrapper{
		Status:  http.StatusAccepted,
		Message: "user profile has been updated",
	})
}
