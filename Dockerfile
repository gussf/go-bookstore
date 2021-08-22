FROM golang

ADD . /src/go-bookstore/
WORKDIR /src/go-bookstore/

RUN go install .
# RUN go-bookstore

EXPOSE 15000
