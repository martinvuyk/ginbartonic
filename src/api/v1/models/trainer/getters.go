package trainer

import (
	"src/api/v1/models"
)

func (trainer *Trainer) Get_relationships() (Relationships, error) {
	var relationships Relationships
	search := models.Database.Model(&TrainerDB{}).First(&relationships, trainer.ID)
	if err := search.Error; err != nil {
		return Relationships{}, err
	}
	return relationships, nil
}

func (trainer *TrainerDB) Get_data_repr() Trainer {
	return trainer.Trainer.Trainer
}
