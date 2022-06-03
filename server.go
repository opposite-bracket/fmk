package fmk

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

type Endpoint func(c *Context) error

type Api struct {
	router *httprouter.Router
}

func NewApi() *Api {
	r := httprouter.New()

	return &Api{r}
}

func (a *Api) handler(method, url string, handler Endpoint) {
	// TODO: recover from panic
	log := ApiLog()
	a.router.Handle(method, url, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		log.StartTimer()
		log.ResetTxId()
		c := NewContext(w, r, p)
		if err := handler(c); err != nil {
			// TODO: handle error
		}
		log.EndTimer()
		log.Logf("GET %s", url)
	})
}

func (a *Api) Get(url string, handler Endpoint) {
	a.handler(http.MethodGet, url, handler)
}

func (a *Api) Post(url string, handler Endpoint) {
	a.handler(http.MethodPost, url, handler)
}

func (a *Api) Put(url string, handler Endpoint) {
	a.handler(http.MethodPut, url, handler)
}

func (a *Api) Delete(url string, handler Endpoint) {
	a.handler(http.MethodDelete, url, handler)
}

func (a *Api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}

func (a *Api) Run() error {
	ApiLog().Logf("running on %s", "http://localhost:8080")
	return http.ListenAndServe(":8080", a)
}
