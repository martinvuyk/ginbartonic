package examples

import (
	"src/api/v1/views/conventions"
	"src/api/v1/views/endpoints/examples/greeting"
	"src/api/v1/views/endpoints/examples/pokemon"
	"src/api/v1/views/endpoints/examples/trainer"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.RouterGroup, m *conventions.Metrics) {
	greeting.Setup(router, m)
	trainer.Setup(router, m)
	pokemon.Setup(router, m)
}
