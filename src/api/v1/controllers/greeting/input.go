package greeting

type GreetUserInput struct {
	Name string `path:"name" description:"User name" validate:"required"`
}
