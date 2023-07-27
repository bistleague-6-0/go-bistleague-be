package auth

import (
	"bistleague-be/model/config"
	"bistleague-be/model/entity"
	"bistleague-be/services/middleware/guard"
	"bistleague-be/services/utils/httpclient"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
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
	app.Get("/token", guard.DefaultGuard(r.CreateSign))
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	strToken, err := token.SignedString([]byte(r.cfg.Secret.JWTSecret))
	if err != nil {
		return g.ReturnError(500, err.Error())
	}
	return g.ReturnSuccess(map[string]interface{}{
		"token": strToken,
	})
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
