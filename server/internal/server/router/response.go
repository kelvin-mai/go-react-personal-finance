package router

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func generateMetadata(ctx *fiber.Ctx) interface{} {
	return fiber.Map{
		"timestamp": time.Now(),
		"path":      ctx.Path(),
		"method":    ctx.Method(),
	}
}

func Response(ctx *fiber.Ctx, code int, data interface{}) error {
	return ctx.Status(code).JSON(fiber.Map{
		"success": true,
		"data":    data,
		"meta":    generateMetadata(ctx),
	})
}

func Ok(ctx *fiber.Ctx, data interface{}) error {
	return Response(ctx, fiber.StatusOK, data)
}

func Created(ctx *fiber.Ctx, data interface{}) error {
	return Response(ctx, fiber.StatusCreated, data)
}
