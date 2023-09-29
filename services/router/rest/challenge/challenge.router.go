package challenge

import (
	"bistleague-be/model/config"
	"bistleague-be/model/dto"
	"bistleague-be/services/middleware/guard"
	"bistleague-be/services/usecase/challenge"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Router struct {
	cfg *config.Config
	vld *validator.Validate
	uc  *challenge.Usecase
}

func New(cfg *config.Config, vld *validator.Validate, uc *challenge.Usecase) *Router {
	return &Router{
		cfg: cfg,
		vld: vld,
		uc:  uc,
	}
}

func (r *Router) Register(app *fiber.App) {
	g := app.Group("/challenge")
	g.Get("", guard.AuthGuard(r.cfg, r.GetChallengeRouter)...)
	g.Post("", guard.AuthGuard(r.cfg, r.AddNewChallengeRouter)...)
	g.Put("", guard.AuthGuard(r.cfg, r.UpdateChallengeRouter)...)
}

func (r *Router) AddNewChallengeRouter(g *guard.AuthGuardContext) error {
	req := dto.InsertChallengeRequestDTO{}
	err := g.FiberCtx.BodyParser(&req)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, "cannot find json body")
	}
	err = r.vld.StructCtx(g.FiberCtx.Context(), &req)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, err.Error())
	}
	if req.IgUsername == "" || req.IgContentURl == "" {
		return g.ReturnError(http.StatusBadRequest, "ig cannot be empty")
	}
	resp, err := r.uc.AddNewChallengeUsecase(g.FiberCtx.Context(), req, g.Claims.UserID)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, err.Error())
	}
	return g.ReturnSuccess(resp)
}

func (r *Router) UpdateChallengeRouter(g *guard.AuthGuardContext) error {
	req := dto.UpdateChallengeRequestDTO{}
	err := g.FiberCtx.BodyParser(&req)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, "cannot find json body")
	}
	err = r.vld.StructCtx(g.FiberCtx.Context(), &req)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, err.Error())
	}
	if req.UID != g.Claims.UserID {
		return g.ReturnError(http.StatusBadRequest, "user id does not match")
	}
	resp, err := r.uc.UpdateChallengeUsecase(g.FiberCtx.Context(), req, g.Claims.UserID)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, err.Error())
	}
	return g.ReturnSuccess(resp)
}

func (r *Router) GetChallengeRouter(g *guard.AuthGuardContext) error {
	userID := g.Claims.UserID
	resp, err := r.uc.GetChallengeUsecase(g.FiberCtx.Context(), userID)
	if err != nil {
		return g.ReturnError(http.StatusNotFound, "cannot find user's data")
	}
	return g.ReturnSuccess(resp)
}
