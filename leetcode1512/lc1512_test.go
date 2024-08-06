package leetcode1512

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumIdenticalPairs(t *testing.T) {
	assert.Equal(t, NumIdenticalPairs([]int{1,2,3,1,1,3}), 4)
	assert.Equal(t, NumIdenticalPairs([]int{1,1,1,1}), 6)
	assert.Equal(t, NumIdenticalPairs([]int{1,2,3}), 0)
	assert.Equal(t, NumIdenticalPairs([]int{1}), 0)
}