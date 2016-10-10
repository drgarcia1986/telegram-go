package server

import (
	"testing"

	"github.com/drgarcia1986/telegram-go/router"
)

type fakeRouter struct {
	lastPattern         string
	lastToResolveString string
}

func (fr *fakeRouter) Register(p string, h router.Handler) {
	fr.lastPattern = p
}

func (fr *fakeRouter) Resolve(s string) (string, error) {
	fr.lastToResolveString = s
	return "", nil
}

func TestRegisterHandler(t *testing.T) {
	fakeRouter := &fakeRouter{}

	server := &Server{router: fakeRouter}
	server.HandleFunc(
		"foo", func(string, map[string]string) (string, error) {
			return "", nil
		})

	if fakeRouter.lastPattern != "foo" {
		t.Errorf("Expecter 'foo' as register pattern instead of '%s'", fakeRouter.lastPattern)
	}
}
