package middleware

type Middleware struct {
}

// return a new middleware struct pointer.
func NewMiddleware() *Middleware {
	return &Middleware{}
}
