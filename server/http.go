package server

import (
	"encoding/json"
	"fmt"
	"github.com/IuryAlves/cleaning-robot/robot"
	"net/http"
	"time"
)

type Start struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Command struct {
	Direction robot.Direction `json:"direction"`
	Steps     int             `json:"steps"`
}

type EnterPathData struct {
	Start    Start     `json:"start"`
	Commands []Command `json:"commands"`
}

type Result struct {
	Id        int
	Timestamp time.Time
	Commands  int
	Result    int
	Duration  int64
}

func execCommands(r *robot.Robot, commands []Command) error {
	for _, c := range commands {
		err := robot.MoveToDirection(r, c.Direction, c.Steps)
		if err != nil {
			return err
		}
	}
	return nil
}

func getCleanedSpaces(r *robot.Robot) (int, error) {
	c, err := r.GetCommand("clean")
	if err != nil {
		return 0, err
	}
	return c.(*robot.CleanCommand).CleanedSpaces(), nil
}

func writeResults(w http.ResponseWriter, r *robot.Robot, c []Command, t time.Time) {
	// Get number of cleaned spaces
	cs, err := getCleanedSpaces(r)
	if err != nil {
		fmt.Fprint(w, err.Error())
		w.WriteHeader(500)
		return
	}
	result := Result{
		Id:        1,
		Timestamp: time.Now(),
		Commands:  len(c),
		Result:    cs,
		Duration:  time.Since(t).Nanoseconds(),
	}

	// Write result
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		fmt.Fprint(w, err.Error())
		w.WriteHeader(500)
	}
}

func EnterPathHandler(w http.ResponseWriter, req *http.Request) {
	// parse request data
	var d EnterPathData
	err := json.NewDecoder(req.Body).Decode(&d)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	// Create a new robot
	r, err := robot.New(d.Start.X, d.Start.Y, &robot.CleanCommand{})
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	t := time.Now()
	// run robot commands
	err = execCommands(r, d.Commands)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	writeResults(w, r, d.Commands, t)
}
