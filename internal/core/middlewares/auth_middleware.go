package middlewares

import (
	"net/http"
	"strings"

	"address-book-server-v2/internal/common/utils"
	"address-book-server-v2/internal/core/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(serverCfg *config.ServerConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		
		jwtSecret := serverCfg.JwtSecret

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.Error(c, http.StatusUnauthorized, "missing token")
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.Error(c, http.StatusUnauthorized, "invalid token format")
			return
		}

		tokenStr := parts[1]

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			utils.Error(c, http.StatusUnauthorized, "invalid token")
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := uint64(claims["user_id"].(float64))
		userEmail := string(claims["user_email"].(string))

		c.Set("user_id", userID)
		c.Set("user_email", userEmail)
		c.Next()
	}
}