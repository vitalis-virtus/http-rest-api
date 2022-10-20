# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /app

ADD go.mod ./
ADD go.sum ./

RUN go mod download

COPY . .
COPY ./sql-scripts/ /docker-entrypoint-initdb.d/

RUN go build -o /main

EXPOSE 8080

CMD ["/main"]