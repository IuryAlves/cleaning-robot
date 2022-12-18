package cleaning_robot

import (
	"github.com/IuryAlves/cleaning-robot/robot"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRobot_Move_north_then_west(t *testing.T) {
	r, err := robot.New(0, 0)
	assert.NoError(t, err)
	err = Move(r, North, 2)
	assert.NoError(t, err)
	err = Move(r, West, 2)
	assert.NoError(t, err)

	assert.Equal(t, -2, r.Location().X)
	assert.Equal(t, 2, r.Location().Y)
}

func TestRobot_Move_south_then_east(t *testing.T) {
	r, err := robot.New(0, 5)
	assert.NoError(t, err)
	err = Move(r, South, 1)
	assert.NoError(t, err)
	err = Move(r, East, 3)
	assert.NoError(t, err)

	assert.Equal(t, 3, r.Location().X)
	assert.Equal(t, 4, r.Location().Y)
}

func TestRobot_Clean(t *testing.T) {
	r, err := robot.New(1, 1, &robot.CleanCommand{})
	assert.NoError(t, err)

	err = Move(r, North, 4)
	assert.NoError(t, err)
	err = Move(r, South, 2)
	assert.NoError(t, err)
	err = Move(r, East, 2)
	assert.NoError(t, err)
	err = Move(r, West, 1)
	assert.NoError(t, err)
	err = Move(r, East, 3)
	assert.NoError(t, err)

	clean, err := r.GetCommand("clean")
	assert.NoError(t, err)

	assert.Equal(t, 9, clean.(*robot.CleanCommand).CleanedSpaces())
}

func TestRobot_Clean_whole_area(t *testing.T) {
	r, err := robot.New(1, 1, &robot.CleanCommand{})
	assert.NoError(t, err)

	err = Move(r, North, 4)
	assert.NoError(t, err)
	err = Move(r, East, 2)
	assert.NoError(t, err)
	err = Move(r, South, 4)
	assert.NoError(t, err)
	err = Move(r, West, 1)
	assert.NoError(t, err)
	err = Move(r, North, 4)
	assert.NoError(t, err)

	clean, err := r.GetCommand("clean")
	assert.NoError(t, err)

	assert.Equal(t, 15, clean.(*robot.CleanCommand).CleanedSpaces())
}
