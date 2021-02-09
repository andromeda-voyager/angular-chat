package router

import (
	"net/http"
)

type Middleware func(w http.ResponseWriter, r *http.Request, c *Context) bool

type Group struct {
	middleware []Middleware
}

func NewGroup() *Group {
	group := &Group{
		middleware: []Middleware{},
	}
	return group
}

func (g *Group) Use(m Middleware) {
	g.middleware = append(g.middleware, m)
}

func (g *Group) runMiddleware(w http.ResponseWriter, req *http.Request, c *Context) bool {
	for _, m := range g.middleware {
		ok := m(w, req, c)
		if !ok {
			return false
		}
	}
	return true
}
