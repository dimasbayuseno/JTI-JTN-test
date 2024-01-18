package jwt

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"time"
)

func ValidateTokenMiddleware(c *fiber.Ctx) error {
	// Get the Bearer token from the Authorization header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Missing Authorization header"})
	}

	// Parse the token
	tokenString := extractBearerToken(authHeader)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("JWT_SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid or expired token"})
	}

	// Set the user ID from the token in the context for further processing
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal server error"})
	}

	// Extract user ID from claims and store it in the context
	userID, ok := claims["sub"].(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal server error"})
	}
	c.Locals("userID", userID)

	// Continue to the next handler
	return c.Next()
}

func extractBearerToken(authHeader string) string {
	// Check if the Authorization header starts with "Bearer"
	const prefix = "Bearer "
	if len(authHeader) > len(prefix) && authHeader[:len(prefix)] == prefix {
		return authHeader[len(prefix):]
	}
	return ""
}

// generateToken generates a JWT token for the user
func GenerateToken(userID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiration time (24 hours)

	// TODO: Set any additional claims needed

	secret := []byte(viper.GetString("JWT_SECRET_KEY")) // Replace with your JWT secret
	return token.SignedString(secret)
}
