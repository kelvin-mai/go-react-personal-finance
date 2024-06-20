package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/kelvin-mai/personal-finance/internal/server/router/response"
)

func errorHandler(c *fiber.Ctx, err error) error {
	if err.Error() == jwtware.ErrJWTMissingOrMalformed.Error() {
		return response.ErrorBadRequest(err)
	}
	return response.ErrorUnauthorized(err, "Invalid or expired token")
}

func Authenticate(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(secret)},
		ErrorHandler: errorHandler,
	})
}
