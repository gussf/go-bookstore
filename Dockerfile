FROM golang

ADD . /go/src/github.com/gussf/go-bookstore/
WORKDIR /go/src/github.com/gussf/go-bookstore/

RUN go install .
CMD [ "bookstore" ]

EXPOSE 8080