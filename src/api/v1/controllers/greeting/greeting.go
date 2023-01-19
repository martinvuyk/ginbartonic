package greeting

import (
	"fmt"

	"github.com/gin-gonic/gin"
	juju "github.com/juju/errors"
)

func Controller(c *gin.Context, in *GreetUserInput) (*GreetUserOutput, error) {
	if in.Name == "satan" {
		return nil, juju.NewForbidden(nil, "go to hell")
	}
	return &GreetUserOutput{Message: fmt.Sprintf("Hello %s!", in.Name)}, nil
}
