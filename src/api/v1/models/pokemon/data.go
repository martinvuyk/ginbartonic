package pokemon

type Pokemon struct {
	ID      uint
	Name    string
	Species string `json:"species" binding:"required"`
	Height  int
	Weight  int
	Types   string `json:"types" binding:"required"`
}
