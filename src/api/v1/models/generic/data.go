package generic

// data repr of the model that will remain the same regardless of which DB library is used
type DataAspect struct {
	ID uint
}

// example json representations that your API uses
// either internally or externally
type JsonRepr1 struct {
	ID uint `json:"id" binding:"required"`
}
type JsonRepr2 struct {
	ID uint `json:"id" binding:"required"`
}

// interface to enable a generic method for json parsing
type JsonRepresentations interface {
	JsonRepr1 | JsonRepr2
}

type DataAspectInterf interface {
	AsJson1() JsonRepr1
	AsJson2() JsonRepr2
	Get_relationships() (Relationships, error)
}
