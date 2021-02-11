package router

import "strconv"

type Context struct {
	Keys map[string]interface{}
}

func (c *Context) addPathParam(r *route, value string) bool {
	if r.isIntParam {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		c.Keys[r.paramName] = intValue
	} else {
		c.Keys[r.paramName] = value
	}
	return true
}
