// Copyright [2023] [Martin Vuyk]
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0

package generic

import (
	"github.com/gin-gonic/gin"
	"github.com/martinvuyk/gadgeto/tonic"
)

type endpointInterface[T, A any] interface {
	docs(c *gin.Context)
	handler(*gin.Context, T) (A, error)
}
type Endpoint[T, A any] struct {
	OkCode int
	endpointInterface[T, A]
}

type exampleEndpoint struct {
	Endpoint[input, output]
}

// gin.HandlerFunc with docstrings in format for github.com/swaggo/swag/cmd/swag
func (*exampleEndpoint) docs(*gin.Context) {}

type input = *any
type output = *any

// the real handler for the request, to be wrapped in tonic.Handler
func (*exampleEndpoint) handler(*gin.Context, input) (output, error) { return nil, nil }

func Setup(router *gin.RouterGroup) {
	exmp := exampleEndpoint{}
	router.GET("example", exmp.docs, tonic.Handler(exmp.handler, exmp.OkCode))
}
