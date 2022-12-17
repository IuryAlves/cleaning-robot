package cleaning_robot

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPolygon_Area(t *testing.T) {

	testCases := map[string]struct {
		input    []Coordinate
		expected int
	}{
		"Calculate area square": {
			input: []Coordinate{
				{
					X: 100,
					Y: 100,
				},
				{
					X: 100,
					Y: 200,
				},
				{
					X: 200,
					Y: 200,
				},
				{
					X: 200,
					Y: 100,
				},
			},
			expected: 20000,
		},
		"Calculate area triangle": {
			input: []Coordinate{
				{
					X: 1,
					Y: 1,
				},
				{
					X: 4,
					Y: 4,
				},
				{
					X: 5,
					Y: 10,
				},
			},
			expected: 6,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			p := Polygon{
				Coordinates: tc.input,
			}
			assert.Equal(t, tc.expected, p.Area())
		})
	}
}
