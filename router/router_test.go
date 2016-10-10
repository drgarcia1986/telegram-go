package router

import (
	"fmt"
	"testing"
)

func TestRegisterRoute(t *testing.T) {
	regexRouter := New()
	regexRouter.Register("foo", func(string, map[string]string) (string, error) {
		return "Ok!", nil
	})
	regexRouter.Register("bar", func(string, map[string]string) (string, error) {
		return "Golang is awesome", nil
	})

	if len(regexRouter.routes) != 2 {
		t.Error("Expected at last two route")
	}

	handlerResult, _ := regexRouter.routes[0].handler("", make(map[string]string))
	if handlerResult != "Ok!" {
		t.Errorf("Expected 'Ok!' in handler result msg, but it was %s instead", handlerResult)
	}
}

func TestResolveValidRoute(t *testing.T) {
	rePattern := "Hi I'm (?P<name>\\w+)"
	handler := func(m string, mValues map[string]string) (string, error) {
		return fmt.Sprintf("Hello %s", mValues["name"]), nil
	}

	regexRouter := New()
	regexRouter.Register(rePattern, handler)

	result, _ := regexRouter.Resolve("Hi I'm Diego")
	if result != "Hello Diego" {
		t.Errorf("Expected 'Hello Diego', but it was %s", result)
	}
}

func TestResolveValidRouteWithoutMatchValues(t *testing.T) {
	rePattern := "Hi"
	expectedResult := "Hello"
	handler := func(m string, mValues map[string]string) (string, error) {
		return expectedResult, nil
	}

	regexRouter := New()
	regexRouter.Register(rePattern, handler)

	result, _ := regexRouter.Resolve(rePattern)
	if result != expectedResult {
		t.Errorf("Expected '%s', but it was %s", expectedResult, result)
	}
}

func TestNonMatchingRoute(t *testing.T) {
	regexRouter := New()
	regexRouter.Register("invalid", func(string, map[string]string) (string, error) {
		return "Ok!", nil
	})

	_, err := regexRouter.Resolve("Some cool message")
	if err == nil {
		t.Errorf("Expected some error for a non matching message")
	}
}
