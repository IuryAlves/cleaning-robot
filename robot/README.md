# Robot

Package robot implements a robot.

## Creating a new robot

```go
x := 0 // x position
y := 0 // y position
robot := New(x, y)
```

## Moving the robot

Robots can move in a 2D plane.

```go
newXPos := 1
robot.Move(newXPos, 0) // Moves the robot in the x-axis
```

## Commands

The robot functionality can be extended by using commands.

A command is a struct that implements the `Command` interface:

```go
type Command interface {
    // OnInit Called when the robot is initialised
    OnInit(args ...any) error
    // OnMove Called when the robot moves
    OnMove(args ...any) error
    // Name String representation of the command
    Name() string
}
```

### Cleaning command

The robot package comes with the `CleanCommand` by default.
This commands cleans any vertex that the robot touches. This includes the robot's initial position.

#### Using the cleaning command

```go
robot := New(0, 0, &CleanCommand)
robot.Move(1, 0)

// GetCommand returns a Command interface
clean, err := robot.GetCommand("clean")
if err != nil {
	// handle error
}
// Since clean is a command interface, we need to convert it to a concrete type
cs := clean.(*CleanCommand).CleanedSpaces()
fmt.Println(cs) // 2
```

### Adding custom commands

Let's implement a custom command that mops the robot's current location

```go
// We create a struct to represent the mop command
// This struct has a list of coordinates of mopped spaces
type MopCommand struct {
	moppedSpaces Coordinates
}

// Implement the Command interface
func (m *MopCommand) Name() string {
	return "mop"
}

// Mop the starting position of the robot
func (m *MopCommand) OnInit(args ...any) error {
    x := args[0].(int)
    y := args[1].(int)
    m.Mop(x, y)
    return nil
}

// Mop the vertex that the robot moved to
func (m *MopCommand) OnMove(args ...any) error {
	x := args[0].(int)
	y := args[1].(int)
	m.Mop(x, y)
	return nil
}

// Mop logic
// Add the current coordinate to the list of mopped spaces
func (m *MopCommand) Mop(x, y int) {
	m.moppedSpaces.Add(x, y)
}

// Returns the list of mopped spaces
func (m *MopCommand) MoppedSpaces() int {
	return m.moppedSpaces.Length()
}
```

#### Integrating the mop command in a robot

```go
robot := New(0,0, &MopCommand{})
mop, err := robot.GetCommand("mop")
if err != nil {
	// handle error
}

ms := mop.(MopCommand).MoppedSpaces()
fmt.Println(ms) // 1
```

Commands can also be added after the robot is initialised by calling the `RegisterCommands` method.

```go
// Initialise a robot with no commands
robot := New(0,0)
commands := []Command{
    &MopCommand{},
}
// Register the mop command
robot.RegisterCommands(commands)
```


