module services/authentication

go 1.25.4

require (
	github.com/coreos/go-oidc/v3 v3.17.0
	github.com/google/uuid v1.6.0
	golang.org/x/oauth2 v0.33.0
	libraries/domain v0.0.0-00010101000000-000000000000
	libraries/http v0.0.0-00010101000000-000000000000
)

require (
	github.com/auth0/go-jwt-middleware/v2 v2.3.1 // indirect
	github.com/go-jose/go-jose/v4 v4.1.3 // indirect
	github.com/rs/cors v1.11.1 // indirect
	golang.org/x/crypto v0.45.0 // indirect
	golang.org/x/sync v0.18.0 // indirect
	gopkg.in/go-jose/go-jose.v2 v2.6.3 // indirect
)

replace libraries/db => ../../libraries/db

replace libraries/domain => ../../libraries/domain

replace libraries/http => ../../libraries/http
