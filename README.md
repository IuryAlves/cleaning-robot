# cleaning-robot

Package cleaning-robot implements a robot

## Installing

```shell
go get github.com/IuryAlves/cleaning-robot
```

## Using as a library

````go
package main

import "github.com/IuryAlves/cleaning-robot/robot"


r := robot.New(0, 0)
r.Move(1, 1)

fmt.Println(r.Location()) // {1, 1}
````

For the full robot documentation see [here](https://github.com/IuryAlves/cleaning-robot/tree/main/robot/README.md).