package server

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/kelvin-mai/personal-finance/internal/controller"
	"github.com/kelvin-mai/personal-finance/internal/server/router"
	"github.com/kelvin-mai/personal-finance/internal/server/router/middleware"
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

func (s *Server) SetupRoutes(uc *controller.AuthController) {
	api := s.app.Group("/api")
	api.Get("/", healthCheck(s.db))

	api.Post("/login", uc.Login)
	api.Post("/register", uc.Register)
	api.Get("/me", middleware.Authenticate(s.jwtSecret), uc.Me)
}
