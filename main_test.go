package main

import (
	"backend/twitch-bot/models"
	"backend/twitch-bot/server/middleware"
	"net/http"
	"testing"
)

func TestServerRoutes(t *testing.T) {
	tests := models.UnitTesting{
		{
			Description:          "route used to enter song's into the song request queue",
			Path:                 "/song-request",
			Method:               "GET",
			ExpectedCodeResponse: 200,
		},
		{
			Description:          "non-exestient route",
			Path:                 "/i-dont-exist",
			Method:               "GET",
			ExpectedCodeResponse: 404,
		},
		{
			Description:          "route used to delete song's from the song request queue",
			Path:                 "/song-request-delete",
			Method:               "GET",
			ExpectedCodeResponse: 200,
		},
		{
			Description:          "route used to fetch all the song's from the song request queue",
			Path:                 "/songs",
			Method:               "GET",
			ExpectedCodeResponse: 200,
		},
		{
			Description:          "route used to fetch the users twitch infromation",
			Path:                 "/twitch/user",
			Method:               "GET",
			ExpectedCodeResponse: 401,
		},
	}

	app := middleware.Server()

	for _, test := range tests {
		req, _ := http.NewRequest(test.Method, test.Path, nil)
		if test.Path == "/song-request" {
			q := req.URL.Query()
			q.Add("channel", "testchannel")
			q.Add("user", "testuser")
			q.Add("q", "test")
			req.URL.RawQuery = q.Encode()
		}

		if test.Path == "/song-request-delete" {
			q := req.URL.Query()
			q.Add("channel", "testchannel")
			q.Add("id", "1")
			req.URL.RawQuery = q.Encode()
		}

		if test.Path == "/songs" {
			q := req.URL.Query()
			q.Add("channel", "testchannel")
			req.URL.RawQuery = q.Encode()
		}

		res, err := app.Test(req, -1)
		if err != nil {
			t.Fatal(err.Error())
		}
		if test.ExpectedCodeResponse != res.StatusCode {
			t.Fatalf("FAILED, the response code for path %s was not the same!", test.Path)
		}
		t.Log("PASSED. the response code was the expected code!")
	}
}
