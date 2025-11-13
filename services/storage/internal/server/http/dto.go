package http

import "time"

type ErrorResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Errors    []string  `json:"errors"`
}

type HTTPError interface {
	HTTPStatus() int
	Error() string
}

