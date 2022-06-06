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
		Message("something went wrong").
		Build()

	expected := "srv1:op1:ISE:400:something went wrong"

	if actual.Error() != expected {
		t.Errorf("got %s but want %s", actual, expected)
	}
}

func TestApiError_Service(t *testing.T) {
	actual := &ApiError{
		service: "some value",
	}

	expected := "updated value"
	actual.Service(expected)

	if actual.service != expected {
		t.Errorf("got %s but want %s", actual, expected)
	}
}

func TestApiError_Operation(t *testing.T) {
	actual := &ApiError{
		operation: "some value",
	}

	expected := "updated value"
	actual.Operation(expected)

	if actual.operation != expected {
		t.Errorf("got %s but want %s", actual, expected)
	}
}

func TestApiError_StatusCode(t *testing.T) {
	actual := &ApiError{
		statusCode: 404,
	}

	expected := 200
	actual.StatusCode(expected)

	if actual.statusCode != expected {
		t.Errorf("got %s but want %v", actual, expected)
	}
}

func TestApiError_Message(t *testing.T) {
	actual := &ApiError{
		message: "some value",
	}

	expected := "updated value"
	actual.Message(expected)

	if actual.message != expected {
		t.Errorf("got %s but want %s", actual, expected)
	}
}
