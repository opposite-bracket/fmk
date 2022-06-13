package fmk

import (
	"net/http"
	"testing"
)

func TestErrorBuilder(t *testing.T) {
	actual := NewErrorBuilder().
		Service("srv1").
		Operation("op1").
		Category(InternalServerErrorCategory).
		StatusCode(http.StatusBadRequest).
		AddGenericMessage(GenericValidation, "something went wrong").
		Build()

	expected := "srv1:op1:ISE:400:something went wrong"

	if actual.Error() != expected {
		t.Errorf("got %s but shouldFail %s", actual, expected)
	}
}

func TestApiError_Service(t *testing.T) {
	actual := &ApiError{
		Service: "some value",
	}

	expected := "updated value"
	actual.Service(expected)

	if actual.Service != expected {
		t.Errorf("got %s but shouldFail %s", actual, expected)
	}
}

func TestApiError_Operation(t *testing.T) {
	actual := &ApiError{
		Operation: "some value",
	}

	expected := "updated value"
	actual.Operation(expected)

	if actual.Operation != expected {
		t.Errorf("got %s but shouldFail %s", actual, expected)
	}
}

func TestApiError_StatusCode(t *testing.T) {
	actual := &ApiError{
		StatusCode: 404,
	}

	expected := 200
	actual.StatusCode(expected)

	if actual.StatusCode != expected {
		t.Errorf("got %s but shouldFail %v", actual, expected)
	}
}

func TestApiError_Message(t *testing.T) {
	expected := ErrorField{
		GenericValidation, "some value",
	}
	actual := &ApiError{
		Messages: []ErrorField{
			{expected.EType, expected.EMessage},
		},
	}
	actual.AddGenericMessage(GenericValidation, expected.EMessage)

	if actual.Messages[0] != expected {
		t.Errorf(
			"got %s but shouldFail %s", actual.Messages[0],
			expected,
		)
	}
}
