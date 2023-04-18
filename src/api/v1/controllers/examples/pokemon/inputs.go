package pokemon

type PokemonCharactInput struct {
	ID int `path:"id" description:"Pokemon ID" validate:"required"`
}
