package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager{
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...Middleware){
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
}

func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
		n := next

		// for i := len(middleware) - 1; i >= 0; i --{
		// 	middleware := middleware[i]
		// 	n = middleware(n)
		// }
		
		for _, middleware := range middlewares {
			n = middleware(n)
		}

		for _, globamiddleware := range mngr.globalMiddlewares {
			n = globamiddleware(n)
		}

		return n
	
}