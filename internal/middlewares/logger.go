// internal/middleware/request_logger.go

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// LogRequest is a middleware that logs incoming requests
func LogRequest(c *gin.Context) {
	logrus.WithFields(logrus.Fields{
		"method": c.Request.Method,
		"url":    c.Request.URL.Path,
		"ip":     c.ClientIP(),
	}).Info("Incoming request")

	// Continue to the next handler
	c.Next()
}
