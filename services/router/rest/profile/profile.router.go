package profile

import (
	"bistleague-be/model/config"
	"bistleague-be/services/middleware/guard"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Router {
	return &Router{
		cfg: cfg,
	}
}

func (r *Router) Register(app *fiber.App) {
	g := app.Group("/profile")
	g.Get("", guard.AuthGuard(r.cfg, r.GetUserProfile)...)
	g.Put("", guard.AuthGuard(r.cfg, r.UpdateUserProfile)...)
}

func (r *Router) GetUserProfile(g *guard.AuthGuardContext) error {
	uid := g.FiberCtx.Params("uid")
	return g.ReturnSuccess(uid)
}

func (r *Router) UpdateUserProfile(g *guard.AuthGuardContext) error {
	return g.ReturnSuccess(nil)
}
