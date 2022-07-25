package main

import (
	"net/http"
	"testing"

	webserver "github.com/MaximillianoNico/go-clean-architecture/app/infrastructure"
	utils "github.com/gofiber/fiber/v2/utils"
)

func TestIndexRoute(t *testing.T) {
	apps := webserver.NewApp("v1.0.0")

	svc := apps.RunServer()

	tests := []struct {
		description string

		// Test input
		route string

		// Expected output
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "index route",
			route:         "/api/health-check",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "OK",
		},
	}

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route
		// from the test case
		req, err := http.NewRequest(
			"GET",
			test.route,
			nil,
		)

		resp, err := svc.Test(req)

		utils.AssertEqual(t, nil, err, "app.Test")
		utils.AssertEqual(t, test.expectedCode, resp.StatusCode, "Status code")
	}

}
