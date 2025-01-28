package domain

type Material struct {
	comment string
}

func NewMaterial() *Material {
	return &Material{}
}

func (m *Material) AddComment(comment string) {
	m.comment = comment
}

func (m *Material) HasThisComment(comment string) bool {
	return m.comment == comment
}
