package middleware

import (
	"github.com/MinhSang97/order_app/pkg/sercurity"
	"github.com/MinhSang97/order_app/pkg/sercurity/claims"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func JWTMiddlewareAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.JSON(401, gin.H{"error": "invalid Token in Bear"})
			c.Abort()
			return
		}

		tokenString := authParts[1]

		token, _ := jwt.ParseWithClaims(tokenString, &claims.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(sercurity.SECRET_KEY_ADMIN), nil
		})

		if token.Valid {
			claims, ok := token.Claims.(*claims.JwtCustomClaims)
			if !ok || time.Now().Unix() > claims.ExpiresAt {
				c.JSON(401, gin.H{"error": "invalid or expired Token"})
				c.Abort()
				return
			}
		} else {
			c.JSON(401, gin.H{"error": "invalid or expired Token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*claims.JwtCustomClaims)
		if !ok {
			c.JSON(401, gin.H{"error": "invalid in Bear or expired Token"})
			c.Abort()
			return
		}

		c.Set("user", claims.UserId)
		c.Next()
	}
}
