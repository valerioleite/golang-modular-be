# Build stage
FROM golang:1.25.4-alpine AS build-stage

RUN apk add --no-cache git ca-certificates tzdata
WORKDIR /build

COPY go.work go.work.sum ./
COPY libraries/ ./libraries/
COPY services/monolith/ ./services/monolith/

WORKDIR /build/services/monolith

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/monolith ./cmd

# Runtime stage
FROM gcr.io/distroless/static-debian12:nonroot AS release-stage

WORKDIR /app

COPY --from=build-stage /app/monolith .

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app/monolith"]
