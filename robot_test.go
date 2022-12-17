package cleaning_robot

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRobot_Move_north(t *testing.T) {
	robot := New(0, 0)
	robot.Move(North, 2)

	assert.Equal(t, 0, robot.Location().X)
	assert.Equal(t, 2, robot.Location().Y)
}

func TestRobot_Move_south_then_east(t *testing.T) {
	robot := New(0, 5)
	robot.Move(South, 1)
	robot.Move(East, 3)

	assert.Equal(t, 3, robot.Location().X)
	assert.Equal(t, 4, robot.Location().Y)
}

func TestRobot_Clean(t *testing.T) {
	robot := New(1, 1)

	err := robot.Move(North, 4)
	assert.NoError(t, err)
	err = robot.Move(South, 2)
	assert.NoError(t, err)
	err = robot.Move(East, 2)
	assert.NoError(t, err)
	err = robot.Move(West, 1)
	assert.NoError(t, err)
	err = robot.Move(East, 3)
	assert.NoError(t, err)

	assert.Equal(t, 8, robot.commands["clean"].(*CleanCommand).CleanedSpaces())
}

func TestRobot_Clean_whole_area(t *testing.T) {
	robot := New(1, 1)

	err := robot.Move(North, 4)
	assert.NoError(t, err)
	err = robot.Move(East, 2)
	assert.NoError(t, err)
	err = robot.Move(South, 4)
	assert.NoError(t, err)
	err = robot.Move(West, 1)
	assert.NoError(t, err)
	err = robot.Move(North, 4)
	assert.NoError(t, err)

	assert.Equal(t, 15, robot.commands["clean"].(*CleanCommand).CleanedSpaces())
}
