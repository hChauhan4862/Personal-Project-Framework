package middleware

import (
	"time"

	MODEL_LOG "SMART_OFFICE_APP/pkg/log"

	"github.com/gofiber/fiber/v3"
)

func Logger(c fiber.Ctx) (err error) {
	startTime := time.Now()
	defer func() {
		var userID string
		var USERROLE string

		// Get userID from Locals or set to "Anonymous"
		if id, ok := c.Locals("USERID").(string); ok {
			userID = id
		} else {
			userID = "Anonymous"
		}

		if role, ok := c.Locals("USERROLE").(string); ok {
			USERROLE = role
		} else {
			USERROLE = ""
		}
		log := MODEL_LOG.LOG{
			REQUESTID:      c.Locals("requestid").(string),
			REQUEST_METHOD: c.Method(),
			REQUEST_URI:    c.OriginalURL(),
			REQUEST_BODY:   string(c.Request().Body()),
			USER_AGENT:     c.Get("User-Agent"),
			IP_ADDRESS:     c.IP(),
			DEVICE_ID:      c.Get("X-Device-ID"),
			USERID:         userID,
			USERROLE:       USERROLE,
			STATUS_CODE:    c.Response().StatusCode(),
			EXECUTION_TIME: float64(time.Since(startTime).Milliseconds()),
		}

		if log.STATUS_CODE != fiber.StatusOK {
			log.RESPONSE_BODY = string(c.Response().Body())
		}

		if err != nil {
			log.STATUS_CODE = fiber.StatusInternalServerError
			c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		go log.Insert()
	}()

	return c.Next()
}
