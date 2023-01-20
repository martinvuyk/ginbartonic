package generic

import (
	"gorm.io/gorm"

	"src/api/v1/models"

	juju "github.com/juju/errors"
)

// transforms the DBAspect of an instance into its DataAspect
func (thing *DbAspect) Get_data_repr() DataAspect {
	return thing.InterDataAspect.DataAspect
}

// search in the DB and return
func searchBy(resultVar any, fieldValue any) *gorm.DB {
	return models.Database.Model(&DbAspect{}).First(&resultVar, fieldValue)
}

func GetByID[T any, A DbAspectInterface[T]](id uint) (*T, error) {
	var resultVar A
	result := models.Database.First(&resultVar, id)
	amount, err := result.RowsAffected, result.Error
	if err != nil || amount == 0 {
		return nil, juju.NotFound
	}
	return resultVar.Get_data_repr(), nil
}

// finds and returns the instances in the DB that have a relationship to the current one
func (data *DataAspect) Get_relationships() (Relationships, error) {
	var relationships Relationships
	search := searchBy(&relationships, &data.ID)
	if err := search.Error; err != nil {
		return Relationships{}, err
	}
	return relationships, nil
}
