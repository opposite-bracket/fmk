package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type IApi interface {
	handler(method, url string, handler Endpoint)
	Get(url string, handler Endpoint)
	Post(url string, handler Endpoint)
	Put(url string, handler Endpoint)
	Delete(url string, handler Endpoint)
	ServeHTTP(w http.ResponseWriter, req *http.Request)
	Run() error
}

type Api struct{}
type Endpoint func(c *Context)

type api struct {
	router *httprouter.Router
}

func NewApi() *api {
	r := httprouter.New()

	return &api{r}
}

func (a *api) handler(method, url string, handler Endpoint) {
	log := ApiLog()
	a.router.Handle(method, url, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		log.StartTimer()
		log.ResetTxId()
		c := NewContext(w, r, p)
		handler(c)
		log.EndTimer()
		log.Logf("GET %s", url)
	})
}

func (a *api) Get(url string, handler Endpoint) {
	a.handler(http.MethodGet, url, handler)
}

func (a *api) Post(url string, handler Endpoint) {
	a.handler(http.MethodPost, url, handler)
}

func (a *api) Put(url string, handler Endpoint) {
	a.handler(http.MethodPut, url, handler)
}

func (a *api) Delete(url string, handler Endpoint) {
	a.handler(http.MethodDelete, url, handler)
}

func (a *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}

func (a *api) Run() error {
	ApiLog().Logf("running on %s", "http://localhost:8080")
	return http.ListenAndServe(":8080", a)
}
