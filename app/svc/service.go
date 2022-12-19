package svc

import (
	"context"
	"github.com/IuryAlves/cleaning-robot/app/storage"
	"github.com/IuryAlves/cleaning-robot/robot"
	"time"
)

type Service struct {
	robot *robot.Robot
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

func (svc *Service) Move(ctx context.Context, commands []Command) (storage.InsertRequest, error) {
	t := time.Now()
	for _, c := range commands {
		err := robot.MoveToDirection(svc.robot, c.Direction, c.Steps)
		if err != nil {
			return storage.InsertRequest{}, err
		}
	}
	return svc.AddExecution(ctx, len(commands), t)
}


func (svc *Service) AddExecution(ctx context.Context, commands int, t time.Time) (storage.InsertRequest, error) {
	duration := time.Since(t).Nanoseconds()
	cleanCommand, err := svc.robot.GetCommand("clean")
	if err != nil {
		return storage.InsertRequest{}, err
	}
	i := storage.InsertRequest{
		Result: cleanCommand.(*robot.CleanCommand).CleanedSpaces(),
		Commands: commands,
		Duration: time.Duration(duration),
		Timestamp: t,
	}
	return i, svc.StorageClient.Insert(ctx, i)
}

func WithRobot(r *robot.Robot) func (*Service) {
	return func (s *Service) {
		s.robot = r
	}
}

func WithDefaultStorageClient() func (*Service) {
	return func (s *Service) {
		s.StorageClient = storage.New(storage.WithPostgres())
	}
}
