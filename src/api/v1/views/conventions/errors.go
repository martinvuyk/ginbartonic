package conventions

import (
	"github.com/martinvuyk/gadgeto/tonic"

	"github.com/gin-gonic/gin"
	juju "github.com/juju/errors"
)

func is(base error, expected error) bool {
	if juju.Is(base, expected) {
		return true
	}
	return false
}

func are(baseError error, possibleErrors []error) bool {
	for _, expected := range possibleErrors {
		if is(baseError, expected) {
			return true
		}
	}
	return false
}

func ErrHook(c *gin.Context, e error) (int, any) {
	errcode, errpl := 500, e.Error()
	if _, ok := e.(tonic.BindError); ok {
		errcode = 400
	} else {
		switch {
		case are(e, []error{juju.BadRequest, juju.NotValid, juju.AlreadyExists, juju.NotSupported, juju.NotAssigned, juju.NotProvisioned}):
			errcode = 400
		case is(e, juju.Unauthorized):
			errcode = 401
		case is(e, juju.Forbidden):
			errcode = 403
		case are(e, []error{juju.NotFound, juju.UserNotFound}):
			errcode = 404
		case is(e, juju.MethodNotAllowed):
			errcode = 405
		case is(e, juju.NotImplemented):
			errcode = 501
		}
	}
	return errcode, ApiResponse[any]{Success: false, Err: juju.ConstError(errpl), Code: uint(errcode)}
}
