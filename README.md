## A simple bookstore API 
#### Intended for learning communication protocols and database access in Go

[![GoDoc](https://godoc.org/github.com/github.com/gussf/go-bookstore/?status.svg)](https://pkg.go.dev/github.com/gussf/go-bookstore/?tab=doc)


``` shell
go get github.com/gussf/go-bookstore
```

``` shell
# Required environment variables in order to connect to database
export DRIVER=<driver>  ex: postgres
export USER=<db-user>
export HOST=<db-endpoint>
export PORT=<port>
export PASSWORD=<db-password>
export DBNAME=<dbname>
```

``` sql
# Table creation
CREATE TABLE books (
    id SERIAL,
    title  varchar(100) NOT NULL,
	author varchar(100) NOT NULL,
	copies integer NOT NULL,
	price  integer NOT NULL,
	creation_date date
);
```

<br>
<hr>

# To Do
* Create a CI/CD pipeline (deploy to AWS ECS Fargate?)
* Handle DELETE http method
* Authentication (JWT?)
* Implement gRPC?

<hr> 
<br>

``` shell 
# pull and run docker image
docker pull gussf/bookstore
# sh .env  # Load environment variables
./run-sh
``` 