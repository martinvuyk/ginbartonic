package pokemon

import "src/api/v1/models/generic"

type Pokemon struct {
	generic.DataAspect[Pokemon]
	ID   uint
	Name string
	// Species string `json:"species" binding:"required"`
	Height int
	Weight int
	// Types   string `json:"types" binding:"required"`
}
