package networking

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"

	"src/api/networking/monitoring"
	"src/api/v1/views/conventions"
	v1 "src/api/v1/views/endpoints"
)

type Endpoints struct {
	api     *gin.Engine
	metrics *gin.Engine
}

func getRoutersV1() Endpoints {
	baseUrl := "/api/v1"

	reg := prometheus.NewRegistry()
	m := conventions.NewMetrics(reg, "v1")

	routerV1 := gin.New()
	routerV1.Use(gin.Recovery())
	v1.Setup(routerV1.Group(baseUrl), m)

	monitor := gin.New()
	monitor.Use(gin.Recovery())
	monitoring.Setup(monitor.Group(baseUrl), reg)

	return Endpoints{api: routerV1, metrics: monitor}
}
