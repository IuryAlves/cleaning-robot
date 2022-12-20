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

## Running the application

### Configure the .env file

Create a copy of `.env.example` and replace the values with the configuration from your environment.
> **NOTE:** The default configuration `.env.example` is meant to work on a local environment.


```shell
cp .env.example .env
```

The `.env` file is read automatically by `docker compose`.

```shell
docker compose build
docker compose up
```

### Running the migrations

When the database is created, the migrations need to run.
The easiest way of running the migrations is by doing `docker compose exec`

```shell
docker compose exec server ./cleaning-robot --migrate 
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