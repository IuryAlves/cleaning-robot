# cleaning-robot

This project simulates the behavior of a cleaning robot

This project is divided into two directories: `robot` and `app`

The `robot` package contains the core logic of the robot, and it can be used as a standalone library.
The `app` package is a http api that stores the robot's cleaning data into a Postgres database.

## Installing

```shell
go get github.com/IuryAlves/cleaning-robot
```

## Using as a standalone library

````go
package main

import "github.com/IuryAlves/cleaning-robot/robot"


r := robot.New(0, 0)
r.Move(1, 1)

fmt.Println(r.Location()) // {1, 1}
````

For the full robot documentation see [here](https://github.com/IuryAlves/cleaning-robot/tree/main/robot/README.md).

## Running the application

### Configure the .env file

Create a copy of `.env.example` and replace the values with the configuration from your environment.
> **NOTE:** The default configuration `.env.example` is meant to work on a local environment.


```shell
cp .env.example .env
```

The `.env` file is read automatically by `docker compose`.

```shell
docker compose up --build
```

### Running the migrations

When the database is created, the migrations need to run.
The easiest way of running the migrations is by doing:

```shell
docker compose run server --migrate 
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

## Tests

### Unit tests

```shell
make test
```

### Integration tests

!! NOTE: If running the integration tests against a new database, the migrations must be run first.

```shell
make integration-test
```