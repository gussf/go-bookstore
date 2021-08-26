## A simple bookstore API 
#### Intended for learning communication protocols and database access in Go

[![GoDoc](https://godoc.org/github.com/github.com/gussf/go-bookstore/?status.svg)](https://pkg.go.dev/github.com/gussf/go-bookstore/?tab=doc)


``` shell
go get github.com/gussf/go-bookstore
```

``` shell
# Run PostgreSQL container on port 5432
docker compose up -d

# Log into database
PGPASSWORD=postgres psql -h localhost -U postgres --dbname bookstore
```

``` sql
# Table creation
CREATE TABLE books (
    id SERIAL,
    title  varchar(100) NOT NULL,
    author varchar(100) NOT NULL,
    copies integer NOT NULL,
    price  integer NOT NULL,
    creation_date timestamp
);
```

``` shell
# Required environment variables in order to connect to database
export DRIVER=postgres
export USER=postgres
export HOST=localhost
export PORT=5432
export PASSWORD=postgres
export DBNAME=bookstore
```


<br>

# To Do
* Create a CI/CD pipeline (deploy to AWS ECS Fargate?)
* Authentication (JWT?)
* Implement gRPC?
