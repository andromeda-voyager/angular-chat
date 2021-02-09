package router

import (
	"net/http"
	"strconv"
	"strings"
)

type route struct {
	Name         string
	paramName    string
	paramIsInt   bool
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

		if nestedRoute, ok := r.nestedRoutes[pathSegments[0]]; ok {
			nestedRoute.Build(pathSegments[1:], fn, g)
		} else {
			nestedRoute := &route{
				Name:         pathSegments[0],
				nestedRoutes: make(map[string]*route),
			}
			if strings.HasPrefix(nestedRoute.Name, ":") {
				nestedRoute.makePathParam()
			}

			nestedRoute.Build(pathSegments[1:], fn, g)
			r.nestedRoutes[nestedRoute.Name] = nestedRoute
		}
	}
}

func (r *route) Match(pathSegments []string, w http.ResponseWriter, req *http.Request, c *Context) {

	nestedRoute, ok := r.nestedRoutes[pathSegments[0]]
	if !ok {
		nestedRoute, ok = r.nestedRoutes["param"]
	}
	if ok {
		if nestedRoute.isPathParam {
			ok := c.addPathParam(nestedRoute.paramName, pathSegments[0], nestedRoute.paramIsInt)
			if !ok {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
		if len(pathSegments) == 1 {
			ok := nestedRoute.group.runMiddleware(w, req, c)
			if ok {
				nestedRoute.callback(w, req, c)
			}
		} else {
			nestedRoute.Match(pathSegments[1:], w, req, c)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (r *route) makePathParam() {
	r.isPathParam = true
	if strings.HasSuffix(r.Name, "<int>") {
		r.paramIsInt = true
	}
	r.paramName = strings.TrimSuffix(r.Name[1:], "<int>")
	r.Name = "param"
}

func (c *Context) addPathParam(name, value string, isInt bool) bool {
	if isInt {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		c.Keys[name] = intValue
	} else {
		c.Keys[name] = value
	}
	return true
}
