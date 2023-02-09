package generic

// data repr of the model that will remain the same regardless of which DB library is used
type DataAspect[T any] struct {
	ID uint
}

// interface to enable a generic method for json parsing
type JsonRepresentations[T any] interface{}

type JsonRepr[T any] struct{}

type JsonSerializer[A any, B any] struct {
	A JsonRepr[A]
	B JsonRepr[B]
}
type JsonSerializerInterf[A any, B any] interface {
	JsonSerializer[A, B]
	FromAtoB(input A) (B, error)
	FromBtoA(input A) (B, error)
}
