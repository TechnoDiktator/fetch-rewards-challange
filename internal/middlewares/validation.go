package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"reflect"
)

// ValidateRequest is a middleware for validating incoming JSON payloads
func ValidateRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Assuming we can decode into the ReceiptRequest struct
		var receipt ReceiptRequest
		err := json.NewDecoder(r.Body).Decode(&receipt)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Warn("Failed to decode request payload")
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Validate using reflection (fields cannot be empty)
		val := reflect.ValueOf(receipt)
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			if field.Kind() == reflect.String && field.Len() == 0 {
				logrus.WithFields(logrus.Fields{
					"field": val.Type().Field(i).Name,
				}).Warn("Empty required field detected")
				http.Error(w, fmt.Sprintf("%s cannot be empty", val.Type().Field(i).Name), http.StatusBadRequest)
				return
			}
		}

		logrus.Info("Request payload is valid")
		next.ServeHTTP(w, r)
	})
}
