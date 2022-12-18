package robot

import (
	"fmt"
	"github.com/IuryAlves/cleaning-robot/logger"
)

type Logger interface {
	Log(msg string, args ...any)
}

type Command interface {
	OnInit(args ...any) error
	OnMove(args ...any) error
	Name() string
}

type Robot struct {
	Logger   Logger
	location Coordinate
	commands map[string]Command
}

// New instantiates a new robot
func New(x, y int, c ...Command) *Robot {
	r := &Robot{
		Logger: &logger.BasicLogger{},
		location: Coordinate{
			X: x,
			Y: y,
		},
		commands: map[string]Command{},
	}
	r.RegisterCommands(c)
	r.OnInit()
	return r
}

func (r *Robot) OnInit() {
	for _, c := range r.commands {
		_ = c.OnInit(r.location.X, r.location.Y)
	}
}

func (r *Robot) OnMove() {
	for _, c := range r.commands {
		_ = c.OnMove(r.location.X, r.location.Y)
	}
}

func (r *Robot) RegisterCommands(commands []Command) {
	for _, c := range commands {
		r.commands[c.Name()] = c
	}
}

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

func (r *Robot) Move(x, y int) {
	r.location.X = x
	r.location.Y = y
	r.OnMove()
}
