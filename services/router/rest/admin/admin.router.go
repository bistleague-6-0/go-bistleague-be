package admin

import (
	"bistleague-be/model/config"
	"bistleague-be/model/dto"
	"bistleague-be/services/middleware/guard"
	"bistleague-be/services/usecase/admin"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	cfg     *config.Config
	usecase *admin.Usecase
	vld     *validator.Validate
}

func New(cfg *config.Config, usecase *admin.Usecase, vld *validator.Validate) *Router {
	return &Router{
		cfg:     cfg,
		usecase: usecase,
		vld:     vld,
	}
}

func (r *Router) RegisterRoute(app *fiber.App) {
	g := app.Group("/admin")
	g.Post("/login", guard.DefaultGuard(r.SignInAdmin))
	g.Post("/register", guard.ZeusGuard(r.cfg, r.RegisterAdmin)...)
	g.Get("/payments", guard.AdminGuard(r.cfg, r.GetTeamPayment)...)
	g.Get("/users", guard.AdminGuard(r.cfg, r.GetUserDocsList)...)
	g.Put("/payments/status/:teamID", guard.AdminGuard(r.cfg, r.UpdatePaymentStatus)...)
	g.Put("/users/status/:uid", guard.AdminGuard(r.cfg, r.UpdateUserDocumentStatus)...)
	g.Get("/challenge/mini/:uid", guard.AdminGuard(r.cfg, r.GetMiniChallenge)...)
	g.Get("/challenge/mini", guard.AdminGuard(r.cfg, r.GetAllMiniChallenge)...)
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *Router) SignInAdmin(g *guard.GuardContext) error {
	req := dto.SignInAdminRequestDTO{}
	err := g.FiberCtx.BodyParser(&req)
	if err != nil {
		fmt.Println("error", err)
		return g.ReturnError(http.StatusBadRequest, "cannot find json body")
	}
	err = r.vld.StructCtx(g.FiberCtx.Context(), &req)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, err.Error())
	}
	resp, err := r.usecase.SignInAdmin(g.FiberCtx.Context(), req)
	if err != nil {
		return g.ReturnError(http.StatusNotFound, "wrong username or password")
	}
	return g.ReturnSuccess(resp)
}

func (r *Router) RegisterAdmin(g *guard.GuardContext) error {
	req := dto.RegisterAdminRequestDTO{}
	err := g.FiberCtx.BodyParser(&req)
	if err != nil {
		fmt.Println("error", err)
		return g.ReturnError(http.StatusBadRequest, "cannot find json body")
	}
	err = r.vld.StructCtx(g.FiberCtx.Context(), &req)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, err.Error())
	}
	resp, err := r.usecase.InsertNewAdmin(g.FiberCtx.Context(), req)
	if err != nil {
		fmt.Println(err)
		return g.ReturnError(http.StatusInternalServerError, "cannot register admin")
	}
	return g.ReturnSuccess(resp)
}

func (r *Router) GetTeamPayment(g *guard.AuthGuardContext) error {
	pageStr := g.FiberCtx.Queries()["page"]
	pageSizeStr := g.FiberCtx.Queries()["page_size"]
	if pageStr == "" {
		pageStr = "1"
	}
	if pageSizeStr == "" {
		pageSizeStr = "10"
	}
	page, err := strconv.ParseInt(pageStr, 10, 16)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, "page is not valid int")
	}

	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 16)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, "page size is not valid int")
	}
	resp, err := r.usecase.GetTeamPayment(g.FiberCtx.Context(), int(page), int(pageSize))
	if err != nil {
		fmt.Println(err)
		return g.ReturnError(http.StatusNotFound, "cannot get teams payment")
	}
	return g.ReturnSuccess(resp)
}

func (r *Router) GetUserDocsList(g *guard.AuthGuardContext) error {
	pageStr := g.FiberCtx.Queries()["page"]
	pageSizeStr := g.FiberCtx.Queries()["page_size"]
	if pageStr == "" {
		pageStr = "1"
	}
	if pageSizeStr == "" {
		pageSizeStr = "10"
	}
	page, err := strconv.ParseInt(pageStr, 10, 16)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, "page is not valid int")
	}

	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 16)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, "page size is not valid int")
	}
	resp, err := r.usecase.GetUserList(g.FiberCtx.Context(), int(page), int(pageSize))
	if err != nil {
		fmt.Println(err)
		return g.ReturnError(http.StatusNotFound, "cannot get user docs")
	}
	return g.ReturnSuccess(resp)
}

