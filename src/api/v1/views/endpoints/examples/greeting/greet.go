package greeting

import (
	controller "src/api/v1/controllers/examples/greeting"
	"src/api/v1/views/conventions"
	"src/api/v1/views/endpoints/generic"

	"github.com/gin-gonic/gin"
)

type getGreetIn = *controller.GreetUserInput
type getGreetOut = *conventions.ApiResponse[controller.GreetUserOutput]
type getGreetEndpoint = generic.Endpoint[getGreetIn, getGreetOut]
type getGreetType struct {
	getGreetEndpoint
}

var getGreet = getGreetType{getGreetEndpoint{OkCode: 200}}

// @Summary		Greeting
// @Description	get greeting with name
// @Tags		greeting
// @Produce		json
// @Success		200	{object}	getGreetOut
// @Router		/hello/:name [get]
func (*getGreetType) docs(c *gin.Context) {}

func (g *getGreetType) handler(c *gin.Context, in getGreetIn) (getGreetOut, error) {
	return conventions.Respond(c, in, controller.GreetAnyone, g.OkCode)
}
