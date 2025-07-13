// middlewares/rbac.go
package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// CustomClaims defines the JWT claims with roles.
type CustomClaims struct {
    Roles []string `json:"roles"`
    jwt.RegisteredClaims
}

// Authorize ensures the user has at least one of the allowed roles.
func Authorize(allowedRoles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        claims := c.MustGet("claims").(*CustomClaims)
        for _, role := range claims.Roles {
            for _, allowed := range allowedRoles {
                if role == allowed {
                    c.Next()
                    return
                }
            }
        }
        c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
    }
}
