package generic

import (
	"encoding/json"

	"github.com/jinzhu/copier"
)

func copy[T JsonRepresentations](input *DataAspect) T {
	var jsonRepr T
	copier.Copy(&jsonRepr, &input)
	return jsonRepr
}

func (data *DataAspect) AsJson1() JsonRepr1 {
	return copy[JsonRepr1](data)
}
func (data *DataAspect) AsJson2() JsonRepr2 {
	return copy[JsonRepr2](data)
}

func FromJson[T JsonRepresentations](input []byte) (T, error) {
	var jsonRepr T
	err := json.Unmarshal(input, &jsonRepr)
	if err != nil {
		return jsonRepr, err
	}
	return jsonRepr, nil
}
