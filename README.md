# cleaning-robot

![CI workflow](https://github.com/iuryalves/cleaning-robot/actions/workflows/ci.yaml/badge.svg)

This project simulates the behavior of a cleaning robot.

The project is divided into two packages: `robot` and `app`

The `robot` package contains the core logic of the robot, and it can be used as a standalone library.
The `app` package is a http rest api that stores the robot's cleaning data into a Postgres database.

## Installing

```shell
go get github.com/IuryAlves/cleaning-robot
```

## Using as a standalone library

````go
package main

import "github.com/IuryAlves/cleaning-robot/robot"


// Initialise the robot into position (0, 0)
r := robot.New(0, 0)
// Move robot to position (1, 1)
r.Move(1, 1)

fmt.Println(r.Location()) // {1, 1}
````

For the full robot documentation see [here](https://github.com/IuryAlves/cleaning-robot/tree/main/robot/README.md).

## Running the application

### Configure the .env file

Create a copy of `.env.example` and replace the values with the configuration from your environment.
> **NOTE:** The default configuration from `.env.example` is meant to work on a local environment.

```shell
cp .env.example .env
```

> **NOTE:** The `.env` file is read automatically by `docker compose`.

### Running

```shell
docker compose up --build
```

### Testing the app

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

output:

```json
{"id":10,"timestamp":"2022-12-27T19:30:56.244522731Z","commands":2,"result":4,"duration":116578}
```
## Tests

### Unit tests

```shell
make test
```

### Integration tests

> **NOTE:** If running the integration tests against a new database, the migrations must be run first.

```shell
make integration-test
```