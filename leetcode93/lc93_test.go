package leetcode93

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRestoreIpAddresses(t *testing.T) {
	type Input struct {
		s      string
		expect []string
	}
	testInput := []Input{
		{
			"25525511135",
			[]string{"255.255.11.135", "255.255.111.35"},
		},
		{
			"0000",
			[]string{"0.0.0.0"},
		},
	}

	for _, input := range testInput {
		result := RestoreIpAddresses(input.s)
		assert.Equal(t, input.expect, result)
	}
}
