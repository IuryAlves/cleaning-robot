package robot

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRobot_Move_north_then_west(t *testing.T) {
	r, err := New(0, 0)
	assert.NoError(t, err)
	err = MoveToDirection(r, North, 2)
	assert.NoError(t, err)
	err = MoveToDirection(r, West, 2)
	assert.NoError(t, err)

	assert.Equal(t, -2, r.Location().X)
	assert.Equal(t, 2, r.Location().Y)
}

func TestRobot_Move_south_then_east(t *testing.T) {
	r, err := New(0, 5)
	assert.NoError(t, err)
	err = MoveToDirection(r, South, 1)
	assert.NoError(t, err)
	err = MoveToDirection(r, East, 3)
	assert.NoError(t, err)

	assert.Equal(t, 3, r.Location().X)
	assert.Equal(t, 4, r.Location().Y)
}

func TestRobot_Clean(t *testing.T) {
	r, err := New(1, 1, &CleanCommand{})
	assert.NoError(t, err)

	err = MoveToDirection(r, North, 4)
	assert.NoError(t, err)
	err = MoveToDirection(r, South, 2)
	assert.NoError(t, err)
	err = MoveToDirection(r, East, 2)
	assert.NoError(t, err)
	err = MoveToDirection(r, West, 1)
	assert.NoError(t, err)
	err = MoveToDirection(r, East, 3)
	assert.NoError(t, err)

	clean, err := r.GetCommand("clean")
	assert.NoError(t, err)

	assert.Equal(t, 9, clean.(*CleanCommand).CleanedSpaces())
}

func TestRobot_Clean_whole_area(t *testing.T) {
	r, err := New(1, 1, &CleanCommand{})
	assert.NoError(t, err)

	err = MoveToDirection(r, North, 4)
	assert.NoError(t, err)
	err = MoveToDirection(r, East, 2)
	assert.NoError(t, err)
	err = MoveToDirection(r, South, 4)
	assert.NoError(t, err)
	err = MoveToDirection(r, West, 1)
	assert.NoError(t, err)
	err = MoveToDirection(r, North, 4)
	assert.NoError(t, err)

	clean, err := r.GetCommand("clean")
	assert.NoError(t, err)

	assert.Equal(t, 15, clean.(*CleanCommand).CleanedSpaces())
}
