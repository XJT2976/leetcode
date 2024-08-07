package leetcode22

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateParenthesis(t *testing.T) {
	assert.Equal(t, []string{"()"}, GenerateParenthesis(1))
	assert.Equal(t, []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		GenerateParenthesis(3))
}
