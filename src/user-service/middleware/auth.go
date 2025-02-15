package middleware

import (
	"strconv"
	"time"
	"net/http"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gin-gonic/gin"
)


var jwtSecret = []byte("my_secret_key")

// --- JWT Helper ---
func GenerateJWT(ID int) (string, error) {
	claims := jwt.MapClaims{
		"userID":    strconv.Itoa(ID),
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// --- JWT Middleware Functionality ---
// AuthMiddleware is a Gin middleware that protects routes by verifying JWT tokens.
// 1. It reads the 'Authorization' header from the request.
// 2. If the header is missing, it returns a 401 Unauthorized error.
// 3. It parses and validates the token using the secret key.
// 4. If the token is invalid or expired, it returns a 401 Unauthorized error.
// 5. If the token is valid, the request proceeds to the protected route.

// --- How Routes Are Protected ---
// In the 'routes()' function, we apply `s.Router.Use(middleware.AuthMiddleware())`.
// - This means all routes defined after this line require a valid JWT to access.
// - Routes defined before this middleware are public (e.g., signup, login).
// Example:
// - `/api/v1/auth/signup` (public) - No token needed.
// - `/api/v1/user/profile` (protected) - Requires a valid token.

// --- Security Benefits ---
// - Prevents unauthorized access to sensitive routes.
// - Ensures requests are from authenticated users.
// - Can be extended for role-based access control.

// If you want, I can help you add a route that returns user profile data using this protection.

// --- JWT Middleware Functionality ---
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
			tokenString := c.GetHeader("Authorization")
			if tokenString == "" {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
					c.Abort()
					return
			}

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
					return jwtSecret, nil
			})
			if err != nil || !token.Valid {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
					c.Abort()
					return
			}
			c.Next()
	}
}