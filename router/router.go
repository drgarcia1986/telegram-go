package router

import (
	"errors"
	"regexp"
)

type route struct {
	pattern *regexp.Regexp
	handler Handler
}

type RegexRouter struct {
	routes []*route
}

func matchMesage(message string, pattern *regexp.Regexp) map[string]string {
	match := pattern.FindStringSubmatch(message)
	if match == nil {
		return nil
	}
	matchValues := make(map[string]string)
	for i, name := range pattern.SubexpNames() {
		if i != 0 {
			matchValues[name] = match[i]
		}
	}
	return matchValues
}

func (r *RegexRouter) Register(pattern string, handler Handler) {
	re := regexp.MustCompile(pattern)
	r.routes = append(r.routes, &route{re, handler})
}

func (r RegexRouter) Resolve(message string) (string, error) {
	for _, route := range r.routes {
		matchValues := matchMesage(message, route.pattern)
		if matchValues != nil {
			return route.handler(message, matchValues)
		}
	}
	return "", errors.New("Not found valid route for this message")
}

func New() *RegexRouter {
	var routes []*route
	return &RegexRouter{routes}
}
