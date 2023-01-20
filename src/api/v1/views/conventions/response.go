package conventions

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ApiResponse[A any] struct {
	Success bool  `json:"success"`
	Data    *A    `json:"data"`
	Err     error `json:"err"`
	Code    uint  `json:"code"`
}

type Handler struct{}

func Respond[T any, A any](c *gin.Context, in *T, handler func(*gin.Context, *T) (*A, error), code int) (ApiResponse[A], error) {
	resp, err := handler(c, in)
	fmt.Println(err)
	return ApiResponse[A]{Success: true, Code: uint(code), Data: resp}, err
}
