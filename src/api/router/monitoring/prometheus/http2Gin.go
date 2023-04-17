// credits to
// https://github.com/turtlemonvh/gin-wraphh

package prometheus

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// A wrapper that turns a http.ResponseWriter into a gin.ResponseWriter, given an existing gin.ResponseWriter
// Needed if the middleware you are using modifies the writer it passes downstream
// FIXME: Wrap more methods: https://golang.org/pkg/net/http/#ResponseWriter
type WrappedResponseWriter struct {
	gin.ResponseWriter
	writer http.ResponseWriter
}

func (w *WrappedResponseWriter) Write(data []byte) (int, error) {
	return w.writer.Write(data)
}

func (w *WrappedResponseWriter) WriteString(s string) (n int, err error) {
	return w.writer.Write([]byte(s))
}

type NextRequestHandler struct {
	c *gin.Context
}

// Run the next request in the middleware chain and return
// See: https://godoc.org/github.com/gin-gonic/gin#Context.Next
func (h *NextRequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.c.Writer = &WrappedResponseWriter{h.c.Writer, w}
	h.c.Next()
}

// Wrap something that accepts an http.Handler, returns an http.Handler
func WrapHH(hh func() http.Handler) gin.HandlerFunc {
	// func WrapHH(hh func(h http.Handler) http.Handler) gin.HandlerFunc {
	// Steps:
	// - create an http handler to pass `hh`
	// - call `hh` with the http handler, which returns a function
	// - call the ServeHTTP method of the resulting function to run the rest of the middleware chain

	return func(c *gin.Context) {
		// hh(&nextRequestHandler{c}).ServeHTTP(c.Writer, c.Request)
		hh().ServeHTTP(c.Writer, c.Request)
	}
}
