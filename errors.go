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

// ApiError is the error that will be processed and formatted
// to the client on the response.
type ApiError struct {
	Service    string           `json:"Service"`
	Operation  string           `json:"Operation"`
	Category   ApiErrorCategory `json:"Category"`
	StatusCode int              `json:"-"`
	Messages   []ErrorField     `json:"Messages"`
}

func (e *ApiError) AddGenericMessage(etype EType, message string) {
	if e.Messages == nil {
		e.Messages = []ErrorField{
			{EType: etype, ETag: message},
		}
	} else {
		e.Messages = append(
			e.Messages,
			ErrorField{EType: etype, ETag: message},
		)
	}
}

func (e *ApiError) AddFieldMessage(etype EType, tag string, message string) {
	if e.Messages == nil {
		e.Messages = []ErrorField{
			{EType: etype, EField: tag, ETag: message},
		}
	} else {
		e.Messages = append(
			e.Messages,
			ErrorField{EType: etype, EField: tag, ETag: message},
		)
	}
}

func (e *ApiError) ContainsErrors() bool {
	return e.Messages != nil && len(e.Messages) >= 1
}

// Error will print out an error with all its data
// required to build said error. Contains ErrorCode &
// Service & message
func (e *ApiError) Error() string {
	return fmt.Sprintf(
		"%s:%s:%s:%v:%s",
		e.Service,
		e.Operation,
		e.Category,
		e.StatusCode,
		e.Messages,
	)
}
