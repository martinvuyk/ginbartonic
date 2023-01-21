package generic

import (
	"src/api/v1/models"

	juju "github.com/juju/errors"
)

// transforms the DBAspect of an instance into its DataAspect
func (dbAspect *DbAspect[T]) Get_data_repr() *T {
	return &dbAspect.InterDataAspect.DataAspect
}

// search in the DB and return
func getFirst[T any, A any](table A, resultVar *DbAspect[T], fieldValue any) error {
	result := models.Database.Model(table).First(resultVar, fieldValue)
	if result.RowsAffected == 0 {
		return juju.NotFound
	}
	return result.Error
}

// TODO: when the language supports it, enable A as DbAspectInterf[T]
func GetByID[T any, A any](table A, id uint) (*T, error) {
	var resultVar DbAspect[T]
	err := getFirst(table, &resultVar, id)
	if err != nil {
		return nil, err
	}
	return resultVar.Get_data_repr(), nil
}

// finds and returns the instances in the DB that have a relationship to the current one
// func (data *DataAspect) Get_relationships() (Relationships, error) {
// 	var relationships Relationships
// 	search := searchBy(&relationships, &data.ID)
// 	if err := search.Error; err != nil {
// 		return Relationships{}, err
// 	}
// 	return relationships, nil
// }
