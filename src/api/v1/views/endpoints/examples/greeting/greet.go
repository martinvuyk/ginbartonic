package greeting

import (
	controller "src/api/v1/controllers/examples/greeting"
	"src/api/v1/views/conventions"

	"github.com/gin-gonic/gin"
)

type getGreetIn = *controller.GreetUserInput
type getGreetOut = *conventions.ApiResponse[controller.GreetUserOutput]
type getGreetEndpoint = conventions.Endpoint[getGreetIn, getGreetOut]
type getGreetType struct {
	getGreetEndpoint
}

func getGreet(m *conventions.Metrics) getGreetType {
	return getGreetType{getGreetEndpoint{OkCode: 200, Url: "/hello/:name", Metrics: m}}
}

// @Summary		Greeting
// @Description	get greeting with name
// @Tags		greeting
// @Produce		json
// @Param		name	path	string true "name"
// @Success		200	{object}	getGreetOut
// @Router		/hello/{name} [get]
func (*getGreetType) docs(c *gin.Context) {}

func (g *getGreetType) handler(c *gin.Context, in getGreetIn) (getGreetOut, error) {
	return conventions.Respond(c, in, controller.GreetAnyone, g.OkCode, g.Url, g.Metrics)
}
