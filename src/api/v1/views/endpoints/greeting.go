package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/martinvuyk/gadgeto/tonic"

	controller "src/api/v1/controllers/greeting"
	"src/api/v1/views/conventions"
)

type greetingType struct{}

var greeting = greetingType{}

type getGreetingType struct {
	successCode int
}

var getGreeting = getGreetingType{successCode: 200}

// ShowAccount godoc
//
//	@Summary		Greeting
//	@Description	get greeting with name
//	@Tags			greeting
//	@Produce		json
//	@Success		200	{object}	any
//	@Router			/hello/:name [get]
func (*getGreetingType) docs(c *gin.Context) {}
func (g *getGreetingType) handler(c *gin.Context, in *controller.GreetUserInput) (*conventions.ApiResponse[controller.GreetUserOutput], error) {
	return conventions.Respond(c, in, controller.GreetAnyone, g.successCode)
}
func (g *getGreetingType) wrappedHandler() gin.HandlerFunc {
	return tonic.Handler(g.handler, g.successCode)
}
func (*greetingType) setup(router *gin.RouterGroup) {
	router.GET("hello/:name", getGreeting.docs, getGreeting.wrappedHandler())
	// register(router.GET, "hello/:name", greeting.getGreeting, 200)
}
