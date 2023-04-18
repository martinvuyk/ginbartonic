package monitoring

import (
	"src/api/router/monitoring/prometheus"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.RouterGroup) {
	prometheus.Setup(router)
}
