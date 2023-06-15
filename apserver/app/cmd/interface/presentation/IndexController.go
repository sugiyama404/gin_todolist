package presentaion

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheck(c *gin.Context) {
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
	})
}
