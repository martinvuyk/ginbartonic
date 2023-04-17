// Copyright [2023] [Martin Vuyk]
// https://github.com/martinvuyk/ginbartonic
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0

package generic

// data repr of the model that will remain the same regardless of which DB library is used
type DataAspect[T any] struct {
	ID uint
}

// interface to enable a generic method for json parsing
type JsonRepresentations[T any] interface{}

type JsonRepr[T any] struct{}

type JsonSerializer[A any, B any] struct {
	A JsonRepr[A]
	B JsonRepr[B]
}
type JsonSerializerInterf[A any, B any] interface {
	JsonSerializer[A, B]
	FromAtoB(input A) (B, error)
	FromBtoA(input A) (B, error)
}
