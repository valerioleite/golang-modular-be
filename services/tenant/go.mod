module services/tenant

go 1.24.3

require github.com/google/uuid v1.6.0
require github.com/joho/godotenv v1.5.1

replace libraries/db => ../libraries/db
replace libraries/http => ../libraries/http
