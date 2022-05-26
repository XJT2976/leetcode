package leetcode225

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyStack(t *testing.T) {
	q := Constructor()
	q.Push(1)
	q.Push(2)

	assert.Equal(t, 2, q.Top())
	assert.Equal(t, 2, q.Pop())
	assert.Equal(t, false, q.Empty())
}
