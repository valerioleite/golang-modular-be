package middleware

import (
	httpLib "libraries/http"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

func HandleWithValidateToken(handlerFunc http.HandlerFunc) http.Handler {
	issuerURL, err := url.Parse(os.Getenv("OIDC_ISSUER_URL"))
	if err != nil {
		slog.Error("Failed to parse the issuer url", "error", err)
		os.Exit(1)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)
	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{os.Getenv("OIDC_CLIENT_ID")},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
	)

	if err != nil {
		slog.Error("Failed to set up the jwt validator", "error", err)
		os.Exit(1)
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		slog.Error("Failed to validate JWT", "error", err)
		httpLib.HandleErrorWithStatus(w, http.StatusUnauthorized, "failed to validate JWT")
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
	)

	return middleware.CheckJWT(handlerFunc)
}
