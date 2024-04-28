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

## Make sure application is running


```shell
curl http://localhost:8080/status
```

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

## Validate class bookings

This is additional feature, I wanted to introduce out of the scope of the task. This is a utility tool where you can check capacity left on class and names of participants in a particular day.

```shell
curl http://localhost:8080/day/1970-01-01
```

> In a real-life situation, I never implement extra features without consultation. This should be discussed during refinement and included in the requirements. I did this only to simplify testing, taking some shortcuts (for prod-ready solution this would look a bit different).

# Tests

```shell
$ go test ./...
```
