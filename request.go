package fmk

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

const paramTag = "param"
const jsonTag = "json"
const queryTag = "query"
const headerTag = "header"
const validationTag = "validate"

type EType string

const (
	FieldValidation   EType = "field"
	GenericValidation       = "generic"
)

type ErrorField struct {
	EType   EType
	Message string
}

func (c *Context) ValidateBody(b interface{}) error {
	bt := reflect.TypeOf(b)

	var data map[string]interface{}
	json.NewDecoder(c.Req.Body).Decode(&data)
	//issues := map[string] string

	for i := 0; i < bt.NumField(); i++ {
		f := bt.Field(i)
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

	return nil
}

func (c *Context) ValidateHeader(h interface{}) error {
	ht := reflect.TypeOf(h)

	for i := 0; i < ht.NumField(); i++ {
		f := ht.Field(i)
		hTag := f.Tag.Get(headerTag)
		vTags := strings.Split(f.Tag.Get(validationTag), ",")

		val := c.Req.Header.Get(hTag)

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

	return nil
}

func (c *Context) ValidateQuery(q interface{}) error {
	qt := reflect.TypeOf(q)

	for i := 0; i < qt.NumField(); i++ {
		f := qt.Field(i)
		qTag := f.Tag.Get(queryTag)
		vTags := strings.Split(f.Tag.Get(validationTag), ",")

		val := c.Req.URL.Query().Get(qTag)

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

	return nil
}

func (c *Context) ValidateParam(p interface{}) error {

	pt := reflect.TypeOf(p)

	for i := 0; i < pt.NumField(); i++ {
		f := pt.Field(i)
		pTag := f.Tag.Get(paramTag)
		vTags := strings.Split(f.Tag.Get(validationTag), ",")

		val := c.Param.ByName(pTag)

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
	return nil
}
