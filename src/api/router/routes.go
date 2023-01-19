package router

import (
	"github.com/gin-gonic/gin"

	v1 "src/api/v1/views/endpoints"
)

func getRouters() []*gin.Engine {

	// Routes API v1
	routerV1 := gin.New()
	routerV1.Use(gin.Recovery())
	v1.Setup(routerV1)

	routers := []*gin.Engine{routerV1}
	return routers
}
