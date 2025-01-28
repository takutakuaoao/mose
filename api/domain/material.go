package domain

type Material struct {
	Comment string
}

func (m *Material) AddComment(comment string) {
	m.Comment = comment
}
