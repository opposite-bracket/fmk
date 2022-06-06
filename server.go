package fmk

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type IServer interface {
	handler(method, url string, handler Endpoint)
	Get(url string, handler Endpoint)
	Post(url string, handler Endpoint)
	Put(url string, handler Endpoint)
	Delete(url string, handler Endpoint)
	ServeHTTP(w http.ResponseWriter, req *http.Request)
	Run() error
}

type Endpoint func(c *Context) error

type Server struct {
	router *httprouter.Router
}

func NewServer() *Server {
	r := httprouter.New()

	return &Server{r}
}

func (a *Server) handler(method, url string, handler Endpoint) {
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

func (a *Server) Get(url string, handler Endpoint) {
	a.handler(http.MethodGet, url, handler)
}

func (a *Server) Post(url string, handler Endpoint) {
	a.handler(http.MethodPost, url, handler)
}

func (a *Server) Put(url string, handler Endpoint) {
	a.handler(http.MethodPut, url, handler)
}

func (a *Server) Delete(url string, handler Endpoint) {
	a.handler(http.MethodDelete, url, handler)
}

func (a *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}

func (a *Server) Run() error {
	ApiLog().Logf("running on %s", "http://localhost:8080")
	return http.ListenAndServe(":8080", a)
}
