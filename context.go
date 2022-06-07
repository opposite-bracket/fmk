package fmk

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type IContext interface {
	Json(statusCode int, body interface{})

	ValidateBody(body interface{}) error
	ValidateHeader(header interface{}) error
	ValidateQuery(query interface{}) error
	ValidateParam(param interface{}) error
	TenantDoc() *TenantDoc
}

type Context struct {
	Res   http.ResponseWriter
	Req   *http.Request
	Param httprouter.Params
}

func (c *Context) Json(statusCode int, body interface{}) {
	c.Res.Header().Set("Content-Type", "application/json")
	c.Res.WriteHeader(statusCode)
	json.NewEncoder(c.Res).Encode(body)
}

func (c *Context) TenantDoc() *TenantDoc {
	// TODO implement me
	return &TenantDoc{}
}

func NewContext(w http.ResponseWriter, r *http.Request, p httprouter.Params) *Context {
	return &Context{
		Res:   w,
		Req:   r,
		Param: p,
	}
}
