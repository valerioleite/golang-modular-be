package http

import "time"

type ErrorResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Errors    []string  `json:"errors"`
}

type Error interface {
	HTTPStatus() int
	Error() string
}
