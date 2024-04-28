Glofox Challenge
=

# Introduction

This is a solution to the [task](<Backend Task.pdf>) for a Senior SWE Glofox position at ABC Fitness.

# How to run

```shell
$ go build .
$ go run
```

# How to use

## Create a class

```shell
curl -X PUT http://localhost:8080/classes \
	-H 'Content-Type: application/json' \
	-H 'Accept: application/json' \
	-d '{"name": "Foo", "start_date": "1970-01-01", "end_date": "1970-01-02", "capacity": 1}'
```

## Register a booking

```shell
curl -X POST http://localhost:8080/bookings \
	-H 'Content-Type: application/json' \
	-H 'Accept: application/json' \
	-d '{"name": "Bar", "date": "1970-01-01"}'
```

# Tests

```shell
$ go test ./...
```
