package middleware

import "net/http"

type middleware func(http.Handler) http.Handler

// middlware manager to manage middlwares.
type Manager struct {
	globalMiddlewares []middleware
}

// return a new manager struct pointer.
func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]middleware, 0),
	}
}

// add global middlwares.
func (mngr *Manager) Use(mdlws ...middleware) {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, mdlws...)
}

// take a http.Handler and wraped it with global middlewares.
// follow FIFO(first in first out) approach
func (mngr *Manager) GlobalWraper(h http.Handler) http.Handler {
	handler := h

	for idx := len(mngr.globalMiddlewares) - 1; idx >= 0; idx-- {
		handler = mngr.globalMiddlewares[idx](handler)
	}

	return handler
}

// take a http.Handler and wraped it with local middlewares.
// follow FIFO(first in first out) approach
func (mngr *Manager) With(h http.Handler, mdlws ...middleware) http.Handler {
	handler := h

	for idx := len(mdlws) - 1; idx >= 0; idx-- {
		handler = mdlws[idx](handler)
	}

	return handler
}
