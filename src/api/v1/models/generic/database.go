package generic

import (
	"gorm.io/gorm"
)

// struct to establish relations to other DB Models following the DB library structure
type Relationships struct{}
type RelationshipsInterf interface {
	Get_relationships() (Relationships, error)
}

// struct for overriding fields and adding tags or anything
// the DB library requires before ingestion into DB Model
type InterDataAspect struct {
	DataAspect    `gorm:"embedded"`
	Relationships `gorm:"embedded"`
}

// struct for establishing a DB Model
type DbAspect struct {
	gorm.Model
	InterDataAspect `gorm:"embedded"`
}
type DbInterface interface {
	Get_data_repr() DataAspect
}
