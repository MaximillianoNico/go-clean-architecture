package webserver

import (
	"log"
	"os"

	router "github.com/MaximillianoNico/go-clean-architecture/app/interface"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type App struct {
	version string
}

func NewApp(
	version string,
) *App {
	return &App{
		version: version,
	}
}

func (svc *App) RunServer() *fiber.App {
	app := fiber.New(fiber.Config{
		ServerHeader: "Fiber",
		AppName:      "Personal Finance Service " + svc.version,
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.Next()
	})
	app.Use(logger.New())

	Router := router.NewAPIRouter(
		app,
	)

	app = Router.Init()

	port := ":3000"
	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	}

	log.Printf("Starting payment server... address %v", port)
	app.Listen(port)

	return app
}
