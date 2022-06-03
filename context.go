package fmk

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type IContext interface {
	Json(statusCode int, body interface{})
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

func NewContext(w http.ResponseWriter, r *http.Request, p httprouter.Params) *Context {
	return &Context{
		Res:   w,
		Req:   r,
		Param: p,
	}
}
