// internal/middleware/logger.go

package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

// LogRequest is a middleware that logs incoming requests using logrus
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.WithFields(logrus.Fields{
			"method": r.Method,
			"url":    r.URL.Path,
			"ip":     r.RemoteAddr,
		}).Info("Incoming request")
		next.ServeHTTP(w, r)
	})
}
