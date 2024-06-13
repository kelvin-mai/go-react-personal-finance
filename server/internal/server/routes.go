package server

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/kelvin-mai/personal-finance/internal/server/router"
)

func healthCheck(db *sqlx.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var result int
		err := db.Get(&result, "select 1")
		if err != nil {
			return errors.New("database unavailable")
		}
		return router.Ok(ctx, fiber.Map{
			"database": "available",
		})
	}
}

func (s *Server) SetupRoutes() {
	api := s.app.Group("/api")
	api.Get("/", healthCheck(s.db))
}
