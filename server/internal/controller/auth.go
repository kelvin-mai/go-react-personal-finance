package controller

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kelvin-mai/personal-finance/internal/model"
	"github.com/kelvin-mai/personal-finance/internal/server/router"
	"github.com/kelvin-mai/personal-finance/internal/service"
	"github.com/kelvin-mai/personal-finance/pkg/util"
)

type AuthController struct {
	s      *service.UserService
	secret string
}

func NewAuthController(s *service.UserService, secret string) *AuthController {
	return &AuthController{
		s:      s,
		secret: secret,
	}
}

func (c *AuthController) createToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		"iat": time.Now(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(c.secret))
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	input := model.AuthInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}
	password, err := util.HashPassword(input.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid password")
	}
	user, err := c.s.Create(input.Username, password)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "registration error")
	}
	token, err := c.createToken(user.Username)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "registration error")
	}
	return router.Created(ctx, fiber.Map{
		"user":  user,
		"token": token,
	})
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	input := model.AuthInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}
	user, err := c.s.GetByUsername(input.Username)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "login error")
	}
	if !util.CheckPassword(input.Password, user.Password) {
		return fiber.NewError(fiber.StatusUnauthorized, "login error")
	}
	token, err := c.createToken(user.Username)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "registration error")
	}
	return router.Ok(ctx, fiber.Map{
		"user":  user,
		"token": token,
	})
}

func (c *AuthController) Me(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["sub"].(string)
	currentUser, err := c.s.GetByUsername(username)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid token")
	}
	token, err := c.createToken(currentUser.Username)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid token")
	}
	return router.Ok(ctx, fiber.Map{
		"user":  currentUser,
		"token": token,
	})
}
