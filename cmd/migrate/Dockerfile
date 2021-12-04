# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -o migrations cmd/migrate/main.go 
RUN ls -al

ENTRYPOINT ./migrations