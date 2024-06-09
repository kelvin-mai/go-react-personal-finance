package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/kelvin-mai/personal-finance/internal/config"
	"github.com/kelvin-mai/personal-finance/internal/database"
)

type Server struct {
	app  *fiber.App
	port string
	db   *sqlx.DB
}

func NewServer(cfg *config.Config) *Server {
	app := fiber.New()
	port := ":" + cfg.Port
	db := database.Connect(cfg.DatabaseUrl)
	return &Server{
		app:  app,
		port: port,
		db:   db,
	}
}

func (s *Server) Start() error {
	return s.app.Listen(s.port)
}

func (s *Server) Stop() error {
	s.db.Close()
	return s.app.Shutdown()
}
