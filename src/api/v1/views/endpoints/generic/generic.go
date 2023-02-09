package generic

import (
	"github.com/gin-gonic/gin"
)

type endpointInterface[T, A any] interface {
	docs(c *gin.Context)
	handler(*gin.Context, T) (A, error)
}
type Endpoint[T, A any] struct {
	OkCode int
	endpointInterface[T, A]
}

type exampleEndpoint struct {
	Endpoint[input, output]
}

// gin.HandlerFunc with docstrings in format for github.com/swaggo/swag/cmd/swag
func (exampleEndpoint) docs(*gin.Context) {}

type input = *any
type output = *any

// the real handler for the request, to be wrapped in tonic.Handler
func (exampleEndpoint) handler(*gin.Context, input) (output, error) { return nil, nil }
