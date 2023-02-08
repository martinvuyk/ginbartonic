package greeting

import (
	"fmt"

	"github.com/gin-gonic/gin"
	juju "github.com/juju/errors"
)

// @Summary			Greet Anyone
// @Description		endpoint returns a greeting according to the given name
func GreetAnyone(c *gin.Context, in *GreetUserInput) (*GreetUserOutput, error) {
	if in.Name == "satan" {
		return nil, juju.NewForbidden(nil, "go to hell")
	}
	return &GreetUserOutput{Message: fmt.Sprintf("Hello %s!", in.Name)}, nil
}
