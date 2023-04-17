package router

import (
	"github.com/gin-gonic/gin"

	"src/api/router/monitoring"
	v1 "src/api/v1/views/endpoints"
)

type Endpoints struct {
	api     []*gin.Engine
	metrics []*gin.Engine
}

func getRouters() Endpoints {

	// Routes API v1
	routerV1 := gin.New()
	routerV1.Use(gin.Recovery())
	v1.Setup(routerV1.Group("/api/v1"))

	api := []*gin.Engine{routerV1}

	monitor := gin.New()
	monitor.Use(gin.Recovery())
	monitoring.Setup(monitor.Group("/api/v1"))

	monitors := []*gin.Engine{monitor}

	return Endpoints{api: api, metrics: monitors}
}
