## A simple bookstore API 
#### Intended for learning RESTful implementations and database access in Go

[![GoDoc](https://godoc.org/github.com/github.com/gussf/go-bookstore/?status.svg)](https://pkg.go.dev/github.com/gussf/go-bookstore/?tab=doc)

## Setup
``` shell
# Run PostgreSQL container on port 5432
# It should automatically setup the database as described in ./script/init_bookstore.sql
docker compose up -d
```

<br>

## Running the app
``` shell
. .env  # Load enviroment variables to access postgres
make run
```

<br>

## Test + Build + golangci-lint 
``` shell
# default target
make
```


<br>

# To Do
* Create a CD pipeline (deploy to AWS ECS Fargate?)
* Authentication (JWT?)
* Implement gRPC?
