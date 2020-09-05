package icors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//CorsMiddleware Middleware to check the cors policies
func CorsMiddleware(acceptedOrigins []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		// fmt.Println("Origin ", c.GetHeader("origin"))
		origin := c.GetHeader("origin")
		if !contains(acceptedOrigins, origin) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   true,
				"message": http.StatusText(http.StatusForbidden),
			})
			return
		}

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
				"error":   true,
				"message": http.StatusText(http.StatusNotAcceptable),
			})
			return
		}

		c.Next()
	}
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
