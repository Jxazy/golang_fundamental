package web

type CategoryCreateReuqest struct {
	Name string `validate:"required,min=1,max=200"`
}
