package leetcode232

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyQueue(t *testing.T) {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	assert.Equal(t, 1, q.Peek())
	assert.Equal(t, 1, q.Pop())
	assert.Equal(t, 2, q.Peek())
	assert.Equal(t, false, q.Empty())
}
