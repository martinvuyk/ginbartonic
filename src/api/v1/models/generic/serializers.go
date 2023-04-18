// Copyright [2023] [Martin Vuyk]
// https://github.com/martinvuyk/ginbartonic
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0

package generic

import (
	"github.com/jinzhu/copier"
)

type Some = struct{}

func FromTo[T any, A any](input *T, dataRepr *A) (*A, error) {
	err := copier.Copy(dataRepr, input)
	return dataRepr, err
}
