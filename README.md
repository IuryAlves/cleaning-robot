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

## HTTP Server

### Running locally

```shell
go run main.go
```

### Running in docker

```shell
docker compose build
docker compose up
```

### Testing the HTTP server

```shell
curl \
  -XPOST  \
  -H 'Content-Type: application/json' \
  -d '{
    "start": {
      "x":10,
      "y": 22
    },
    "commands": [
      {"direction": "east", "steps": 2}, 
      {"direction": "north", "steps": 1}
    ]
   }' \
  localhost:8080/tibber-developer-test/enter-path
```