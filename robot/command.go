package robot

type Command interface {
	Execute(args ...any) error
}

type CleanCommand struct {
	cleanedArea Coordinates
}

func (c *CleanCommand) Execute(args ...any) error {
	x := args[0].(int)
	y := args[1].(int)
	c.Clean(x, y)
	return nil
}

// Clean cleans an uncleaned space
func (c *CleanCommand) Clean(x, y int) {
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
