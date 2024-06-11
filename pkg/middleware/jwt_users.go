package middleware

import (
	"github.com/MinhSang97/order_app/dbutil"
	"github.com/MinhSang97/order_app/pkg/sercurity"
	"github.com/MinhSang97/order_app/pkg/sercurity/claims"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"strings"
)

func JWTMiddlewareUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := dbutil.ConnectDB()
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

		var count int64
		err := db.Table("users").Where("token = ?", tokenString).Count(&count).Error
		if err != nil {
			c.JSON(401, gin.H{"error": "invalid or expired Token"})
			c.Abort()
			return
		}
		if count == 0 {
			c.JSON(401, gin.H{"error": "invalid or expired Token"})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &claims.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(sercurity.SECRET_KEY_USERS), nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "invalid in Bear or expired Token"})
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
