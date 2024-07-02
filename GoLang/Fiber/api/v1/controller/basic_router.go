package basic_router

import (
	"context"

	DB "SMART_OFFICE_APP/pkg/db"

	"github.com/gofiber/fiber/v3"
)

func RootHandler(c fiber.Ctx) error {
	MAIN_DB := "Disconnected"
	LOG_DB := "Disconnected"
	REDIS_DB := "Disconnected"

	if DB.DBMAIN != nil {
		if conn, err := DB.DBMAIN.DB(); err == nil {
			if conn.Ping() == nil {
				MAIN_DB = "Connected"
			}
		}
	}

	if DB.DBLOG.Connect() == nil {
		LOG_DB = "Connected"
	}

	if DB.DBSTORAGE != nil {
		if conn := DB.DBSTORAGE.Conn(); conn.Ping(context.Background()).Val() == "PONG" {
			REDIS_DB = "Connected"
		}
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Welcome to Smart Office API",
		"server_status": fiber.Map{
			"MAIN":   MAIN_DB,
			"LOGGER": LOG_DB,
			"CACHE":  REDIS_DB,
		},
	})
}
