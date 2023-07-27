package hello

import (
	"bistleague-be/model/config"
	"bistleague-be/services/middleware/guard"
	"bistleague-be/services/usecase/hello"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
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
	g.Get("", guard.DefaultGuard(r.HelloGet))
	g.Post("", guard.AuthGuard(r.HelloPost))
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

func (r *Router) HelloGet(g *guard.GuardContext) error {
	resp, err := r.helloUC.GetHello(g.FiberCtx)
	if err != nil {
		log.Println(err)
		return g.ReturnError(http.StatusBadRequest, "cannot generate response")
	}
	return g.ReturnSuccess(HelloResponse{
		Msg: resp,
	})
}
