Glofox Challenge
=

# Introduction

This is a solution to the [task](<Backend Task.pdf>) for a Senior SWE Glofox position at ABC Fitness.

# How to run

## Docker

Application comes dockerized and requires runnning only these two commands to run.

```shell
$ docker build . -t glofox
$ docker run -p 8080:8080 -d glofox
```

> By default it listens on 8080, but feel free to choose whatever you want. Image name is also editable.

## Source

If no docker is present on your system (why?), you can run the application executing following commands:

```shell
$ go run main.go
```

> First, make sure no other application listens to 8080.

# How to use

## Make sure application is running

Execute following curl command and confirm received response status of 200. This endpoint has no other purpose than performing a sanity check.

```shell
curl -i http://localhost:8080/status
```

## Create a class

Before booking a class, a system has to be populated with classes. A class can be added with similar data as these:

```shell
curl -X PUT http://localhost:8080/classes \
	-H 'Content-Type: application/json' \
	-H 'Accept: application/json' \
	-d '{"name": "Foo", "start_date": "1970-01-01", "end_date": "1970-01-02", "capacity": 1}'
```

## Register a booking

Once classes are injected, a booking can be registered. A booking can be registered with similar data as these:

```shell
curl -X POST http://localhost:8080/bookings \
	-H 'Content-Type: application/json' \
	-H 'Accept: application/json' \
	-d '{"name": "Bar", "date": "1970-01-01"}'
```

Keep in mind, there has to be a class happening on the specified date, and class has to have capacity left to accommodate request.

## Validate class bookings

This is additional feature, I wanted to introduce out of the scope of the task. This is a utility tool where you can check capacity left on class and names of participants in a particular day.

```shell
curl http://localhost:8080/day/1970-01-01
```

> In a real-life situation, I never implement extra features without consultation. This should be discussed during refinement and included in the requirements. I did this only to simplify testing, taking some shortcuts (for prod-ready solution this would look a bit different).

# Implementation

## Solution

- There's an openapi spec available. It can be imported in postman to facilitate testing.
- I've decided to use built in libraries, and not install any dependencies (like `mux`). I'm relatively new to go, and I'm not familar with anything else.
- I've decided to follow the suggestion and not use any external databases and store everything in memory. For this, I'm using a hashmaps. They aren't particularly lightweight, but could be replaced by databases without a problem (DI).
- I've decided to validate capacity when booking a class. That makes more sense to me and gives a chance to write a little bit more logic.
- I separated input reading and validation, business and storage layers to avoid putting everything in a one place.

## Comments

- I'm not happy with validation. Perhaps there are libraries that help with that and let perform more sophisticated calculations.
- Maybe I'd separate domain models from request data. It might seem they're the same objects, but I would feel more comfortable having request DTOs passed deeper into the domain and then instantiating domain models.
- I would like service to return result objects (DTSs) containing operation status (if it was successful or not, an error if any) and pick it up in app. That's something I'd feel beter about than simply returning models and errors.
- I'm not sure about the correct way to do logging. It feels not right to have it all scattered around, but maybe that's a go "thing".
- If the solution was using a database, I would provide a docker-compose specs.
- I'm not sure if this is possible in go, but I would appreciate if domain models were created with New* methods on unmarshalling.
- I've not implemented any locking mechanisms, but I know about possibility of a race condition. Lock should be obtained on storages.
- I'm not happy with app and service tests. These tests are dependent on each other and I don't know how to make go tests run in parallel.

# Tests

```shell
$ go test ./...
```
