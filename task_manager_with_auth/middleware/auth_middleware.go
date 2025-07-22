package middleware

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte("your_jwt_secret")
func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.IndentedJSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}
	
		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer"{
			c.IndentedJSON(401, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(authParts[1], func(token *jwt.Token)(interface{}, error){
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return jwtSecret, nil
		})
		
		if err != nil || !token.Valid {
			c.IndentedJSON(401, gin.H{"error": "Invalid JWT"})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)


		// Set Values that are needed in the controller 
		c.Set("user_id", claims["user_id"])
		c.Set("username", claims["username"])
		c.Set("role", claims["role"])

		c.Next()

	}

}

func AdminAuthMiddelware() gin.HandlerFunc {
		return func(c *gin.Context){

		roleVal, ok := c.Get("role")
		if !ok {
			c.IndentedJSON(401, gin.H{"error": "Cannot verify user role"})
			c.Abort()
			return
		}

		role, ok := roleVal.(string)
		if !ok || role != "admin" {
			c.IndentedJSON(403, gin.H{"error": "Access denied, only admins perform this operation"})
			c.Abort()
			return
		}

		c.Next()

	}

}