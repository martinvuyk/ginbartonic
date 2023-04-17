// Copyright [2023] [Martin Vuyk]
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0

package generic

import (
	"encoding/json"

	"github.com/jinzhu/copier"
)

func FromTo[T any, A any](input *T) (*A, error) {
	var dataRepr A
	err := copier.Copy(&dataRepr, input)
	return &dataRepr, err
}

func FromJsonTo[T any, A JsonRepresentations[T]](input []byte) (*A, error) {
	var jsonRepr A
	err := json.Unmarshal(input, &jsonRepr)
	if err != nil {
		return nil, err
	}
	return &jsonRepr, nil
}
