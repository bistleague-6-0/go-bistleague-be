package team

import (
	"bistleague-be/model/config"
	"bistleague-be/model/dto"
	"bistleague-be/services/middleware/guard"
	"bistleague-be/services/usecase/team"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Router struct {
	cfg     *config.Config
	vld     *validator.Validate
	usecase *team.Usecase
}

func New(cfg *config.Config, vld *validator.Validate, usecase *team.Usecase) *Router {
	return &Router{
		cfg:     cfg,
		vld:     vld,
		usecase: usecase,
	}
}

func (r *Router) Register(app *fiber.App) {
	group := app.Group("/team")
	group.Post("", guard.AuthGuard(r.cfg, r.CreateTeam)...)
}

//MARK : NEED TO UPDATE
func (r *Router) CreateTeam(g *guard.AuthGuardContext) error {
	req := dto.CreateTeamRequestDTO{}
	err := g.FiberCtx.BodyParser(&req)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, "cannot find json body")
	}
	err = r.vld.StructCtx(g.FiberCtx.Context(), &req)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, err.Error())
	}
	err = r.usecase.CreateTeam(g.FiberCtx.Context(), req, g.Claims.UserID)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, "cannot create team")
	}
	return g.FiberCtx.JSON(dto.NoBodyDTOResponseWrapper{
		Status:  http.StatusAccepted,
		Message: "team has been created",
	})
}
