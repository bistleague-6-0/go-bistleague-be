package hello

import (
	"bistleague-be/model/config"
	"bistleague-be/services/middleware/guard"
	"bistleague-be/services/usecase/hello"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	cfg     *config.Config
	helloUC *hello.Usecase
}

func New(cfg *config.Config, helloUC *hello.Usecase) *Router {
	return &Router{
		cfg:     cfg,
		helloUC: helloUC,
	}
}

func (r *Router) Register(app *fiber.App) {
	g := app.Group("/hello")
	//g.Use()
	g.Get("", guard.AuthGuard(r.cfg, r.HelloGet)...)
	g.Post("", guard.AuthGuard(r.cfg, r.HelloPost)...)
}

type HelloRequest struct {
	Body string `json:"body"`
}

type HelloResponse struct {
	Msg string `json:"msg"`
}

func (r *Router) HelloPost(g *guard.AuthGuardContext) error {
	req := HelloRequest{}
	g.FiberCtx.BodyParser(&req)
	return g.ReturnSuccess(req)
}

func (r *Router) HelloGet(g *guard.AuthGuardContext) error {
	return g.ReturnSuccess(g.Claims)
}
