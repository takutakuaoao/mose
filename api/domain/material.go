package domain

type Material struct {
	Comment string
}

func NewMaterial() *Material {
	return &Material{}
}

func (m *Material) AddComment(comment string) {
	m.Comment = comment
}

func (m *Material) HasThisComment(comment string) bool {
	return m.Comment == comment
}
