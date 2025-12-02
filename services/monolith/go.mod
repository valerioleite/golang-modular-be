module services/monolith

go 1.25.4

require (
	github.com/aws/aws-sdk-go-v2 v1.40.0
	github.com/aws/aws-sdk-go-v2/config v1.32.2
	github.com/aws/aws-sdk-go-v2/credentials v1.19.2
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.20.12
	github.com/aws/aws-sdk-go-v2/service/s3 v1.92.1
	github.com/coreos/go-oidc/v3 v3.17.0
	github.com/google/uuid v1.6.0
	github.com/joho/godotenv v1.5.1
	golang.org/x/oauth2 v0.33.0
)

require (
	github.com/auth0/go-jwt-middleware/v2 v2.3.1 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.7.3 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.18.14 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.4.14 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.7.14 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.4 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.4.14 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.13.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.9.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.13.14 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.19.14 // indirect
	github.com/aws/aws-sdk-go-v2/service/signin v1.0.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.30.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.35.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.41.2 // indirect
	github.com/aws/smithy-go v1.24.0 // indirect
	github.com/go-jose/go-jose/v4 v4.1.3 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/rs/cors v1.11.1 // indirect
	golang.org/x/crypto v0.45.0 // indirect
	golang.org/x/sync v0.18.0 // indirect
	github.com/go-jose/go-jose/v4 v4.1.3 // indirect
)

replace libraries/db => ../../libraries/db
replace libraries/domain => ../../libraries/domain
replace libraries/http => ../../libraries/http
