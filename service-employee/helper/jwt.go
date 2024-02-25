package helper

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"service-employee/dto"
	"strings"
)

var secretKey = "this_is_jwt_secret_shhh"

func TokenMiddleware(next fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("access_token")
		if authHeader == "" {
			return c.Status(http.StatusUnauthorized).JSON(dto.WebResponse{
				Code:   401,
				Status: "Unauthorized",
				Data:   "Need Token",
			})
		}
		
		var tokenString string
		
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenString = authHeader[7:] // Mengambil bagian setelah "Bearer "
		} else {
			tokenString = authHeader // Menggunakan token langsung jika tidak menggunakan skema "Bearer"
		}
		
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
		
		if err != nil || !token.Valid {
			return c.Status(http.StatusUnauthorized).JSON(dto.WebResponse{
				Code:   401,
				Status: "Unauthorized",
				Data:   "Invalid Token",
			})
		}
		//
		var userPayload = map[string]string{
			"email": token.Claims.(jwt.MapClaims)["email"].(string),
		}
		
		if token.Valid {
			c.Locals("email", userPayload)
		}
		
		return next(c)
	}
}
