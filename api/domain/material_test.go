package domain_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/takutakuaoao/mose/api/domain"
	"testing"
)

func Test_コメントを追加する(t *testing.T) {
	material := domain.NewMaterial()
	material.AddComment("コメント")

	assert.True(t, material.HasThisComment("コメント"))
}

func Test_指定したコメントを保持しているかのテーブルテスト(t *testing.T) {
	material := domain.NewMaterial()
	material.AddComment("コメント")

	cases := []struct {
		name     string
		comment  string
		expected bool
	}{
		{
			name:     "指定したコメントを保持している場合にtrue",
			comment:  "コメント",
			expected: true,
		},
		{
			name:     "指定したコメントを保持していない場合にfalse",
			comment:  "正しくないコメント",
			expected: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expected, material.HasThisComment(c.comment))
		})
	}

}
