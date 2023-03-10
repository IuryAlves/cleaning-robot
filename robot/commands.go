package robot

type CleanCommand struct {
	cleanedArea Coordinates
}

func (c *CleanCommand) Name() string {
	return "clean"
}

// OnMove cleans the robot's new position
func (c *CleanCommand) OnMove(args ...any) error {
	return c.Execute(args...)
}

// OnInit cleans the robot's current position
func (c *CleanCommand) OnInit(args ...any) error {
	return c.Execute(args...)
}

func (c *CleanCommand) Execute(args ...any) error {
	x := args[0].(int)
	y := args[1].(int)
	c.clean(x, y)
	return nil
}

// clean cleans an uncleaned space
func (c *CleanCommand) clean(x, y int) {
	if !c.isSpaceAlreadyCleaned(x, y) {
		c.cleanedArea.Add(
			x,
			y,
		)
	}
}

// isSpaceAlreadyCleaned checks if a given coordinate is already cleaned
func (c *CleanCommand) isSpaceAlreadyCleaned(x, y int) bool {
	for _, c := range c.cleanedArea.GetCoordinates() {
		if c.X == x && c.Y == y {
			return true
		}
	}
	return false
}

// CleanedSpaces returns the number of cleaned spaces
func (c *CleanCommand) CleanedSpaces() int {
	return c.cleanedArea.Length()
}
