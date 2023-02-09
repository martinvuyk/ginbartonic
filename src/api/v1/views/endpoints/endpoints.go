package endpoints

import (
	"github.com/gin-gonic/gin"

	"src/api/v1/views/conventions"

	"github.com/martinvuyk/gadgeto/tonic"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"src/api/v1/views/endpoints/greeting"
	"src/api/v1/views/endpoints/pokemon"
	"src/api/v1/views/endpoints/trainer"
)

func Setup(router *gin.RouterGroup) {
	// Readme for tonic in https://github.com/martinvuyk/gadgeto/tree/master/tonic
	// middleware and http error handler
	tonic.SetErrorHook(conventions.ErrHook)

	// setup endpoints
	greeting.Setup(router)
	trainer.Setup(router)
	pokemon.Setup(router)

	// setup docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
