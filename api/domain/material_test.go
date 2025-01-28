package domain_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/takutakuaoao/mose/api/domain"
	"testing"
)

func Test_コメントを追加する(t *testing.T) {
	material := domain.Material{}
	material.AddComment("コメント")

	assert.Equal(t, "コメント", material.Comment)
}
