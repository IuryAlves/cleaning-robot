package robot

type CleanCommand struct {
	start         Coordinate
	end           Coordinate
	cleanedSpaces int
}

func (c *CleanCommand) Name() string {
	return "clean"
}

// OnMove cleans the robot's new position
func (c *CleanCommand) OnMove(args ...any) error {
	x := args[0].(int)
	y := args[1].(int)
	c.clean(x, y)
	c.end.X = x
	c.end.Y = y
	return nil
}

// OnInit cleans the robot's current position
func (c *CleanCommand) OnInit(args ...any) error {
	x := args[0].(int)
	y := args[1].(int)
	c.start = Coordinate{X: x, Y: y}
	c.clean(x, y)
	return nil
}

// clean cleans an uncleaned space
func (c *CleanCommand) clean(x, y int) {
	if !c.isSpaceAlreadyCleaned(x, y) {
		c.cleanedSpaces++
	}
}

// isSpaceAlreadyCleaned checks if a given coordinate is already cleaned
func (c *CleanCommand) isSpaceAlreadyCleaned(x, y int) bool {
	return x >= c.start.X && y >= c.start.Y && x <= c.end.X && y <= c.end.Y
}

// CleanedSpaces returns the number of cleaned spaces
func (c *CleanCommand) CleanedSpaces() int {
	return c.cleanedSpaces
}
