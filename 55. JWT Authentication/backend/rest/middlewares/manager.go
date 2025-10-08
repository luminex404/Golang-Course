package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}
func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)

}

func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	h := next
	
    // Local Middleware Handle kore 
	//middlewares = [Arekta]
	for _, middleware := range middlewares {
		h = middleware(h)
	}	 

	return h
}


func (mngr *Manager) WrapMux(handler http.Handler, middlewares ...Middleware) http.Handler {
	h := handler
    // Global Middleware Handle kore 

	for _,globalMiddleware := range mngr.globalMiddlewares{
		h = globalMiddleware(h)
	}

	return h
}
