package robot

import (
	"fmt"
)

// Command specifies the methods that a Robot command must implement
type Command interface {
	// OnInit Called when the robot is initialised
	OnInit(args ...any) error
	// OnMove Called when the robot moves
	OnMove(args ...any) error
	// Name String representation of the command
	Name() string
}

type Robot struct {
	// The current robot location
	location Coordinate
	// List of commands that the robot can execute
	commands map[string]Command
}

// New instantiates a new robot
func New(x, y int, c ...Command) (*Robot, error) {
	r := &Robot{
		location: Coordinate{
			X: x,
			Y: y,
		},
		commands: map[string]Command{},
	}
	if err := r.RegisterCommands(c); err != nil {
		return nil, err
	}
	r.OnInit()
	return r, nil
}

// OnInit is an event called when the robot is initialised
// It calls all the commands that are subscribed
func (r *Robot) OnInit() {
	for _, c := range r.commands {
		_ = c.OnInit(r.location.X, r.location.Y)
	}
}

// OnMove is an event that is called when the robot moves
// It calls all the commands that are subscribed
func (r *Robot) OnMove() {
	for _, c := range r.commands {
		_ = c.OnMove(r.location.X, r.location.Y)
	}
}

// RegisterCommands registers one or move commands to the Robot
// If the robot already has the command registered an error is returned
func (r *Robot) RegisterCommands(commands []Command) error {
	for _, c := range commands {
		name := c.Name()
		if r.commands[name] != nil {
			return fmt.Errorf("command %s is already registered", name)
		}
		r.commands[name] = c
	}
	return nil
}

// GetCommand finds the command registered in the robot and returns it
// clients must convert the result of GetCommand to a concrete command type
// c, _ := r.GetCommand("<my-command>")
// c.(*ConcreteCommandType).ConcreteCommandMethod()
func (r *Robot) GetCommand(name string) (Command, error) {
	c := r.commands[name]
	if c == nil {
		return nil, fmt.Errorf("command %s is not registered", name)
	}
	return c, nil
}

// Location returns the current robot location
func (r *Robot) Location() Coordinate {
	return r.location
}

// Move moves the robot to a new x,y position
func (r *Robot) Move(x, y int) {
	r.location.X = x
	r.location.Y = y
	r.OnMove()
}
