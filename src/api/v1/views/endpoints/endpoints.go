package endpoints

import (
	"github.com/gin-gonic/gin"

	"src/api/v1/views/conventions"
	"src/controller"

	"github.com/martinvuyk/gadgeto/tonic"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type routerFuncType interface {
	func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
}

func register[T, A any, RT routerFuncType](routerFunc RT, url string, handler func(*gin.Context, T) (A, error), code int) {
	routerFunc(url, tonic.Handler(handler, code))
}

func Setup(router *gin.RouterGroup) {
	// Readme for tonic in https://github.com/martinvuyk/gadgeto/tree/master/tonic
	tonic.SetErrorHook(conventions.ErrHook)

	// setup endpoints
	greeting.setup(router)
	trainer.setup(router)
	pokemon.setup(router)

	c := controller.Controller{}
	router.GET("/trying", c.TryingSomethingOut)
	// setup docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
