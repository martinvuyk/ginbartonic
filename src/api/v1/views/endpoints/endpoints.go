package endpoints

import (
	"github.com/gin-gonic/gin"

	"src/api/v1/views/conventions"

	"github.com/martinvuyk/gadgeto/tonic"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup(router *gin.RouterGroup) {
	// Readme for tonic in https://github.com/martinvuyk/gadgeto/tree/master/tonic
	// middleware and http error handler
	tonic.SetErrorHook(conventions.ErrHook)

	// setup endpoints

	// setup docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
