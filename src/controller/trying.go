package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ShowAccount godoc
//
//	@Summary		Show an account
//	@Description	get string by ID
//	@Tags			accounts
//	@Produce		json
//	@Success		200	{object}	any
//	@Router			/trying [get]
func (c *Controller) TryingSomethingOut(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, struct{ prueba string }{prueba: "asd1234"})
}
