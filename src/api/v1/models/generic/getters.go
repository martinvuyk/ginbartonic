package generic

import (
	"src/api/v1/models"

	juju "github.com/juju/errors"
)

// transforms the DBAspect of an instance into its DataAspect
func (dbAspect *DbAspect[T, A]) GetDataRepr() *T {
	return &dbAspect.DataAspect
}

// search in the DB and return
func getFirst[T any, A any](table A, resultVar *DbAspect[T, A], fieldValue any) error {
	result := models.Database.Model(table).First(resultVar, fieldValue)
	if result.RowsAffected == 0 {
		return juju.NotFound
	}
	return result.Error
}

// TODO: when the language supports it, enable A as DbAspectInterf[T]
func GetByID[T any, A any](table A, id uint) (*T, error) {
	var resultVar DbAspect[T, A]
	err := getFirst(table, &resultVar, id)
	if err != nil {
		return nil, err
	}
	return resultVar.GetDataRepr(), nil
}
