package guard

import (
	"bistleague-be/model/config"
	"bistleague-be/model/dto"
	"bistleague-be/model/entity"
	"net/http"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type GuardContext struct {
	FiberCtx *fiber.Ctx
}

type AuthGuardContext struct {
	FiberCtx *fiber.Ctx
	Claims   entity.CustomClaim
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

func DefaultGuard(handlerFunc func(g *GuardContext) error) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		guardCtx := GuardContext{
			FiberCtx: ctx,
		}
		return handlerFunc(&guardCtx)
	}
}

func AuthGuard(cfg *config.Config, handlerFunc func(g *AuthGuardContext) error) []fiber.Handler {
	handlers := []fiber.Handler{
		jwtware.New(jwtware.Config{SigningKey: jwtware.SigningKey{
			Key: []byte(cfg.Secret.JWTSecret),
		}}),
		func(ctx *fiber.Ctx) error {
			user := ctx.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			expireAt, err := claims.GetExpirationTime()
			if err != nil {
				return ctx.Status(http.StatusUnauthorized).JSON(dto.NoBodyDTOResponseWrapper{
					Status:  http.StatusUnauthorized,
					Message: "unauthorized",
				})
			}
			issuedAt, err := claims.GetIssuedAt()
			if err != nil {
				return ctx.Status(http.StatusUnauthorized).JSON(dto.NoBodyDTOResponseWrapper{
					Status:  http.StatusUnauthorized,
					Message: "unauthorized",
				})
			}
			teamID, ok := claims["team_id"].(string)
			if !ok {
				return ctx.Status(http.StatusUnauthorized).JSON(dto.NoBodyDTOResponseWrapper{
					Status:  http.StatusUnauthorized,
					Message: "unauthorized",
				})
			}
			userID, ok := claims["user_id"].(string)
			if !ok {
				return ctx.Status(http.StatusUnauthorized).JSON(dto.NoBodyDTOResponseWrapper{
					Status:  http.StatusUnauthorized,
					Message: "unauthorized",
				})
			}
			ety := entity.CustomClaim{
				TeamID: teamID,
				UserID: userID,
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: expireAt,
					IssuedAt:  issuedAt,
				},
			}
			authGuardCtx := AuthGuardContext{
				FiberCtx: ctx,
				Claims:   ety,
			}
			return handlerFunc(&authGuardCtx)
		},
	}
	return handlers
}

func AdminGuard(cfg *config.Config, handlerFunc func(g *AuthGuardContext) error) []fiber.Handler {
	handlers := []fiber.Handler{
		jwtware.New(jwtware.Config{SigningKey: jwtware.SigningKey{
			Key: []byte(cfg.Secret.AdminJWT),
		}}),
		func(ctx *fiber.Ctx) error {
			user := ctx.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			expireAt, err := claims.GetExpirationTime()
			if err != nil {
				return ctx.Status(http.StatusUnauthorized).JSON(dto.NoBodyDTOResponseWrapper{
					Status:  http.StatusUnauthorized,
					Message: "unauthorized",
				})
			}
			issuedAt, err := claims.GetIssuedAt()
			if err != nil {
				return ctx.Status(http.StatusUnauthorized).JSON(dto.NoBodyDTOResponseWrapper{
					Status:  http.StatusUnauthorized,
					Message: "unauthorized",
				})
			}
			userID, ok := claims["user_id"].(string)
			if !ok {
				return ctx.Status(http.StatusUnauthorized).JSON(dto.NoBodyDTOResponseWrapper{
					Status:  http.StatusUnauthorized,
					Message: "unauthorized",
				})
			}
			ety := entity.CustomClaim{
				TeamID: "",
				UserID: userID,
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: expireAt,
					IssuedAt:  issuedAt,
				},
			}
			adminGuardCtx := AuthGuardContext{
				FiberCtx: ctx,
				Claims:   ety,
			}
			return handlerFunc(&adminGuardCtx)
		},
	}
	return handlers
}

func ZeusGuard(cfg *config.Config, handlerFunc func(g *GuardContext) error) []fiber.Handler {
	handlers := []fiber.Handler{
		func(ctx *fiber.Ctx) error {
			headers := ctx.GetReqHeaders()
			tokenValue := headers["BIST-Zeus-Token"]
			if tokenValue != cfg.Secret.AdminSecret {
				return ctx.Status(http.StatusUnauthorized).JSON(dto.NoBodyDTOResponseWrapper{
					Status:  http.StatusUnauthorized,
					Message: "unauthorized",
				})
			}

			zeusGuardCtx := GuardContext{
				FiberCtx: ctx,
			}
			return handlerFunc(&zeusGuardCtx)
		},
	}
	return handlers
}
