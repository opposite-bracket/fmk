package fmk

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

// Name of the struct tag used in examples
const paramTag = "param"
const jsonTag = "json"
const queryTag = "query"
const headerTag = "header"
const validationTag = "validate"

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

func processParam(d interface{}, ps httprouter.Params) {

	st := reflect.TypeOf(d)

	for i := 0; i < st.NumField(); i++ {
		f := st.Field(i)
		pTag := f.Tag.Get(paramTag)
		vTags := strings.Split(f.Tag.Get(validationTag), ",")

		val := ps.ByName(pTag)

		fmt.Printf(
			"%d. %v (%v), param: '%v', validation: '%v', val: '%v'\n",
			i+1,
			f.Name,
			f.Type.Name(),
			pTag,
			vTags,
			val,
		)
	}
}

func processHeader(d interface{}, h http.Header) {

	st := reflect.TypeOf(d)

	for i := 0; i < st.NumField(); i++ {
		f := st.Field(i)
		hTag := f.Tag.Get(headerTag)
		vTags := strings.Split(f.Tag.Get(validationTag), ",")

		val := h.Get(hTag)

		fmt.Printf(
			"%d. %v (%v), param: '%v', validation: '%v', val: '%v'\n",
			i+1,
			f.Name,
			f.Type.Name(),
			hTag,
			vTags,
			val,
		)
	}
}

func processBody(d interface{}, body io.ReadCloser) {
	st := reflect.TypeOf(d)

	var data map[string]interface{}
	json.NewDecoder(body).Decode(&data)

	for i := 0; i < st.NumField(); i++ {
		f := st.Field(i)
		hTag := f.Tag.Get(jsonTag)
		vTags := strings.Split(f.Tag.Get(validationTag), ",")

		val := data[hTag]

		fmt.Printf(
			"%d. %v (%v), param: '%v', validation: '%v', val: '%v'\n",
			i+1,
			f.Name,
			f.Type.Name(),
			hTag,
			vTags,
			val,
		)
	}
}

func processQuery(d interface{}, q url.Values) {

	//fmt.Fprintf(w, "hello, %s!\n", queryValues.Get("name"))
	st := reflect.TypeOf(d)

	for i := 0; i < st.NumField(); i++ {
		f := st.Field(i)
		qTag := f.Tag.Get(queryTag)
		vTags := strings.Split(f.Tag.Get(validationTag), ",")

		val := q.Get(qTag)

		fmt.Printf(
			"%d. %v (%v), param: '%v', validation: '%v', val: '%v'\n",
			i+1,
			f.Name,
			f.Type.Name(),
			qTag,
			vTags,
			val,
		)
	}
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
