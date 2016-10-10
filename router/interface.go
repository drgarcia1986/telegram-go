package router

type Router interface {
	Register(string, Handler)
	Resolve(string) (string, error)
}

type Handler func(string, map[string]string) (string, error)
