# go-car-storage-api
A simple API that allows a user to search for locations to store vehicles


## Getting Started
### Prerequisites

For running/testing locally you'll need either [golang](https://go.dev/doc/install) or [docker](https://docs.docker.com/engine/install/) installed


### Setup

The app can be run through docker with the included [Dockerfile](Dockerfile) by running
`docker compose up --build` in this directory

Alternatively if you have golang installed you can run `go mod download && go run main.go`

The app defaults to using port 8083 but that can be changed by setting the `$APP_PORT` environment variable if running locally or by editing the port numbers in [docker-compose.yml](docker-compose.yml#L8) to whatever port you would like.

## Usage

The app has two routes defined.

- `/` is defined as an easy `GET` option to test you have the app running correctly and are able to hit it
- `/listings` is defined as a `POST` as a way to search for storage location options and expects a json object that matches the following shape
```

[
    {
        "length": 10,
        "quantity": 1
    },
    {
        "length": 20,
        "quantity": 2
    },
    {
        "length": 25,
        "quantity": 1
    },
]
```

### Testing

The app has a basic unit test to check the fitment functionality and all tests can be run by running `go test ./...`
