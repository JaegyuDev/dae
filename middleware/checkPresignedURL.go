package middleware

import (
	"github.com/3AM-Developer/dae/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckPreSignedURL(c *gin.Context) {
	token := c.DefaultQuery("token", "")

	if err := models.VerifyOTU(token); err != nil {
		if err == models.ErrOTUInvalid {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Invalid or expired registration link"})
			return
		}
		// Handle other possible errors, if needed
	}

	c.Next()
}
