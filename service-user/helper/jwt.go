package helper

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"service-user/domain"
	"strings"
)

var secretKey = "this_is_jwt_secret_shhh"

type JwtClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(userData domain.User) string {
	claims := &JwtClaims{
		Email:            userData.Email,
		RegisteredClaims: jwt.RegisteredClaims{},
	}
	
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	t, err := tokenClaim.SignedString([]byte(secretKey))
	if err != nil {
		panic(err)
	}
	
	return t
}

func TokenMiddleware(next fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("access_token")
		if authHeader == "" {
			return c.Status(http.StatusUnauthorized).SendString("Need Token")
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
			return c.Status(http.StatusUnauthorized).SendString("Invalid Token")
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
