package greeting

import (
	controller "src/api/v1/controllers/greeting"
	"src/api/v1/views/conventions"
	"src/api/v1/views/endpoints/generic"

	"github.com/gin-gonic/gin"
)

type getGreetType[T, A any] struct{ generic.Endpoint[T, A] }

// @Summary		Greeting
// @Description	get greeting with name
// @Tags		greeting
// @Produce		json
// @Success		200	{object}	any
// @Router		/hello/:name [get]
func (getGreetType[T, A]) docs(c *gin.Context) {}

type getGreetIn = *controller.GreetUserInput
type getGreetOut = *conventions.ApiResponse[controller.GreetUserOutput]
type getGreetAlias = generic.Endpoint[getGreetIn, getGreetOut]

var getGreet = getGreetType[getGreetIn, getGreetOut]{getGreetAlias{OkCode: 200}}

func (g getGreetType[T, A]) handler(c *gin.Context, in getGreetIn) (getGreetOut, error) {
	return conventions.Respond(c, in, controller.GreetAnyone, g.OkCode)
}
