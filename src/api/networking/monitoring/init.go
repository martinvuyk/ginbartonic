package monitoring

import (
	prometheus_pkg "src/api/networking/monitoring/prometheus"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

func Setup(router *gin.RouterGroup, reg *prometheus.Registry) {
	prometheus_pkg.Setup(router, reg)
}