func (r *Router) UpdatePaymentStatus(g *guard.AuthGuardContext) error {
	req := dto.UpdateTeamPaymentStatus{}
	teamID := g.FiberCtx.Params("teamID")
	err := g.FiberCtx.BodyParser(&req)
	if err != nil {
		fmt.Println("error", err)
		return g.ReturnError(http.StatusBadRequest, "cannot find json body")
	}
	if teamID == "" {
		return g.ReturnError(http.StatusBadRequest, "team id is not provided")
	}
	err = r.vld.StructCtx(g.FiberCtx.Context(), &req)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, err.Error())
	}
	if req.Status != -1 && req.Status != 2 {
		return g.ReturnError(http.StatusBadRequest, "invalid status")
	}
	err = r.usecase.UpdateTeamPaymentStatus(g.FiberCtx.Context(), teamID, req.Status, req.Rejection)
	if err != nil {
		fmt.Println(err)
		return g.ReturnError(http.StatusInternalServerError, "cannot update payment status")
	}
	return g.FiberCtx.JSON(dto.NoBodyDTOResponseWrapper{
		Status:  http.StatusAccepted,
		Message: "team payment status has been updated",
	})
}

func (r *Router) UpdateUserDocumentStatus(g *guard.AuthGuardContext) error {
	req := dto.UpdateUserDocumentStatus{}
	uid := g.FiberCtx.Params("uid")
	err := g.FiberCtx.BodyParser(&req)
	if err != nil {
		fmt.Println("error", err)
		return g.ReturnError(http.StatusBadRequest, "cannot find json body")
	}
	if uid == "" {
		return g.ReturnError(http.StatusBadRequest, "user id is not provided")
	}
	if req.Status != -1 && req.Status != 2 {
		return g.ReturnError(http.StatusBadRequest, "invalid status")
	}
	if req.DocumentType != "student_card" && req.DocumentType != "self_portrait" && req.DocumentType != "twibbon" && req.DocumentType != "enrollment" {
		return g.ReturnError(http.StatusBadRequest, "invalid document type")
	}
	err = r.usecase.UpdateUserDocumentStatus(g.FiberCtx.Context(), uid, req.DocumentType, req.Status, req.Rejection)
	if err != nil {
		fmt.Println(err)
		return g.ReturnError(http.StatusInternalServerError, "cannot update user document status")
	}
	return g.FiberCtx.JSON(dto.NoBodyDTOResponseWrapper{
		Status:  http.StatusAccepted,
		Message: "user document status has been updated",
	})
}

func (r *Router) GetMiniChallenge(g *guard.AuthGuardContext) error {
	uid := g.FiberCtx.Params("uid")
	if uid == "" {
		return g.ReturnError(http.StatusNotFound, "cannot find challenge")
	}
	resp, err := r.usecase.GetMiniChallengeByUIDUsecase(g.FiberCtx.Context(), uid)
	if err != nil {
		return g.ReturnError(http.StatusNotFound, "cannot find challenge")
	}
	return g.ReturnSuccess(resp)
}

func (r *Router) GetAllMiniChallenge(g *guard.AuthGuardContext) error {
	strPage := g.FiberCtx.Query("page", "1")
	page, err := strconv.ParseUint(strPage, 10, 32)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, "page is not an integer")
	}
	strLimit := g.FiberCtx.Query("limit", "10")
	limit, err := strconv.ParseUint(strLimit, 10, 32)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, "limit is not an integer")
	}
	resp, err := r.usecase.GetMiniChallengesUsecase(g.FiberCtx.Context(), page, limit)
	if err != nil {
		log.Error(err)
		return g.ReturnError(http.StatusInternalServerError, "internal server error")
	}
	return g.ReturnSuccess(resp)
}
