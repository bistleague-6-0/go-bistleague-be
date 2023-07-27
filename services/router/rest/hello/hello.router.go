package hello

import (
	"bistleague-be/model/config"
	"bistleague-be/services/middleware/guard"
	"bistleague-be/services/usecase/hello"
	"encoding/hex"
	"github.com/gofiber/fiber/v2"
	"math/rand"
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
	g.Get("", guard.DefaultGuard(r.HelloGet))
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

// token randomizer
func (r *Router) HelloGet(g *guard.GuardContext) error {
	b := make([]byte, 4) //equals 8 characters
	rand.Read(b)
	s := hex.EncodeToString(b)
	return g.ReturnSuccess(s)
}
