module services/monolith

go 1.25.4

require (
	github.com/joho/godotenv v1.5.1
	libraries/db v0.0.0-00010101000000-000000000000
	libraries/domain v0.0.0-00010101000000-000000000000
	libraries/http v0.0.0-00010101000000-000000000000
)

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/rs/cors v1.11.1 // indirect
)

replace libraries/db => ../../libraries/db

replace libraries/domain => ../../libraries/domain

replace libraries/http => ../../libraries/http
