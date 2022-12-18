package robot

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Robot_clean_command(t *testing.T) {
	r, err := New(0, 0, &CleanCommand{})
	assert.NoError(t, err)
	r.Move(0, 1)

	clean, err := r.GetCommand("clean")
	assert.NoError(t, err)
	assert.Equal(t, 2, clean.(*CleanCommand).CleanedSpaces())
}

type MopCommand struct {
	moppedSpaces Coordinates
}

func (m *MopCommand) Name() string {
	return "mop"
}

func (m *MopCommand) OnInit(args ...any) error {
	x := args[0].(int)
	y := args[1].(int)
	m.Mop(x, y)
	return nil
}

func (m *MopCommand) OnMove(args ...any) error {
	x := args[0].(int)
	y := args[1].(int)
	m.Mop(x, y)
	return nil
}

func (m *MopCommand) Mop(x, y int) {
	m.moppedSpaces.Add(x, y)
}

func (m *MopCommand) MoppedSpaces() int {
	return m.moppedSpaces.Length()
}

func Test_Robot_mop_command(t *testing.T) {
	r, err := New(0, 0)
	assert.NoError(t, err)
	commands := []Command{&MopCommand{}}
	err = r.RegisterCommands(commands)
	assert.NoError(t, err)
	r.Move(1, 0)

	mop, err := r.GetCommand("mop")
	assert.NoError(t, err)
	assert.Equal(t, 1, mop.(*MopCommand).MoppedSpaces())
}

func TestRobot_RegisterCommands(t *testing.T) {
	clean := &CleanCommand{}
	r, err := New(0, 0, clean)
	assert.NoError(t, err)
	commands := []Command{
		clean,
	}

	err = r.RegisterCommands(commands)
	assert.EqualError(t, err, "command clean is already registered")
}
