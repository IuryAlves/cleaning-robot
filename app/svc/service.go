package svc

import (
	"context"
	"github.com/IuryAlves/cleaning-robot/app/storage"
	"github.com/IuryAlves/cleaning-robot/robot"
	"time"
)

type Service struct {
	robot         *robot.Robot
	StorageClient *storage.Client
}

func New(options ...func(*Service)) *Service {
	s := &Service{}
	for _, o := range options {
		o(s)
	}
	return s
}

type Command struct {
	Direction robot.Direction `json:"direction"`
	Steps     int             `json:"steps"`
}

func (svc *Service) Move(ctx context.Context, commands []Command) (storage.Executions, error) {
	t := time.Now()
	for _, c := range commands {
		err := robot.MoveToDirection(svc.robot, c.Direction, c.Steps)
		if err != nil {
			return storage.Executions{}, err
		}
	}
	return svc.InsertExecution(ctx, len(commands), t)
}

func (svc *Service) InsertExecution(ctx context.Context, commands int, t time.Time) (storage.Executions, error) {
	duration := time.Since(t).Nanoseconds()
	cleanCommand, err := svc.robot.GetCommand("clean")
	if err != nil {
		return storage.Executions{}, err
	}
	e := storage.Executions{
		Result:    cleanCommand.(*robot.CleanCommand).CleanedSpaces(),
		Commands:  commands,
		Duration:  time.Duration(duration),
		Timestamp: t,
	}
	r, err := svc.StorageClient.InsertExecution(ctx, e)
	if err != nil {
		return storage.Executions{}, err
	}
	return r, nil
}

func WithRobot(r *robot.Robot) func(*Service) {
	return func(s *Service) {
		s.robot = r
	}
}

func WithDefaultStorageClient() func(*Service) {
	return func(s *Service) {
		s.StorageClient = storage.New(storage.WithPostgres())
	}
}
