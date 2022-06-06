package fmk

import "fmt"

// ApiErrorCategory describes the type of error
// thrown by the api
type ApiErrorCategory string

const (
	InternalServerErrorCategory ApiErrorCategory = "ISE"
	RequestErrorCategory                         = "RQE"
	DBErrorCategory                              = "DBE"
	CacheErrorCategory                           = "CCE"
	ServiceErrorCategory                         = "SVE"
	UtilityErrorCategory                         = "UTE"
)

// ApiErrorBuilder helps when creating a new error
// the api needs to trigger
type ApiErrorBuilder struct {
	service    string
	operation  string
	category   ApiErrorCategory
	statusCode int
	message    string
}

// ApiError is the error that will be processed and formatted
// to the client on the response.
type ApiError struct {
	service    string
	operation  string
	category   ApiErrorCategory
	statusCode int
	message    string
}

func (e *ApiError) Service(service string) {
	e.service = service
}

func (e *ApiError) StatusCode(statusCode int) {
	e.statusCode = statusCode
}

func (e *ApiError) Operation(operation string) {
	e.operation = operation
}

func (e *ApiError) Message(message string) {
	e.message = message
}

// Error will print out an error with all its data
// required to build said error. Contains ErrorCode &
// service & message
func (e *ApiError) Error() string {
	return fmt.Sprintf(
		"%s:%s:%s:%v:%s",
		e.service,
		e.operation,
		e.category,
		e.statusCode,
		e.message,
	)
}

// Service assigns value to error, so we can identify
// which service the error was triggered from
func (e *ApiErrorBuilder) Service(service string) *ApiErrorBuilder {
	e.service = service
	return e
}

// Operation assigns value to error, so we can identify
// which operation the error was triggered from
func (e *ApiErrorBuilder) Operation(operation string) *ApiErrorBuilder {
	e.operation = operation
	return e
}

// Category assigns value to error, so we can identify
// the type of error that was triggered. See ApiErrorCategory
// for valid options
func (e *ApiErrorBuilder) Category(category ApiErrorCategory) *ApiErrorBuilder {
	e.category = category
	return e
}

// StatusCode assigns an http status code to retrieve to teh client
func (e *ApiErrorBuilder) StatusCode(statusCode int) *ApiErrorBuilder {
	e.statusCode = statusCode
	return e
}

// Message assigns value to error, so we can have more detailed
// information on the error that was triggered
func (e *ApiErrorBuilder) Message(message string) *ApiErrorBuilder {
	e.message = message
	return e
}

// NewErrorBuilder will generate an empty instance
// for further configuration of an API error
func NewErrorBuilder() *ApiErrorBuilder {
	return &ApiErrorBuilder{}
}

// Build validates and retrieves the ApiError that will
// need to be trickled down to the API client.
func (e *ApiErrorBuilder) Build() *ApiError {

	if e.category == "" {
		panic("category is required when creating an ApiError")
	}

	if e.message == "" {
		panic("message is required when creating an ApiError")
	}

	return &ApiError{
		operation:  e.operation,
		category:   e.category,
		service:    e.service,
		statusCode: e.statusCode,
		message:    e.message,
	}
}
