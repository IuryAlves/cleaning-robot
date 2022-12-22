package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IuryAlves/cleaning-robot/app/svc"
	"github.com/IuryAlves/cleaning-robot/logger"
	"github.com/IuryAlves/cleaning-robot/robot"
	"net/http"
)

func EnterPathHandler(w http.ResponseWriter, req *http.Request) {
	l := logger.BasicLogger{}
	l.Log("Handling request")
	defer l.Log("Request ended %s")
	ctx := context.Background()
	// parse request data
	var mr svc.MoveRequest
	err := json.NewDecoder(req.Body).Decode(&mr)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	// Create a new robot
	r, err := robot.New(mr.Start.X, mr.Start.Y, &robot.CleanCommand{})
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	service := svc.New(svc.WithRobot(r), svc.WithDefaultStorageClient())
	resp, err := service.Move(ctx, mr.Commands)
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
