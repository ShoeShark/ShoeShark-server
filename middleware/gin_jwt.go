package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			c.Abort()
			return
		}

		token, err := ParseJWT(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		address := claims["address"].(string)
		c.Set("accountAddress", address)
		c.Next()
	}
}

func ExtractAccountAddress(c *gin.Context) string {
	return c.GetString("accountAddress")
}

func GenContextWithInformation(c *gin.Context) context.Context {
	return context.WithValue(c.Request.Context(), "accountAddress", ExtractAccountAddress(c))
}
