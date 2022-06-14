package fmk

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Sample struct {
	Field1 string `json:"field1"`
}

func TestName(t *testing.T) {
	expected := "{\"field1\":\"Hello\"}\n"
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	c := NewContext(rec, req, nil)

	c.Json(http.StatusOK, Sample{
		Field1: "Hello",
	})

	assert.EqualValues(t, rec.Body.String(), expected)
}
