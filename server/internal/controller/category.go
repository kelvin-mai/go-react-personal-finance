package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kelvin-mai/personal-finance/internal/model"
	"github.com/kelvin-mai/personal-finance/internal/server/router/response"
	"github.com/kelvin-mai/personal-finance/internal/service"
)

type CategoryController struct {
	s *service.CategoryService
}

func NewCategoryController(s *service.CategoryService) *CategoryController {
	return &CategoryController{
		s: s,
	}
}

func (c *CategoryController) List(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["sub"].(string)
	categories, err := c.s.GetAllByUserId(id)
	if err != nil {
		return response.ErrorNotFound(err)
	}
	return response.Ok(ctx, fiber.Map{
		"categories": categories,
	})
}

func (c *CategoryController) Get(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)
	id := ctx.Params("id")
	category, err := c.s.GetById(id)
	if err != nil {
		return response.ErrorNotFound(err)
	}
	if category.UserId != userId {
		return response.ErrorUnauthorized(nil, "unauthorized")
	}
	return response.Ok(ctx, fiber.Map{
		"category": category,
	})
}

func (c *CategoryController) Create(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)
	input := model.CategoryInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}
	category, err := c.s.Create(userId, input.Title)
	if err != nil {
		return response.ErrorBadRequest(err)
	}
	return response.Created(ctx, fiber.Map{
		"category": category,
	})
}

func (c *CategoryController) Update(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)
	id := ctx.Params("id")
	input := model.CategoryInput{}
	if err := ctx.BodyParser(&input); err != nil {
		return response.ErrorBadRequest(err)
	}
	category, err := c.s.Update(id, userId, input.Title)
	if err != nil {
		return response.ErrorBadRequest(err)
	}
	return response.Ok(ctx, fiber.Map{
		"category": category,
	})
}

func (c *CategoryController) Delete(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)
	id := ctx.Params("id")
	category, err := c.s.Delete(id, userId)
	if err != nil {
		return response.ErrorBadRequest(err)
	}
	return response.Ok(ctx, fiber.Map{
		"category": category,
	})
}
