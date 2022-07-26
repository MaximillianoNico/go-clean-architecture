# Sample Go Service
This project shows an example of how to write go with clean architecture.

For this service using gofiber as a framework

## Start Project
- `go mod tidy`
- run `make dev`
- hit endpoint `http://localhost:3000/api/health-check` for makesure the app is works

## Structure Project
```
project
│   README.md
│   Makefile   
│
└───app
│   └───infrastructure
|   |   └───middleware
│   │   webserver.go
│   │
│   └───interface
│       │   router.go
│       └───controllers
│   
└───cmd
|   │   server.go
|
└───test
    │   healthcheck_test.go
```