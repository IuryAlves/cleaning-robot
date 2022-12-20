package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IuryAlves/cleaning-robot/app/svc"
	"github.com/IuryAlves/cleaning-robot/robot"
	"net/http"
)

// Start TODO: maybe move to svc ?
type Start struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// EnterPathData TODO: maybe move to svc ?
type EnterPathData struct {
	Start    Start         `json:"start"`
	Commands []svc.Command `json:"commands"`
}

func EnterPathHandler(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
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

	service := svc.New(svc.WithRobot(r), svc.WithDefaultStorageClient())
	resp, err := service.Move(ctx, d.Commands)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Fprint(w, err.Error())
		w.WriteHeader(500)
	}
}
