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

func (r *Router) Register(app *fiber.App) {

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
	return nil
}
