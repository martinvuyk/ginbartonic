package pokemon

type PokemonCharactOutput struct {
	ID      uint
	Name    string
	Species string `json:"species" binding:"required"`
	Types   string `json:"types" binding:"required"`
}
