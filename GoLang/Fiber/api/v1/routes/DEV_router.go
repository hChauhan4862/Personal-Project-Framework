package routes

import (
	"github.com/gofiber/fiber/v3"
)

// DevRoutes func for describe group of private routes.
func DevRoutes(a *fiber.App) {
	// Create routes group.
	_ = a.Group("/api/v1")

}
