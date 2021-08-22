## A simple bookstore microservice 
#### Intended for learning communication protocols and database access in Go
#### This project is developed using a AWS RDS running PostgreSQL

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

``` shell
# List all books
curl -v localhost:15000/books/

# List specific book 
curl -v localhost:15000/books/1
```

``` shell 
# OUT OF DATE
# pull and run docker image
docker pull gussf/go-bookstore:v0.1
docker run -p 15000:15000 gussf/go-bookstore:v0.1
``` 


