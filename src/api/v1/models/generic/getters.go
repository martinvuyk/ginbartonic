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
	"src/api/v1/models"

	juju "github.com/juju/errors"
)

// search in the DB and return
func GetFirst[A any](table A, resultVar any, fieldValue any) error {
	result := models.Database.Model(table).First(resultVar, fieldValue)
	if result.RowsAffected == 0 {
		return juju.NotFound
	}
	return result.Error
}

// TODO: when the language supports it, enable A as DbAspectInterf[T]
func GetByID[T any](table *T, id uint) (*T, error) {
	var resultVar T
	err := GetFirst(table, &resultVar, id)
	if err != nil {
		return nil, err
	}
	return &resultVar, nil
}
