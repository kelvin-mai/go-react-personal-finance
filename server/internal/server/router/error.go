package router

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
)

func DefaultErrorHandler(ctx *fiber.Ctx, err error) error {
	log.Printf("error: %v\n", err)
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	return ctx.Status(code).JSON(fiber.Map{
		"success": false,
		"error":   e.Message,
		"meta":    generateMetadata(ctx),
	})
}
