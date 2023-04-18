package trainer

import (
	"src/api/v1/models"
)

func (trainer *Trainer) GetRelationships() (relationships, error) {
	var relations relationships
	search := models.Database.Model(&TrainerDB{}).First(&relations, trainer.ID)
	if err := search.Error; err != nil {
		return relationships{}, err
	}
	return relations, nil
}

func (t *TrainerDB) GetDataRepr() *Trainer {
	return &Trainer{
		ID:     t.ID,
		Name:   t.InterDataAspect.DataAspect.Name,
		Height: t.InterDataAspect.DataAspect.Height,
		Weight: t.InterDataAspect.DataAspect.Weight,
	}
}
