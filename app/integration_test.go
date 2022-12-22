package app

import (
	"bytes"
	"encoding/json"
	"github.com/IuryAlves/cleaning-robot/app/server"
	"github.com/IuryAlves/cleaning-robot/app/storage"
	"github.com/IuryAlves/cleaning-robot/app/svc"
	"github.com/IuryAlves/cleaning-robot/robot"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
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

	assert.Equal(t, response.Result().StatusCode, 200)

	var result storage.Executions
	err = json.Unmarshal(response.Body.Bytes(), &result)
	assert.NoError(t, err)
	assert.Equal(t, 4, result.Result)
	assert.Equal(t, 2, result.Commands)
}

func TestServer_clean3x3_area(t *testing.T) {
	d := svc.MoveRequest {
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

	assert.Equal(t, response.Result().StatusCode, 200)

	var result storage.Executions
	err = json.Unmarshal(response.Body.Bytes(), &result)
	assert.NoError(t, err)
	assert.Equal(t, 9, result.Result)
	assert.Equal(t, 5, result.Commands)
}
