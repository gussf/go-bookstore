## A simple bookstore API 
#### Intended for learning RESTful implementations and database access in Go

[![GoDoc](https://godoc.org/github.com/github.com/gussf/go-bookstore/?status.svg)](https://pkg.go.dev/github.com/gussf/go-bookstore/?tab=doc)


``` shell
go get github.com/gussf/go-bookstore
```

``` shell
# Run PostgreSQL container on port 5432
# It should automatically setup the database as described in ./script/init_bookstore.sql
docker compose up -d
```

``` shell
# Default target is build+lint
make

# Starting the app
make run
```

<br>

# To Do
* Create a CD pipeline (deploy to AWS ECS Fargate?)
* Authentication (JWT?)
* Implement gRPC?
