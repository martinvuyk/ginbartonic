package endpoints

import (
	"github.com/gin-gonic/gin"

	"src/api/v1/controllers/greeting"
	"src/api/v1/views/conventions"
)

func (*Endpoints) greeting(c *gin.Context, in *greeting.GreetUserInput) (conventions.ApiResponse[greeting.GreetUserOutput], error) {
	return conventions.Respond(c, in, greeting.Controller, 200)
}