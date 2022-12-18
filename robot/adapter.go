package robot

import (
	"fmt"
)

type Direction string

const (
	North Direction = "north"
	South Direction = "south"
	East  Direction = "east"
	West  Direction = "west"
)

// MoveToDirection is an adapter that translates a direction and a number of steps to x,y coordinates
// It calls robot.Move passing the x,y coordinates
func MoveToDirection(r *Robot, d Direction, steps int) error {
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
