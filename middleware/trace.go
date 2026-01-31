package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

// TraceIDKey is the context key for the trace ID
const TraceIDKey = "TraceID"

// TraceIDMiddleware injects a trace ID into the request context
func TraceIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := uuid.New().String()
		ctx := context.WithValue(r.Context(), TraceIDKey, traceID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

// GetTraceID retrieves the trace ID from the context
func GetTraceID(ctx context.Context) string {
	if v := ctx.Value(TraceIDKey); v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}
