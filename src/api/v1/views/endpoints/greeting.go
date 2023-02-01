package endpoints

import (
	"github.com/gin-gonic/gin"

	controller "src/api/v1/controllers/greeting"
	"src/api/v1/views/conventions"
)

type greetingType struct{}

var greeting = greetingType{}

func (*greetingType) getGreeting(c *gin.Context, in *controller.GreetUserInput) (*conventions.ApiResponse[controller.GreetUserOutput], error) {
	return conventions.Respond(c, in, controller.GreetAnyone, 200)
}

func (*greetingType) setup(router *gin.Engine) {
	// endpoint returns a greeting according to the given name
	register(router.GET, "hello/:name", greeting.getGreeting, 200)
}
