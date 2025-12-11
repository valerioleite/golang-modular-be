package middleware

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"libraries/http/middleware/dto"

	"github.com/google/uuid"
)

const (
	CorrelationIDKey string = "correlation_id"
	UserSubKey       string = "user_sub"
	UserEmailKey     string = "user_email"
	AccessTokenKey   string = "access_token"
)

const (
	HeaderCorrelationID string = "X-Correlation-ID"
	HeaderUserSub       string = "X-User-Sub"
	HeaderUserEmail     string = "X-User-Email"
	HeaderAuthorization string = "Authorization"
	BearerPrefix        string = "Bearer "
)

func WithRequestContext(module string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			correlationID := r.Header.Get(HeaderCorrelationID)
			if correlationID == "" {
				correlationID = uuid.New().String()
			}

			ctx := context.WithValue(r.Context(), CorrelationIDKey, correlationID)
			userSub := r.Header.Get(HeaderUserSub)
			if userSub != "" {
				ctx = context.WithValue(ctx, UserSubKey, userSub)
			}

			userEmail := r.Header.Get(HeaderUserEmail)
			if userEmail != "" {
				ctx = context.WithValue(ctx, UserEmailKey, userEmail)
			}

			authHeader := r.Header.Get(HeaderAuthorization)
			if authHeader != "" {
				accessToken := authHeader
				if len(authHeader) > len(BearerPrefix) && authHeader[:len(BearerPrefix)] == BearerPrefix {
					accessToken = authHeader[len(BearerPrefix):]
				}
				ctx = context.WithValue(ctx, AccessTokenKey, accessToken)
			}

			r = r.WithContext(ctx)
			w.Header().Set(HeaderCorrelationID, correlationID)

			wrapped := &dto.ResponseWriter{ResponseWriter: w, StatusCode: http.StatusOK}

			next.ServeHTTP(wrapped, r)
			duration := time.Since(start)

			if strings.ToLower(os.Getenv("LOGGING_REQUESTS")) != "true" {
				return
			}

			slog.Info("Request completed",
				"module", module,
				"correlation_id", correlationID,
				"method", r.Method,
				"path", r.URL.Path,
				"status", wrapped.StatusCode,
				"duration_ms", duration.Milliseconds(),
				"remote_addr", r.RemoteAddr,
			)
		})
	}
}

func GetTraceID(ctx context.Context) string {
	if traceID, ok := ctx.Value(CorrelationIDKey).(string); ok {
		return traceID
	}
	return ""
}

func GetUserSub(ctx context.Context) string {
	if userSub, ok := ctx.Value(UserSubKey).(string); ok {
		return userSub
	}
	return ""
}

func GetUserEmail(ctx context.Context) string {
	if userEmail, ok := ctx.Value(UserEmailKey).(string); ok {
		return userEmail
	}
	return ""
}

func GetAccessToken(ctx context.Context) string {
	if accessToken, ok := ctx.Value(AccessTokenKey).(string); ok {
		return accessToken
	}
	return ""
}

func NewRequestWithContextAndHeaders(ctx context.Context, method, rawURL string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, rawURL, body)
	if err != nil {
		return nil, err
	}

	addContextHeadersToRequest(ctx, req)
	return req, nil
}

func addContextHeadersToRequest(ctx context.Context, req *http.Request) {
	correlationID := GetTraceID(ctx)
	if correlationID != "" {
		req.Header.Set(HeaderCorrelationID, correlationID)
	}

	userSub := GetUserSub(ctx)
	if userSub != "" {
		req.Header.Set(HeaderUserSub, userSub)
	}

	userEmail := GetUserEmail(ctx)
	if userEmail != "" {
		req.Header.Set(HeaderUserEmail, userEmail)
	}

	accessToken := GetAccessToken(ctx)
	if accessToken != "" {
		req.Header.Set(HeaderAuthorization, BearerPrefix+accessToken)
	}
}
