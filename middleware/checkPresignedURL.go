package middleware

import (
	"net/http"

	"github.com/3AM-Developer/dae/models"
	"github.com/gin-gonic/gin"
)

func CheckPreSignedURL(c *gin.Context) {
	token := c.DefaultQuery("token", "")

	if err := models.VerifyOTU(token); err != nil {
		if err == models.ErrOTUInvalid {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Invalid or expired registration link"})
			return
		}
	}

	c.Next()
}
