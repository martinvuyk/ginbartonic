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
type InterDataAspect[T any] struct {
	DataAspect    T `gorm:"embedded"`
	Relationships `gorm:"embedded"`
}

// struct for establishing a DB Model
type DbAspect[T any, A any] struct {
	gorm.Model
	InterDataAspect[T] `gorm:"embedded"`
}

// type DbAspectInterf[T any, A any] interface {
// 	*DbAspect[T, A]
// 	Get_data_repr() *T
// }
