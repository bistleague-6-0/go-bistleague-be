package auth

import (
	"bistleague-be/model/config"
	"bistleague-be/services/middleware/guard"
	"bistleague-be/services/utils/httpclient"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

type Router struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Router {
	return &Router{
		cfg: cfg,
	}
}

func (r *Router) RegisterRoute(app *fiber.App) {
	app.Post("/login", guard.DefaultGuard(r.GetTokenByEmail))
	app.Post("/login/token", guard.DefaultGuard(r.ValidateToken))
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenAuth struct {
	Token             string `json:"token"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

func (r *Router) ValidateToken(g *guard.GuardContext) error {
	req := TokenAuth{}
	if err := g.FiberCtx.BodyParser(&req); err != nil {
		return err
	}
	req.ReturnSecureToken = true
	var response struct {
		Uid          string `json:"localId"`
		Email        string `json:"email"`
		DisplayName  string `json:"displayName"`
		IDToken      string `json:"idToken"`
		RefreshToken string `json:"refreshToken"`
		ExpiresIn    string `json:"expiresIn"`
	}
	err := httpclient.Request(
		"",
		httpclient.PostMethod,
		map[string]string{
			"Content-Type": "application/json",
		},
		req,
		&response,
	)
	if err != nil {
		log.Println(err.Error())
		return g.ReturnError(http.StatusBadRequest, "cannot authenticate user")
	}
	return g.ReturnSuccess(response)
}

func (r *Router) GetTokenByEmail(g *guard.GuardContext) error {
	credential := AuthRequest{}
	if err := g.FiberCtx.BodyParser(&credential); err != nil {
		return err
	}
	requestBody := map[string]string{
		"email":             credential.Email,
		"password":          credential.Password,
		"returnSecureToken": "true",
	}
	var response struct {
		Uid          string `json:"localId"`
		Email        string `json:"email"`
		DisplayName  string `json:"displayName"`
		IDToken      string `json:"idToken"`
		RefreshToken string `json:"refreshToken"`
		ExpiresIn    string `json:"expiresIn"`
	}

	err := httpclient.Request(
		r.cfg.Firebase.AuthDomain,
		httpclient.PostMethod,
		map[string]string{
			"Content-Type": "application/json",
		},
		requestBody,
		&response,
	)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, err.Error())
	}
	return g.ReturnSuccess(response)
}
