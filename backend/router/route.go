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
	callback     routeFunction
}

func splitPath(path, method string) []string {
	nestedPaths := strings.Split(path, "/")
	nestedPaths[0] = method
	return nestedPaths
}

func (r *route) Build(nestedPaths []string, f routeFunction) string {
	if len(nestedPaths) == 0 {
		r.callback = f
	} else {

		if nestedRoute, ok := r.nestedRoutes[nestedPaths[0]]; ok {
			nestedRoute.Build(nestedPaths[1:], f)
		} else {
			nestedRoute := &route{
				Name:         nestedPaths[0],
				nestedRoutes: make(map[string]*route),
			}
			if strings.HasPrefix(nestedRoute.Name, ":") {
				nestedRoute.makePathParam()
			}

			nestedRoute.Build(nestedPaths[1:], f)
			r.nestedRoutes[nestedRoute.Name] = nestedRoute
		}
	}
	return r.Name
}

func (r *route) Match(nestedPaths []string, w http.ResponseWriter, req *http.Request, c *Context) {

	nestedRoute, ok := r.nestedRoutes[nestedPaths[0]]
	if !ok {
		nestedRoute, ok = r.nestedRoutes["param"]
	}
	if ok {
		if nestedRoute.isPathParam {
			ok := c.addPathParam(nestedRoute.paramName, nestedPaths[0], nestedRoute.paramIsInt)
			if !ok {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
		if len(nestedPaths) == 1 {
			nestedRoute.callback(w, req, c)
		} else {
			nestedRoute.Match(nestedPaths[1:], w, req, c)
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
