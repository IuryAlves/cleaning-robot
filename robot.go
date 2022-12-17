package cleaning_robot

import "fmt"

type Direction string

const (
	North Direction = "NORTH"
	South Direction = "SOUTH"
	East  Direction = "EAST"
	West  Direction = "WEST"
)

type Logger interface {
	Log(msg string, args ...any)
}

type Robot struct {
	logger   Logger
	location Coordinate
	commands map[string]Command
}

// New instantiates a new robot
func New(x, y int) *Robot {
	r := &Robot{
		logger: &BasicLogger{},
		location: Coordinate{
			X: x,
			Y: y,
		},
		commands: map[string]Command{
			"clean": &CleanCommand{},
		},
	}
	_ = r.commands["clean"].Execute(x, y)
	return r
}

func (r *Robot) AddCommand(name string, c Command) {
	r.commands[name] = c
}

// Location returns the current robot location
func (r *Robot) Location() Coordinate {
	return r.location
}

// Move moves the robot in a direction N steps
func (r *Robot) Move(d Direction, steps int) error {
	r.logger.Log("moving %v step(s) in the %s direction", steps, d)
	switch d {
	case North:
		r.moveNorth(r.location.Y, steps)
	case South:
		r.moveSouth(r.location.Y, steps)
	case East:
		r.moveEast(r.location.X, steps)
	case West:
		r.moveWest(r.location.X, steps)
	default:
		return fmt.Errorf("invalid direction: %s", d)
	}
	return nil
}

// moveEast moves the robot in the x-axis towards east cleaning any uncleaned space
func (r *Robot) moveEast(x, steps int) {
	for i := x; i < x+steps; i++ {
		for _, cm := range r.commands {
			cm.Execute(i, r.location.Y)
		}
		r.location.X++
	}
}

// moveWest moves the robot in the x-axis towards west cleaning any uncleaned space
func (r *Robot) moveWest(x, steps int) {
	for i := x; i > x-steps; i-- {
		for _, cm := range r.commands {
			cm.Execute(i, r.location.Y)
		}
		r.location.X--
	}
}

// moveNorth moves the robot in the y-axis towards north cleaning any uncleaned space
func (r *Robot) moveNorth(y, steps int) {
	for i := y; i < y+steps; i++ {
		for _, cm := range r.commands {
			cm.Execute(r.location.X, i)
		}
		r.location.Y++
	}
}

// moveSouth moves the robot in the y-axis towards south cleaning any uncleaned space
func (r *Robot) moveSouth(y, steps int) {
	for i := y; i > y-steps; i-- {
		for _, cm := range r.commands {
			cm.Execute(r.location.X, i)
		}
		r.location.Y--
	}
}
