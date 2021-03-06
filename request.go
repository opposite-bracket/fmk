package fmk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

const paramTag = "param"
const bodyTag = "body"
const queryTag = "query"
const headerTag = "header"
const validationTag = "validate"

type EType string

const (
	HeaderFieldValidation EType = "headerField"
	BodyFieldValidation         = "bodyField"
	QueryFieldValidation        = "queryField"
	ParamFieldValidation        = "paramField"
	GenericValidation           = "generic"
)

type ErrorField struct {
	EType  EType  `json:"eType,omitempty"`
	EField string `json:"eField,omitempty"`
	ETag   string `json:"eTag,omitempty"`
}

func (c *Context) ValidateBody(b interface{}) error {
	bt := reflect.TypeOf(b)

	var data map[string]interface{}
	json.NewDecoder(c.Req.Body).Decode(&data)
	err := ApiError{
		Category:   RequestErrorCategory,
		StatusCode: http.StatusBadRequest,
	}

	for i := 0; i < bt.NumField(); i++ {
		f := bt.Field(i)
		hTag := f.Tag.Get(bodyTag)
		vTags := strings.Split(f.Tag.Get(validationTag), ",")

		val := data[hTag]

		for i := 0; i < len(vTags); i++ {
			switch {
			case vTags[i] == "required" && required(val):
				err.AddFieldMessage(
					BodyFieldValidation,
					hTag,
					vTags[i],
				)
			case vTags[i] == "email" && email(fmt.Sprintf("%v", val)):
				err.AddFieldMessage(
					BodyFieldValidation,
					hTag,
					vTags[i],
				)
			}
		}
	}

	if err.ContainsErrors() {
		return &err
	}

	return nil
}

func (c *Context) ValidateHeader(h interface{}) error {
	ht := reflect.TypeOf(h)
	err := ApiError{
		Category:   RequestErrorCategory,
		StatusCode: http.StatusBadRequest,
	}

	for i := 0; i < ht.NumField(); i++ {
		f := ht.Field(i)
		hTag := f.Tag.Get(headerTag)
		vTags := strings.Split(f.Tag.Get(validationTag), ",")

		val := c.Req.Header.Get(hTag)

		for i := 0; i < len(vTags); i++ {
			switch {
			case vTags[i] == "required" && required(val):
				err.AddFieldMessage(
					HeaderFieldValidation,
					hTag,
					vTags[i],
				)
			case vTags[i] == "email" && email(fmt.Sprintf("%v", val)):
				err.AddFieldMessage(
					HeaderFieldValidation,
					hTag,
					vTags[i],
				)
			}
		}
	}

	if err.ContainsErrors() {
		return &err
	}

	return nil
}

func (c *Context) ValidateQuery(q interface{}) error {
	qt := reflect.TypeOf(q)
	err := ApiError{
		Category:   RequestErrorCategory,
		StatusCode: http.StatusBadRequest,
	}

	for i := 0; i < qt.NumField(); i++ {
		f := qt.Field(i)
		qTag := f.Tag.Get(queryTag)
		vTags := strings.Split(f.Tag.Get(validationTag), ",")

		val := c.Req.URL.Query().Get(qTag)

		for i := 0; i < len(vTags); i++ {
			switch {
			case vTags[i] == "required" && required(val):
				err.AddFieldMessage(
					QueryFieldValidation,
					qTag,
					vTags[i],
				)
			case vTags[i] == "email" && email(fmt.Sprintf("%v", val)):
				err.AddFieldMessage(
					QueryFieldValidation,
					qTag,
					vTags[i],
				)
			}
		}
	}

	if err.ContainsErrors() {
		return &err
	}

	return nil
}

func (c *Context) ValidateParam(p interface{}) error {
	pt := reflect.TypeOf(p)
	err := ApiError{
		Category:   RequestErrorCategory,
		StatusCode: http.StatusBadRequest,
	}

	for i := 0; i < pt.NumField(); i++ {
		f := pt.Field(i)
		pTag := f.Tag.Get(paramTag)
		vTags := strings.Split(f.Tag.Get(validationTag), ",")

		val := c.Param.ByName(pTag)

		for i := 0; i < len(vTags); i++ {
			switch {
			case vTags[i] == "required" && required(val):
				err.AddFieldMessage(
					ParamFieldValidation,
					pTag,
					vTags[i],
				)
			case vTags[i] == "email" && email(fmt.Sprintf("%v", val)):
				err.AddFieldMessage(
					ParamFieldValidation,
					pTag,
					vTags[i],
				)
			}
		}
	}

	if err.ContainsErrors() {
		return &err
	}

	return nil
}
