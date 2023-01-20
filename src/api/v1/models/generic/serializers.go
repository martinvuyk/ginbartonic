package generic

import (
	"encoding/json"

	"github.com/jinzhu/copier"
)

func Copy[T any, A any](input *A) (T, error) {
	var jsonRepr T
	err := copier.Copy(&jsonRepr, &input)
	return jsonRepr, err
}

func (data *DataAspect) AsJson1() (JsonRepr1, error) {
	return Copy[JsonRepr1](data)
}
func (data *DataAspect) AsJson2() (JsonRepr2, error) {
	return Copy[JsonRepr2](data)
}

func FromJson[T JsonRepresentations](input []byte) (T, error) {
	var jsonRepr T
	err := json.Unmarshal(input, &jsonRepr)
	if err != nil {
		return jsonRepr, err
	}
	return jsonRepr, nil
}
