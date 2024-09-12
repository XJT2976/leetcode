package leetcode2150

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLonely(t *testing.T) {
	actual := [][]int{
		FindLonely([]int{10, 6, 5, 8}),
		FindLonely([]int{1, 3, 5, 3}),
	}
	expect := [][]int{
		{10, 8},
		{1, 5},
	}

	for i := range actual {
		sort.Ints(actual[i])
		sort.Ints(expect[i])
		assert.Equal(t, expect[i], actual[i])
	}
}
