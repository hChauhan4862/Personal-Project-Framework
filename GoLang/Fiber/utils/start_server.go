package utils

import (
	DB "SMART_OFFICE_APP/pkg/db"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v3"
)

// StartServerWithGracefulShutdown function for starting server with a graceful shutdown.
func StartServerWithGracefulShutdown(a *fiber.App) {
	// Create a channel to listen for OS signals.
	c := make(chan os.Signal, 1)

	// Notify the channel for the following signals.
	signal.Notify(c, os.Interrupt)

	// Run the server in a goroutine.
	go func() {
		if err := a.Listen(":3000"); err != nil {
			log.Fatalf("Error: %v", err)
		}
	}()

	// Block until a signal is received.
	<-c

	// Shutdown the server.
	if err := a.Shutdown(); err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Close the database connections.
	DBCONN, err := DB.DBMAIN.DB()
	if err == nil {
		DBCONN.Close()
	}

	// Close the cache Redis connections.
	DBCACHE := DB.DBSTORAGE.Conn()
	if DBCACHE != nil {
		DBCACHE.Close()
	}
}
