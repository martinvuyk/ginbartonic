// Copyright [2023] [Martin Vuyk]
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0

package generic

import (
	"gorm.io/gorm"
)

// struct to establish relations to other DB Models following the DB library structure
type Relationships struct{}

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

type DbAspectInterf[T any, A any] interface {
	DbAspect[T, A]
	GetDataRepr() *T
}
