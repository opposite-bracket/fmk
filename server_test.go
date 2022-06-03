package fmk

import (
	"net/http"
	"testing"
)

func TestApi_New(t *testing.T) {
	server := NewApi()

	server.Get(
		http.MethodGet,
		"/hello-world",
		func(c context) {

		},
	)
}
