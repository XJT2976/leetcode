package leetcode332

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindItinerary(t *testing.T) {
	type Input struct {
		tickets [][]string
		expect  []string
	}
	testInput := []Input{
		{
			[][]string{
				{"MUC", "LHR"},
				{"JFK", "MUC"},
				{"SFO", "SJC"},
				{"LHR", "SFO"},
			},
			[]string{"JFK", "MUC", "LHR", "SFO", "SJC"},
		},
		{
			[][]string{
				{"JFK", "SFO"},
				{"JFK", "ATL"},
				{"SFO", "ATL"},
				{"ATL", "JFK"},
				{"ATL", "SFO"},
			},
			[]string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"},
		},
	}

	for _, input := range testInput {
		result := FindItinerary(input.tickets)
		assert.Equal(t, input.expect, result)
	}
}
