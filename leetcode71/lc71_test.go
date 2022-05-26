package leetcode71

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	type Input struct {
		path   string
		expect string
	}
	testInput := []Input{
		{
			"/home/",
			"/home",
		},
		{
			"/../",
			"/",
		},
		{
			"/home//foo/",
			"/home/foo",
		},
		{
			"/.",
			"/",
		},
		{
			"/...",
			"/...",
		},
	}

	for _, input := range testInput {
		result := SimplifyPath(input.path)
		assert.Equal(t, input.expect, result)
	}
}
