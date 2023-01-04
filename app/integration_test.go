package app

import (
	"bytes"
	"encoding/json"
	"github.com/IuryAlves/cleaning-robot/app/server"
	"github.com/IuryAlves/cleaning-robot/app/storage"
	"github.com/IuryAlves/cleaning-robot/app/svc"
	"github.com/IuryAlves/cleaning-robot/robot"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_move_and_clean(t *testing.T) {
	d := svc.MoveRequest{
		Start: svc.Start{X: 10, Y: 22},
		Commands: []svc.Command{
			{
				Direction: robot.East,
				Steps:     2,
			},
			{
				Direction: robot.North,
				Steps:     1,
			},
		},
	}
	b, err := json.Marshal(&d)
	assert.NoError(t, err)

	request, _ := http.NewRequest(http.MethodPost, "/tibber-developer-test/enter-path", bytes.NewReader(b))
	response := httptest.NewRecorder()
	server.EnterPathHandler(response, request)

	assert.Equal(t, 200, response.Result().StatusCode)

	var result storage.Executions
	err = json.Unmarshal(response.Body.Bytes(), &result)
	assert.NoError(t, err)
	assert.Equal(t, 4, result.Result)
	assert.Equal(t, 2, result.Commands)
}

func TestServer_clean3x3_area(t *testing.T) {
	d := svc.MoveRequest{
		Start: svc.Start{X: 0, Y: 0},
		Commands: []svc.Command{
			{
				Direction: robot.North,
				Steps:     2,
			},
			{
				Direction: robot.East,
				Steps:     2,
			},
			{
				Direction: robot.South,
				Steps:     2,
			},
			{
				Direction: robot.West,
				Steps:     1,
			},
			{
				Direction: robot.North,
				Steps:     1,
			},
		},
	}
	b, err := json.Marshal(&d)
	assert.NoError(t, err)

	request, _ := http.NewRequest(http.MethodPost, "/tibber-developer-test/enter-path", bytes.NewReader(b))
	response := httptest.NewRecorder()
	server.EnterPathHandler(response, request)

	assert.Equal(t, 200, response.Result().StatusCode)

	var result storage.Executions
	err = json.Unmarshal(response.Body.Bytes(), &result)
	assert.NoError(t, err)
	assert.Equal(t, 9, result.Result)
	assert.Equal(t, 5, result.Commands)
}

func Test_invalid_input_data(t *testing.T) {
	d := struct{}{}
	b, err := json.Marshal(&d)
	assert.NoError(t, err)

	request, _ := http.NewRequest(http.MethodPost, "/tibber-developer-test/enter-path", bytes.NewReader(b))
	response := httptest.NewRecorder()
	server.EnterPathHandler(response, request)

	assert.Equal(t, 422, response.Result().StatusCode)
	assert.Equal(t, "request must have at least one command", response.Body.String())

}

func Test_constraints(t *testing.T) {
	d := svc.MoveRequest{
		Start: svc.Start{
			X: 0,
			Y: 0,
		},
	}
	rand.Seed(time.Now().Unix())
	directions := []robot.Direction{robot.North, robot.South, robot.West, robot.East}
	for i := 0; i < 10000; i++ {
		randomDirection := directions[rand.Int()%len(directions)]
		d.Commands = append(d.Commands, svc.Command{
			Steps:     1,
			Direction: randomDirection,
		})
	}
	b, err := json.Marshal(&d)
	assert.NoError(t, err)

	request, _ := http.NewRequest(http.MethodPost, "/tibber-developer-test/enter-path", bytes.NewReader(b))
	response := httptest.NewRecorder()
	server.EnterPathHandler(response, request)

	assert.Equal(t, 200, response.Result().StatusCode)

	var result storage.Executions
	err = json.Unmarshal(response.Body.Bytes(), &result)
	assert.NoError(t, err)

	assert.Equal(t, 10000, result.Commands)
}
