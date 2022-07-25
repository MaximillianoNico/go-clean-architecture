package Interface

import (
	middleware "github.com/MaximillianoNico/go-clean-architecture/app/infrastructure/middleware"
	controllers "github.com/MaximillianoNico/go-clean-architecture/app/interface/controllers"
	"github.com/gofiber/fiber/v2"
)

type APIRouter struct {
	app *fiber.App
}

func NewAPIRouter(
	app *fiber.App,
) *APIRouter {
	return &APIRouter{
		app: app,
	}
}

func (r *APIRouter) Init() *fiber.App {
	ctr := controllers.NewController()
	m := middleware.InitMiddleware()

	// Grouping Routing
	api := r.app.Group("/api", m.PublicRoute)

	api.Get("/health-check", ctr.HealthCheck)

	return r.app
}
