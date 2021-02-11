package router

import (
	"net/http"
	"strings"
)

type route struct {
	name         string
	paramName    string
	isIntParam   bool
	isPathParam  bool
	nestedRoutes map[string]*route
	callback     routeCallbackFunc
	group        *Group
}

func splitPath(path, method string) []string {
	nestedPaths := strings.Split(path, "/")
	nestedPaths[0] = method
	return nestedPaths
}

func (r *route) Build(pathSegments []string, fn routeCallbackFunc, g *Group) {
	if len(pathSegments) == 0 {
		r.callback = fn
		r.group = g
	} else {
		nestedRoute, ok := r.getNestedRoute(pathSegments[0])
		if ok {
			nestedRoute.Build(pathSegments[1:], fn, g)
		} else {
			nestedRoute := &route{
				name:         pathSegments[0],
				nestedRoutes: make(map[string]*route),
			}
			if strings.HasPrefix(nestedRoute.name, ":") {
				nestedRoute.makePathParam()
			}

			nestedRoute.Build(pathSegments[1:], fn, g)
			r.nestedRoutes[nestedRoute.name] = nestedRoute
		}
	}
}

func (r *route) Match(pathSegments []string, w http.ResponseWriter, req *http.Request, c *Context) {
	if len(pathSegments) == 0 {
		ok := r.group.runMiddleware(w, req, c)
		if ok {
			if r.callback != nil {
				r.callback(w, req, c)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
			return
		}
	} else {
		nestedRoute, ok := r.getNestedRoute(pathSegments[0])
		if ok {
			if nestedRoute.isPathParam {
				ok := c.addPathParam(nestedRoute, pathSegments[0])
				if !ok {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
			}
			nestedRoute.Match(pathSegments[1:], w, req, c)

		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func (r *route) getNestedRoute(pathSegment string) (*route, bool) {
	nestedRoute, ok := r.nestedRoutes[pathSegment]
	if !ok {
		nestedRoute, ok = r.nestedRoutes["param"]
	}
	return nestedRoute, ok
}

func (r *route) makePathParam() {
	r.isPathParam = true
	if strings.HasSuffix(r.name, "<int>") {
		r.isIntParam = true
	}
	r.paramName = strings.TrimSuffix(r.name[1:], "<int>")
	r.name = "param"
}
