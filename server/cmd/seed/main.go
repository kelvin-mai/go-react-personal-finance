package main

import (
	"log"

	"github.com/kelvin-mai/personal-finance/internal/config"
	"github.com/kelvin-mai/personal-finance/internal/database"
	"github.com/kelvin-mai/personal-finance/internal/model"
	"github.com/kelvin-mai/personal-finance/pkg/util"
)

func main() {
	cfg := config.Load()
	db := database.Connect(cfg.DatabaseUrl)
	password, err := util.HashPassword("password")
	if err != nil {
		log.Fatalf("Error generating password: %v\n", err)
	}
	users := []model.User{
		{
			Username: "admin",
			Password: password,
		},
	}
	_, err = db.NamedExec(
		`insert into users (username, password)
		 values (:username, :password)`,
		users,
	)
	if err != nil {
		log.Fatalf("Error inserting users: %v\n", err)
	}
	log.Printf("Successfully inserted users: %v\n", users)
}
