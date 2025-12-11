#!/bin/sh

set -e

if [ -f /usr/local/kong/kong.yml.source ]; then
  AUTHENTICATION_SERVICE_URL="${AUTHENTICATION_SERVICE_URL:-http://host.docker.internal:8081}"
  USER_SERVICE_URL="${USER_SERVICE_URL:-http://host.docker.internal:8080}"

  export AUTHENTICATION_SERVICE_URL USER_SERVICE_URL
  envsubst '$AUTHENTICATION_SERVICE_URL $USER_SERVICE_URL' < /usr/local/kong/kong.yml.source > /usr/local/kong/kong.yml
fi

exec /docker-entrypoint.sh kong docker-start

