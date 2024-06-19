package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func errorHandler(c *fiber.Ctx, err error) error {
	if err.Error() == jwtware.ErrJWTMissingOrMalformed.Error() {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired JWT")
}

func Authenticate(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(secret)},
		ErrorHandler: errorHandler,
	})
}
