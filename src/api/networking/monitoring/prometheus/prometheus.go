package prometheus

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.RouterGroup, reg *prometheus.Registry) {
	router.Any("/metrics", WrapHH(promhttp.Handler))

	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	router.Any("/metrics/custom", WrapHH(func() http.Handler { return promHandler }))
}
