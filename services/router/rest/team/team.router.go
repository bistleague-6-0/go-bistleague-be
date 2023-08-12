package team

import (
	"bistleague-be/model/config"
	"bistleague-be/model/dto"
	"bistleague-be/services/middleware/guard"
	"bistleague-be/services/usecase/team"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
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
	group.Get("", guard.AuthGuard(r.cfg, r.GetTeamInformation)...)
	group.Post("/redeem", guard.AuthGuard(r.cfg, r.RedeemTeamCode)...)
	group.Post("/document", guard.AuthGuard(r.cfg, r.InsertTeamDocument)...)
}

// MARK : NEED TO UPDATE
func (r *Router) CreateTeam(g *guard.AuthGuardContext) error {
	if g.Claims.TeamID != "" {
		return g.ReturnError(http.StatusNotAcceptable, "user already registered to a team")
	}
	req := dto.CreateTeamRequestDTO{}
	err := g.FiberCtx.BodyParser(&req)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, "cannot find json body")
	}
	err = r.vld.StructCtx(g.FiberCtx.Context(), &req)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, err.Error())
	}
	resp, err := r.usecase.CreateTeam(g.FiberCtx.Context(), req, g.Claims.UserID)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, "cannot create team")
	}
	return g.FiberCtx.JSON(dto.DefaultDTOResponseWrapper{
		Status:  http.StatusAccepted,
		Message: "team has been created",
		Body:    resp,
	})
}

func (r *Router) GetTeamInformation(g *guard.AuthGuardContext) error {
	if g.Claims.TeamID == "" {
		return g.ReturnError(http.StatusNotFound, "user is not registered at any team")
	}
	resp, err := r.usecase.GetTeamInformation(g.FiberCtx.Context(), g.Claims.TeamID, g.Claims.UserID)
	if err != nil {
		log.Println(err)
		return g.ReturnError(http.StatusNotFound, "cannot find team information")
	}
	return g.ReturnSuccess(resp)
}

func (r *Router) RedeemTeamCode(g *guard.AuthGuardContext) error {
	if g.Claims.TeamID != "" {
		return g.ReturnError(http.StatusNotAcceptable, "user already registered to a team")
	}
	req := dto.RedeemTeamCodeRequestDTO{}
	err := g.FiberCtx.BodyParser(&req)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, "cannot find json body")
	}
	err = r.vld.StructCtx(g.FiberCtx.Context(), &req)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, "redeem code is invalid")
	}
	jwtToken, err := r.usecase.RedeemTeamCode(g.FiberCtx.Context(), req, g.Claims.UserID)
	if err != nil {
		return g.ReturnError(http.StatusNotAcceptable, "redeem code is invalid/expired")
	}
	//MARK : Create a better response message
	return g.FiberCtx.JSON(dto.DefaultDTOResponseWrapper{
		Status:  http.StatusAccepted,
		Message: "successfully joined a team",
		Body: map[string]string{
			"jwt_token": jwtToken,
		},
	})
}

func (r *Router) InsertTeamDocument(g *guard.AuthGuardContext) error {
	req := dto.InsertTeamDocumentRequestDTO{}
	err := g.FiberCtx.BodyParser(&req)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, "cannot find json body")
	}
	err = r.vld.StructCtx(g.FiberCtx.Context(), &req)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, err.Error())
	}
	filename, err := r.usecase.InsertTeamDocument(g.FiberCtx.Context(), req, g.Claims.TeamID, g.Claims.UserID)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, err.Error())
	}
	return g.ReturnSuccess(map[string]string{
		"doc_type": req.Type,
		"filename": filename,
	})
}
