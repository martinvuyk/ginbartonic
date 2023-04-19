package monitoring

import (
	"src/api/networking/monitoring/prometheus"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.RouterGroup) {
	prometheus.Setup(router)
}
