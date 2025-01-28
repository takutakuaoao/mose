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
