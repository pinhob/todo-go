## ⚠️ Work in progress ⚠️
# To-do list in Go
This is a simple to-do CLI and API built with Go. The CLI is based in the book [Powerful Command-Line Applications in Go](https://pragprog.com/titles/rggo/powerful-command-line-applications-in-go/) and in the [Abah Joseph tutorial](https://www.youtube.com/watch?v=j1CXoOQXbco).

## Requisites
```
Go 1.22.1
```

## How to run with Docker

## How to run without Docker

To start the API, run thesses commands:
```
cd cmd/api
go run main.go
```

To start the CLI, run theses commands:
```
cd cmd/cli
go run main.go
```
## Testing
Run the following command in the root directory to run the unit tests:
```
go test -v
```
To run the CLI integration test, run:
```
cd cmd/todo
go test -v 
```
