package cleaning_robot

import (
	"fmt"
	"github.com/IuryAlves/cleaning-robot/robot"
)

type Direction string

const (
	North Direction = "NORTH"
	South Direction = "SOUTH"
	East  Direction = "EAST"
	West  Direction = "WEST"
)

// Move converts a direction and steps to x,y coordinates and moves the Robot
func Move(r *robot.Robot, d Direction, steps int) error {
	r.Logger.Log("moving %v step(s) in the %s direction", steps, d)
	switch d {
	case North:
		yPos := r.Location().Y
		for i := 0; i < steps; i++ {
			yPos++
			r.Move(r.Location().X, yPos)
		}
	case South:
		yPos := r.Location().Y
		for i := 0; i < steps; i++ {
			yPos--
			r.Move(r.Location().X, yPos)
		}
	case East:
		xPos := r.Location().X
		for i := 0; i < steps; i++ {
			xPos++
			r.Move(xPos, r.Location().Y)
		}
	case West:
		xPos := r.Location().X
		for i := 0; i < steps; i++ {
			xPos--
			r.Move(xPos, r.Location().Y)
		}
	default:
		return fmt.Errorf("invalid direction: %s", d)
	}
	return nil
}
