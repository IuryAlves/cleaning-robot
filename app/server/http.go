package server

import (
	"context"
	"encoding/json"
	"github.com/IuryAlves/cleaning-robot/app/svc"
	"github.com/IuryAlves/cleaning-robot/logger"
	"github.com/IuryAlves/cleaning-robot/robot"
	"net/http"
)

func EnterPathHandler(w http.ResponseWriter, req *http.Request) {
	l := logger.BasicLogger{}
	l.Info("Handling request")
	// parse request data
	var mr svc.MoveRequest
	err := json.NewDecoder(req.Body).Decode(&mr)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		l.Error(err.Error())
		return
	}

	if mr.Commands == nil {
		errMsg := "request must have at least one command"
		w.WriteHeader(422)
		w.Write([]byte(errMsg))
		l.Error(errMsg)
		return
	}

	// Create a new robot
	r, err := robot.New(mr.Start.X, mr.Start.Y, &robot.CleanCommand{})
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		l.Error(err.Error())
		return
	}

	service := svc.New(svc.WithRobot(r), svc.WithDefaultStorageClient())
	ctx := context.Background()
	resp, err := service.Move(ctx, mr.Commands)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		l.Error(err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		l.Error(err.Error())
	}
}
