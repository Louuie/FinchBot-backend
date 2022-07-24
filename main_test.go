package main

import (
	"backend/twitch-bot/server/middleware"
	"net/http"
	"testing"
)

func TestServerRoutes(t *testing.T) {
	// tests := []struct {
	// 	description string

	// 	// Test input
	// 	route string

	// 	// Expected output
	// 	expectedError bool
	// 	expectedCode  int
	// 	expectedBody  string
	// }{
	// 	{
	// 		description:   "index route",
	// 		route:         "/",
	// 		expectedError: false,
	// 		expectedCode:  200,
	// 		expectedBody:  "OK",
	// 	},
	// 	{
	// 		description:   "non existing route",
	// 		route:         "/i-dont-exist",
	// 		expectedError: false,
	// 		expectedCode:  404,
	// 		expectedBody:  "Cannot GET /i-dont-exist",
	// 	},
	// }
	tests := []struct {
		description          string
		path                 string
		method               string
		expectedResponseCode int
	}{
		{
			"route used to enter song's into the song request queue",
			"/song-request",
			"GET",
			200,
		},
		{
			"non-exestient route",
			"/i-dont-exist",
			"GET",
			404,
		},
		{
			"route used to delete song's from the song request queue",
			"/song-request-delete",
			"GET",
			200,
		},
	}
	app := middleware.Server()

	for _, test := range tests {
		req, _ := http.NewRequest(
			test.method,
			test.path,
			nil,
		)
		if test.path == "/song-request" {
			q := req.URL.Query()
			q.Add("channel", "testchannel")
			q.Add("user", "testuser")
			q.Add("q", "test")
			req.URL.RawQuery = q.Encode()
		}

		if test.path == "/song-request-delete" {
			q := req.URL.Query()
			q.Add("channel", "testchannel")
			q.Add("id", "2")
			req.URL.RawQuery = q.Encode()
		}

		res, err := app.Test(req, -1)
		if err != nil {
			t.Fatal(err.Error())
		}
		if test.expectedResponseCode != res.StatusCode {
			t.Fatalf("FAILED, the response code for path %s was not the same!", test.path)
		}
		t.Log("PASSED. the response code was the expected code!")
	}
}
