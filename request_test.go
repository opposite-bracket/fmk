package fmk

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Name of the struct tag used in examples
type ParamSample struct {
	Sample1 string `param:"sample_1" validate:"required"`
}
type QuerySample struct {
	Sample1 string `query:"sample_1" validate:"required"`
}
type BodySample struct {
	Sample1 string `json:"sample_1" validate:"required"`
}
type HeaderSample struct {
	Sample1 string `header:"sample_1" validate:"required"`
}

func TestHeaderSample(t *testing.T) {

	ps := ParamSample{}
	qs := QuerySample{}
	bs := BodySample{}
	hs := HeaderSample{}

	router := httprouter.New()
	router.GET("/resource/:sample_1", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		processParam(ps, p)
		processHeader(hs, r.Header)
		processQuery(qs, r.URL.Query())
		processBody(bs, r.Body)
		fmt.Fprint(w, "done!\n")
	})

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/resource/hello-param?sample_1=hello-query", strings.NewReader("{\"sample_1\":\"hello-body\"}"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("sample_1", "hello-header")

	router.ServeHTTP(rec, req)

}
