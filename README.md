## A simple bookstore microservice 
#### Intended for learning communication protocols and database access in Go

[![GoDoc](https://godoc.org/github.com/github.com/gussf/go-bookstore/?status.svg)](https://pkg.go.dev/github.com/gussf/go-bookstore/?tab=doc)

``` shell 
# pull and run docker image
docker pull gussf/go-bookstore:v0.1
docker run -p 15000:15000 gussf/go-bookstore:v0.1

# test a HTTP method on port 15000
curl -v  localhost:15000/books
```

