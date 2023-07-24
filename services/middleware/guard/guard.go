package guard

import (
	"bistleague-be/model/dto"
	"firebase.google.com/go/auth"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

type GuardContext struct {
	FiberCtx *fiber.Ctx
}

type AuthGuardContext struct {
	FiberCtx  *fiber.Ctx
	AuthToken *auth.Token
}

func (g *GuardContext) ReturnError(
	status int,
	message string,
) error {
	return g.FiberCtx.Status(status).JSON(dto.NoBodyDTOResponseWrapper{
		Status:  status,
		Message: message,
	})
}

func (g *GuardContext) ReturnSuccess(
	body interface{},
) error {
	return g.FiberCtx.Status(http.StatusOK).JSON(dto.DefaultDTOResponseWrapper{
		Status:  http.StatusOK,
		Message: "ok",
		Body:    body,
	})
}

func (g *AuthGuardContext) ReturnError(
	status int,
	message string,
) error {
	return g.FiberCtx.Status(status).JSON(dto.NoBodyDTOResponseWrapper{
		Status:  status,
		Message: message,
	})
}

func (g *AuthGuardContext) ReturnSuccess(
	body interface{},
) error {
	return g.FiberCtx.Status(http.StatusOK).JSON(dto.DefaultDTOResponseWrapper{
		Status:  http.StatusOK,
		Message: "ok",
		Body:    body,
	})
}

type Guardian struct {
	AuthClient *auth.Client
}

func DefaultGuard(handlerFunc func(g *GuardContext) error) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		guardCtx := GuardContext{
			FiberCtx: ctx,
		}
		return handlerFunc(&guardCtx)
	}
}

func AuthGuard(client *auth.Client, handlerFunc func(g *AuthGuardContext) error) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authorization := ctx.Get("Authorization")
		if !strings.Contains(authorization, "Bearer") {
			return ctx.Status(http.StatusUnauthorized).JSON(dto.NoBodyDTOResponseWrapper{
				Status:  http.StatusUnauthorized,
				Message: "unauthorized",
			})
		}
		bearer := authorization[7:]
		token, err := client.VerifyIDToken(ctx.Context(), bearer)
		if err != nil {
			return ctx.Status(http.StatusUnauthorized).JSON(dto.NoBodyDTOResponseWrapper{
				Status:  http.StatusUnauthorized,
				Message: "unauthorized",
			})
		}
		authGuardCtx := AuthGuardContext{
			FiberCtx:  ctx,
			AuthToken: token,
		}
		return handlerFunc(&authGuardCtx)
	}
}
