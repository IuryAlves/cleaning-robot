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
	logger      Logger
	location    Coordinate
	cleanedArea Polygon
}

// New instantiates a new robot
func New(x, y int) *Robot {
	r := &Robot{
		logger: &BasicLogger{},
		location: Coordinate{
			X: x,
			Y: y,
		},
	}
	r.Clean(x, y)
	return r
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

// moveEast moves the robot in the x-axis towards east cleaning any non-cleaned space
func (r *Robot) moveEast(x, steps int) {
	for i := x; i < x+steps; i++ {
		r.Clean(i, r.location.Y)
		r.location.X++
	}
}

// moveWest moves the robot in the x-axis towards west cleaning any non-cleaned space
func (r *Robot) moveWest(x, steps int) {
	for i := x; i > x-steps; i-- {
		r.Clean(i, r.location.Y)
		r.location.X--
	}
}

// moveNorth moves the robot in the y-axis towards north cleaning any non-cleaned space
func (r *Robot) moveNorth(y, steps int) {
	for i := y; i < y+steps; i++ {
		r.Clean(r.location.X, i)
		r.location.Y++
	}
}

// moveSouth moves the robot in the y-axis towards south cleaning any non-cleaned space
func (r *Robot) moveSouth(y, steps int) {
	for i := y; i > y-steps; i-- {
		r.Clean(r.location.X, i)
		r.location.Y--
	}
}

// Clean cleans a space that is not already cleaned
func (r *Robot) Clean(x, y int) {
	if !r.isSpaceAlreadyCleaned(x, y) {
		r.cleanedArea.Add(
			x,
			y,
		)
	}
}

// isSpaceAlreadyCleaned checks if a given coordinate is already cleaned
func (r *Robot) isSpaceAlreadyCleaned(x, y int) bool {
	for _, c := range r.cleanedArea.Coordinates {
		if c.X == x && c.Y == y {
			return true
		}
	}
	return false
}

// CleanedSpaces returns the number of cleaned spaces
func (r *Robot) CleanedSpaces() int {
	return r.cleanedArea.Length()
}
