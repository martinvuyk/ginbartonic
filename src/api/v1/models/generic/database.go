// Copyright [2023] [Martin Vuyk]
// https://github.com/martinvuyk/ginbartonic
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
type relationships struct{}

// struct for overriding fields and adding tags or anything
// the DB library requires before ingestion into DB Model
type interDataAspect[T any] struct {
	DataAspect    T `gorm:"embedded"`
	relationships `gorm:"embedded"`
}

// struct for establishing a DB Model
type DbAspect[T any] struct {
	gorm.Model
}

type DbAspectInterf[T any] interface {
	GetDataRepr() (T, error)
}

type exampleTable[T any] struct {
	DbAspect[T]
	interDataAspect[T] `gorm:"embedded"`
	DbAspectInterf[T]
}
