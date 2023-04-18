package examples

import (
	"src/api/v1/views/endpoints/examples/greeting"
	"src/api/v1/views/endpoints/examples/pokemon"
	"src/api/v1/views/endpoints/examples/trainer"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.RouterGroup) {
	greeting.Setup(router)
	trainer.Setup(router)
	pokemon.Setup(router)
}
