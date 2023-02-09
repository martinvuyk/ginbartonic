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
