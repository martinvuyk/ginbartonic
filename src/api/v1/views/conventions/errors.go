// Copyright [2023] [Martin Vuyk]
// https://github.com/martinvuyk/ginbartonic
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0

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

func mapJujuError(e error) int {
	err := 500
	if _, ok := e.(tonic.BindError); ok {
		err = 400
	} else {
		switch {
		case are(e, []error{juju.BadRequest, juju.NotAssigned}):
			err = 400
		case is(e, juju.Unauthorized):
			err = 401
		case is(e, juju.Forbidden):
			err = 403
		case are(e, []error{juju.NotFound, juju.UserNotFound, juju.NotProvisioned}):
			err = 404
		case are(e, []error{juju.MethodNotAllowed}):
			err = 405
		case is(e, juju.AlreadyExists):
			err = 409
		case is(e, juju.NotSupported):
			err = 415
		case is(e, juju.NotValid):
			err = 418
		case are(e, []error{juju.NotImplemented, juju.NotYetAvailable}):
			err = 501
		}
	}
	return err
}
func ErrHook(c *gin.Context, e error) (int, any) {
	errpl := e.Error()
	errcode := mapJujuError(e)
	return errcode, ApiResponse[any]{Success: false, Err: juju.ConstError(errpl), Code: uint(errcode)}
}
