package main

import (
	"SMART_OFFICE_APP/utils"

	"SMART_OFFICE_APP/api/middleware"
	basic_router "SMART_OFFICE_APP/api/v1/controller"
	"SMART_OFFICE_APP/pkg/configs"
	DB "SMART_OFFICE_APP/pkg/db"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"github.com/gofiber/fiber/v3/middleware/monitor/skip"
	// "github.com/gofiber/fiber/v2/middleware/idempotency"
)

func main() {

	DB.Init()
	// cache.Init()

	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	middleware.FiberMiddleware(app)
	app.Use(skip.New(
		limiter.New(limiter.Config{
			Max:     60,
			Storage: DB.DBSTORAGE,
		}),
	))

	app.Get("/", basic_router.RootHandler)

	// Custom 404 Not Found
	app.Use(func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusNotFound, "404 Not Found")
	})

	utils.StartServerWithGracefulShutdown(app)
}
