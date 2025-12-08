package dto

import "time"

type ErrorResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Errors    []string  `json:"errors"`
} //@name ErrorResponse

type Error interface {
	HTTPStatus() int
	Error() string
}
