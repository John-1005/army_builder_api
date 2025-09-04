package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type contextKey string

const requestIDKey contextKey = "requestID"

func MiddlewareRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		reqID := uuid.New().String()

		ctx := context.WithValue(r.Context(), requestIDKey, reqID)

		w.Header().Set("X-Request-ID", reqID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetRequestID(r *http.Request) string {
	val := r.Context().Value(requestIDKey)

	if val == nil {
		return ""
	}

	id, ok := val.(string)
	if !ok {
		return ""
	}

	return id
}
