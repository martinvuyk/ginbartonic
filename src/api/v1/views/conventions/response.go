package conventions

import (
	"github.com/gin-gonic/gin"
)

type ApiResponse[A any] struct {
	Success bool  `json:"success"`
	Data    *A    `json:"data"`
	Err     error `json:"err"`
	Code    uint  `json:"code"`
}

type HandlerType[T, A any] interface {
	func(*gin.Context, *T) (*A, error)
}

type ResponseType[T, A any, D HandlerType[T, A]] interface {
	func(c *gin.Context, in *T, handler D, code int) (*ApiResponse[A], error)
}

func Respond[T any, A any, D HandlerType[T, A]](c *gin.Context, in *T, handler D, code int) (*ApiResponse[A], error) {
	resp, err := handler(c, in)
	return &ApiResponse[A]{Success: true, Code: uint(code), Data: resp}, err
}
